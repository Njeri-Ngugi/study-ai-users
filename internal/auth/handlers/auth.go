package handlers

import (
	"github.com/Njeri-Ngugi/toolbox/auth"
	authSerializer "github.com/Njeri-Ngugi/toolbox/auth/serializers"
	"github.com/Njeri-Ngugi/toolbox/helpers"
	"github.com/Njeri-Ngugi/toolbox/postgres"
	"github.com/Njeri-Ngugi/toolbox/validation"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"net/http"
	"users/internal/auth/serializers"
	"users/internal/user/models"
)

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	// validate login request
	var request serializers.UserLoginRequest
	err := validation.ValidateRequest(w, r, &request)
	if err != nil {
		logrus.Error(err)
		return
	}

	// ensure that either username or email is provided
	if request.Email == "" && request.Username == "" {
		logrus.Error("Email or Username is required")
		helpers.HTTPErrorResponse(w, "Email or Username is required", http.StatusBadRequest)
		return
	}

	var condition models.User
	if request.Username != "" {
		condition.Username = request.Username
	}

	if request.Email != "" {
		condition.Email = request.Email
	}

	user, err := postgres.DbService.DAO.Get(r.Context(), &condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "User not found", http.StatusInternalServerError)

		return
	}

	if user != nil {
		userObj := user.(*models.User)

		// compare password with that of retrieved record
		isValid, err := auth.ComparePasswords(userObj.Password, []byte(request.Password))
		if err != nil {
			logrus.Error("error credentials don't match", err, isValid)
			helpers.HTTPErrorResponse(w, "invalid login credentials", http.StatusUnauthorized)
			return
		}

		if !isValid {
			logrus.Error("Username or Password is invalid")
			helpers.HTTPErrorResponse(w, "invalid login credentials", http.StatusUnauthorized)
			return
		}

		// generate auth token
		var authModel authSerializer.AuthModelData
		err = copier.Copy(&authModel, userObj)
		if err != nil {
			return
		}

		logrus.Infoln("auth model: ", authModel)
		token, err := auth.GenerateAuthToken(authModel)
		if err != nil {
			logrus.Error(err)
			helpers.HTTPErrorResponse(w, "Error validating user", http.StatusInternalServerError)
			return
		}

		// return success response
		responseData := serializers.UserLoginResponse{
			UserData: user,
			Token:    token,
		}

		response := helpers.Response{
			Message: "User login successful",
			Data:    responseData,
		}
		logrus.Infoln("successfully authenticated")

		if err = helpers.HTTPResponse(w, response, http.StatusOK); err != nil {
			logrus.Error("error sending http response", err)
			helpers.HTTPErrorResponse(w, "error returning http response", http.StatusInternalServerError)
			return
		}

	} else {
		logrus.Error("error retrieving user data")
		helpers.HTTPErrorResponse(w, "error retrieving user data", http.StatusInternalServerError)
		return
	}
}
