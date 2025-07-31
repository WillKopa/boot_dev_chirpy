package api

import (
	"net/http"

	"github.com/WillKopa/boot_dev_chirpy/internal/auth"
	"github.com/WillKopa/boot_dev_chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) delete_chirp(rw http.ResponseWriter, req *http.Request) {
	auth_token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respond_with_error(rw, http.StatusUnauthorized, "Missing JWT in header")
		return
	}

	token_user_id, err := auth.ValidateJWT(auth_token, cfg.secret)
	if err != nil {
		respond_with_error(rw, http.StatusUnauthorized, "YOU SHALL NOT DELETE!")
		return
	}

	chirp_id, err := uuid.Parse(req.PathValue("chirpID"))
	if err != nil {
		respond_with_error(rw, http.StatusBadRequest, "invalid chirp id")
		return
	}

	chirp, err := cfg.db_queries.GetSingleChirp(req.Context(), chirp_id)
	if err != nil {
		respond_with_error(rw, http.StatusNotFound, "No chirp with that ID")
		return
	}

	if chirp.UserID != token_user_id {
		respond_with_error(rw, http.StatusForbidden, "You cannot delete someone elses chirp")
		return
	}

	params := database.DeleteChirpParams{
		ID: chirp_id,
		UserID: token_user_id,
	}
	err = cfg.db_queries.DeleteChirp(req.Context(), params)
	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "Unable to delete chirp")
		return
	}


	rw.WriteHeader(http.StatusNoContent)
}