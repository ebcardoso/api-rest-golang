package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ebcardoso/api-rest-golang/app/repository"
	"github.com/ebcardoso/api-rest-golang/app/types"
	"github.com/ebcardoso/api-rest-golang/config"
	"github.com/ebcardoso/api-rest-golang/utils/response"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	repository *repository.Users
	configs    *config.Config
}

func NewUsers(configs *config.Config) *Users {
	return &Users{
		repository: repository.NewRepositoryUsers(configs),
		configs:    configs,
	}
}

func (api *Users) GetListUsers(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	items, err := api.repository.ListUsers()
	if err != nil {
		output["message"] = err.Error()
		response.JsonRes(w, output, http.StatusInternalServerError)
		return
	}

	output["content"] = items
	response.JsonRes(w, output, http.StatusOK)
}

func (api *Users) GetUserByID(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	id := chi.URLParam(r, "id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		output["message"] = api.configs.Translations.Errors.ParseId
		response.JsonRes(w, output, http.StatusBadRequest)
		return
	}

	result, err := api.repository.GetUserByID(objID)
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
	response.JsonRes(w, result, http.StatusOK)
}

func (api *Users) UpdateUser(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	id := chi.URLParam(r, "id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		output["message"] = api.configs.Translations.Errors.ParseId
		response.JsonRes(w, output, http.StatusBadRequest)
		return
	}

	var user types.UserDB
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		output["message"] = "Invalid Params"
		response.JsonRes(w, output, http.StatusBadRequest)
		return
	}

	err = api.repository.UpdateUser(objID, user)
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

	user.ID = objID
	response.JsonRes(w, types.MapUserDB(user), http.StatusOK)
}

func (api *Users) DestroyUser(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	id := chi.URLParam(r, "id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		output["message"] = api.configs.Translations.Errors.ParseId
		response.JsonRes(w, output, http.StatusBadRequest)
		return
	}

	err = api.repository.DestroyUser(objID)
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

	output["message"] = api.configs.Translations.Users.Destroy.Success
	response.JsonRes(w, output, http.StatusOK)
}

func (api *Users) Block(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	//Parsing ID
	id := chi.URLParam(r, "id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		output["message"] = api.configs.Translations.Errors.ParseId
		response.JsonRes(w, output, http.StatusBadRequest)
		return
	}

	//Persisting the Block
	err = api.repository.BlockOrUnlockUser(objID, true)
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

	//Success Response
	output["message"] = api.configs.Translations.Users.Block.Success
	response.JsonRes(w, output, http.StatusOK)
}

func (api *Users) Unblock(w http.ResponseWriter, r *http.Request) {
	output := make(map[string]interface{})

	//Parsing ID
	id := chi.URLParam(r, "id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		output["message"] = api.configs.Translations.Errors.ParseId
		response.JsonRes(w, output, http.StatusBadRequest)
		return
	}

	//Persisting the Block
	err = api.repository.BlockOrUnlockUser(objID, false)
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

	//Success Response
	output["message"] = api.configs.Translations.Users.Unblock.Success
	response.JsonRes(w, output, http.StatusOK)
}
