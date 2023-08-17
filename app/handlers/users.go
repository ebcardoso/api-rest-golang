package handlers

import (
	"errors"
	"net/http"

	"github.com/ebcardoso/api-rest-golang/app/repository"
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
