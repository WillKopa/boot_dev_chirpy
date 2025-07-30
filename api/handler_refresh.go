package api

import (
	"net/http"
	"time"

	"github.com/WillKopa/boot_dev_chirpy/internal/auth"
)

func (cfg *apiConfig) refresh(rw http.ResponseWriter, req *http.Request) {
	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respond_with_error(rw, http.StatusBadRequest, "refresh token missing from header")
		return
	}
	
	user, err := cfg.db_queries.GetRefreshTokenUser(req.Context(), token)

	if err != nil {
		respond_with_error(rw, http.StatusUnauthorized, "no refresh token available")
		return
	}

	jwt_token, err := auth.MakeJWT(user.ID, cfg.secret, time.Hour)

	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "error creating jwt token")
		return
	}

	respond_with_json(rw, http.StatusOK, RefreshToken{Token: jwt_token})
}