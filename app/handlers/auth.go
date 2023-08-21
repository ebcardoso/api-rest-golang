package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ebcardoso/api-rest-golang/app/repository"
	"github.com/ebcardoso/api-rest-golang/app/types"
	"github.com/ebcardoso/api-rest-golang/config"
	"github.com/ebcardoso/api-rest-golang/utils"
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

func (api *Auth) Signin(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	var params request.SigninReq
	json.NewDecoder(r.Body).Decode(&params)

	//Find user on BD
	user, err := api.repository.GetUserByEmail(params.Email)
	if err != nil {
		var status int
		if errors.Is(err, repository.ErrUserNotFound) {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
		output["message"] = err.Error()
		response.JsonRes(w, output, status)
		return
	}

	//Check if the user is blocked
	if user.IsBlocked {
		output["message"] = api.configs.Translations.Auth.Signin.UserBlocked
		response.JsonRes(w, output, http.StatusUnauthorized)
		return
	}

	//Call check password
	checkPassword := types.CheckPassword(user, params.Password)
	if !checkPassword {
		output["message"] = api.configs.Translations.Auth.Signin.Invalid
		response.JsonRes(w, output, http.StatusUnauthorized)
		return
	}

	//To do: Generate token jwt
	token, err := utils.EncodeJWT(user.ID.Hex(), api.configs.Env.JWT_KEY)
	if err != nil {
		return
	}
	output["message"] = api.configs.Translations.Auth.Signin.Success
	output["accessToken"] = token

	response.JsonRes(w, output, http.StatusOK)
}

func (api *Auth) CheckToken(w http.ResponseWriter, r *http.Request) {

}
