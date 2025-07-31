package api

import (
	"net/http"

	"github.com/WillKopa/boot_dev_chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) make_chirpy_red(rw http.ResponseWriter, req *http.Request) {
	type MakeRedRequest struct {
		Event 		string						`json:"event"`
		Data		struct {
			UserID		uuid.UUID		`json:"user_id"`
		}		`json:"data"`
	}

	red_request, err := unmarshal_json[MakeRedRequest](req)
	if err != nil {
		respond_with_error(rw, http.StatusBadRequest, "unable to parse request")
		return
	}

	if red_request.Event != "user.upgraded" {
		rw.WriteHeader(http.StatusNoContent)
		return
	}

	user, err := cfg.db_queries.MakeChirpyRed(req.Context(), red_request.Data.UserID)
	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "error updating user")
		return
	}

	empty_user := database.User{}
	if user == empty_user {
		respond_with_error(rw, http.StatusNotFound, "no user found for given id")
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}