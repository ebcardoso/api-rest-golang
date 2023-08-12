package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/ebcardoso/api-rest-golang/config"
	"github.com/ebcardoso/api-rest-golang/utils"
	"github.com/ebcardoso/api-rest-golang/utils/response"
)

type RouterProtector struct {
	configs *config.Config
}

func NewRouterProtector(configs *config.Config) *RouterProtector {
	return &RouterProtector{
		configs: configs,
	}
}

func (rp *RouterProtector) Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		bearer := strings.TrimSpace(strings.Replace(authorization, "Bearer", "", 1))

		if bearer == "" {
			output := make(map[string]interface{})
			output["message"] = rp.configs.Translations.Auth.Protector.TokenRequired
			response.JsonRes(w, output, http.StatusBadRequest)
			return
		}

		token, payload, err := utils.DecodeJWT(bearer, rp.configs.Env.JWT_KEY)
		if err != nil || !token.Valid {
			output := make(map[string]interface{})
			output["message"] = rp.configs.Translations.Auth.Protector.NotAuth
			response.JsonRes(w, output, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "id", payload.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
