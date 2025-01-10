package handlers

import (
	"github.com/Njeri-Ngugi/toolbox/helpers"
	"github.com/Njeri-Ngugi/toolbox/postgres"
	"github.com/Njeri-Ngugi/toolbox/validation"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"users/internal/user/models"
	"users/internal/user/serializers"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request serializers.CreateUserRequest
	err := validation.ValidateRequest(w, r, &request)
	if err != nil {
		logrus.Error(err)
		return
	}

	// copy request into model data
	var model models.User
	err = copier.Copy(&model, &request)
	if err != nil {
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error(err)
		return
	}

	model.Password = hashedPassword

	// create user
	user, err := postgres.DbService.DAO.GetOrCreate(r.Context(), &model, &model)
	if err != nil {
		logrus.Error(err)
		helpers.HTTPErrorResponse(w, "Error creating user", http.StatusInternalServerError)
	}

	response := helpers.Response{
		Message: "User created successfully",
		Data:    user,
	}

	if err = helpers.HTTPResponse(w, response, http.StatusCreated); err != nil {
		logrus.Error(err)
		helpers.HTTPErrorResponse(w, "", http.StatusInternalServerError)
		return
	}

}
