package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ebcardoso/api-rest-golang/app/repository"
	"github.com/ebcardoso/api-rest-golang/app/types"
	"github.com/ebcardoso/api-rest-golang/config"
	"github.com/ebcardoso/api-rest-golang/utils/request"
	"github.com/ebcardoso/api-rest-golang/utils/response"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	repository *repository.Users
	configs    *config.Config
}

func NewAuth(configs *config.Config) *Auth {
	return &Auth{
		repository: repository.NewRepositoryUsers(configs),
		configs:    configs,
	}
}

func (api *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	var params request.SignupReq
	json.NewDecoder(r.Body).Decode(&params)

	//Check if the password are the same
	if params.Password != params.PasswordConfirmation {
		output["message"] = api.configs.Translations.Auth.Signup.Errors.PasswordDifferent
		response.JsonRes(w, output, http.StatusUnprocessableEntity)
		return
	}

	//Check if the user already exists
	_, err := api.repository.GetUserByEmail(params.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserGet) {
			output["message"] = api.configs.Translations.Auth.Signup.Errors.SaveUser
			response.JsonRes(w, output, http.StatusInternalServerError)
			return
		}
	} else {
		output["message"] = api.configs.Translations.Auth.Signup.Errors.AlreadyExists
		response.JsonRes(w, output, http.StatusUnprocessableEntity)
		return
	}

	//Hashed Password
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 8)
	if err != nil {
		output["message"] = api.configs.Translations.Auth.Signup.Errors.SaveUser
		response.JsonRes(w, output, http.StatusInternalServerError)
		return
	}

	//User Object
	user := types.UserDB{
		Name:     params.Name,
		Email:    params.Email,
		Password: string(encryptedPassword[:]),
	}

	//Save User
	result, err := api.repository.CreateUser(user)
	if err != nil {
		output["message"] = api.configs.Translations.Auth.Signup.Errors.SaveUser
		response.JsonRes(w, output, http.StatusInternalServerError)
		return
	}

	response.JsonRes(w, result, http.StatusCreated)
}
