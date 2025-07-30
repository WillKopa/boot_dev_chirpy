package api

import (
	"net/http"

	"github.com/WillKopa/boot_dev_chirpy/internal/auth"
)

func (cfg *apiConfig) revoke(rw http.ResponseWriter, req *http.Request) {
	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respond_with_error(rw, http.StatusBadRequest, "no refresh token in header")
		return
	}
	err = cfg.db_queries.RevokeRefreshToken(req.Context(), token)
	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "error revoking token")
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}