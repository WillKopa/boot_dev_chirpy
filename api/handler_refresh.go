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
	
	refresh_token, err := cfg.db_queries.GetRefreshToken(req.Context(), token)
	if err != nil {
		respond_with_error(rw, http.StatusNotFound, "refresh token not found")
		return
	}

	if refresh_token.ExpiresAt.Before(time.Now()) {
		respond_with_error(rw, http.StatusUnauthorized, "refresh token has expired")
		return
	}

	if refresh_token.RevokedAt.Valid {
		respond_with_error(rw, http.StatusUnauthorized, "refresh token access revoked, please login again")
		return
	}

	jwt_token, err := auth.MakeJWT(refresh_token.UserID, cfg.secret, time.Hour * 1)

	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "error creating jwt token")
		return
	}

	respond_with_json(rw, http.StatusOK, RefreshToken{Token: jwt_token})
}