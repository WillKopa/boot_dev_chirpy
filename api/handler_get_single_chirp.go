package api

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) get_single_chirp(rw http.ResponseWriter, req *http.Request) {
	chirp_id, err := uuid.Parse(req.PathValue("chirpID"))
	if err != nil {
		log.Printf("Error parsing uuid: %s", err)
		respond_with_error(rw, http.StatusBadRequest, "passed chirpID is not a uuid")
	}

	chirp, err := cfg.db_queries.GetSingleChirp(req.Context(), chirp_id)

	if err != nil {
		log.Printf("Error getting chirp from database: %s", err)
		respond_with_error(rw, http.StatusNotFound, "unable to get chirp")
	}

	respond_with_json(rw, http.StatusOK, Chirp(chirp))
}