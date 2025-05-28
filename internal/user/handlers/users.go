package handlers

import (
	"fmt"
	"github.com/Njeri-Ngugi/toolbox/helpers"
	"github.com/Njeri-Ngugi/toolbox/postgres"
	"github.com/Njeri-Ngugi/toolbox/validation"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	academicModels "users/internal/academics/models"
	"users/internal/user/daos"
	"users/internal/user/models"
	"users/internal/user/serializers"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request serializers.CreateUserRequest
	err := validation.ValidateRequest(w, r, &request)
	if err != nil {
		logrus.Error("error validating create user request", err)
		return
	}

	// convert date of birth and completion date into dates
	formattedDOB, err := daos.ConvertDateIntoTime(request.DateOfBirth)
	if err != nil {
		logrus.Error(err)
		helpers.HTTPErrorResponse(w, "invalid date of birth format", http.StatusBadRequest)
		return
	}

	formattedCompletionDate, err := daos.ConvertDateIntoTime(request.CompletionDate)
	if err != nil {
		logrus.Error(err)
		helpers.HTTPErrorResponse(w, "invalid completion date", http.StatusBadRequest)
		return
	}

	// copy request into model data
	var model models.User
	err = copier.Copy(&model, &request)
	if err != nil {
		return
	}

	model.DateOfBirth = *formattedDOB
	model.CompletionDate = *formattedCompletionDate

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error(err)
		return
	}
	model.Password = hashedPassword

	// get or create institution
	conditionInstitution := &academicModels.Institution{
		InstitutionName: request.InstitutionName,
	}

	institution, _, err := postgres.DbService.DAO.GetOrCreate(r.Context(), conditionInstitution, conditionInstitution)
	if err != nil {
		logrus.Error(err)
		helpers.HTTPErrorResponse(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	institutionId := institution.(*academicModels.Institution).Id
	model.InstitutionId = institutionId

	// get or create course
	var newCourse interface{}
	conditionCourse := &academicModels.Course{
		CourseName: request.CourseName,
	}

	course, err := postgres.DbService.DAO.Get(r.Context(), conditionCourse)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			// create a course if the course wasn't found
			modelCourse := &academicModels.Course{
				CourseName:            request.CourseName,
				CourseCode:            daos.GenerateCourseCode(request.InstitutionName, request.CourseName),
				CourseDurationInYears: request.CourseDurationInYears,
				InstitutionId:         institutionId,
			}

			newCourse, _, err = postgres.DbService.DAO.GetOrCreate(r.Context(), modelCourse, modelCourse)
			if err != nil {
				logrus.Error(err)
				helpers.HTTPErrorResponse(w, "Failed to create user", http.StatusInternalServerError)
				return
			}
			model.CourseId = newCourse.(*academicModels.Course).Id
		} else {
			logrus.Error(err)
			helpers.HTTPErrorResponse(w, "Failed to create user", http.StatusInternalServerError)
			return
		}
	}

	if course != nil {
		model.CourseId = course.(*academicModels.Course).Id
	}

	// check if user exists
	var response helpers.Response
	createdUser, userExists, err := daos.CheckUserExistsByUsernameOrEmail(r.Context(), request.Username, request.Email)
	if err != nil {
		logrus.Error(err)
		helpers.HTTPErrorResponse(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	if !userExists {
		// get or create user
		user, _, err := postgres.DbService.DAO.GetOrCreate(r.Context(), &model, &model)
		if err != nil {
			// check for a duplicate key error
			if err.Error() == gorm.ErrDuplicatedKey.Error() {
				fmt.Println("Duplicate key error: Record already exists.")
				return
			} else {
				logrus.Error("error saving user", err)
				helpers.HTTPErrorResponse(w, "Error creating user", http.StatusInternalServerError)
				return
			}
		}
		response.Data = user
	} else {
		response.Data = createdUser
	}

	response.Message = "User created successfully"

	if err = helpers.HTTPResponse(w, response, http.StatusCreated); err != nil {
		logrus.Error(err)
		helpers.HTTPErrorResponse(w, "", http.StatusInternalServerError)
		return
	}

}
