package api

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) get_chirps(rw http.ResponseWriter, req *http.Request) {
	author_id, _ := uuid.Parse(req.URL.Query().Get("author_id"))

	db_chirps, err := cfg.db_queries.GetChirps(req.Context(), author_id)

	if err != nil {
		log.Printf("Error getting all chirps from database: %s", err)
		respond_with_error(rw, http.StatusInternalServerError, "unable to read from databse")
	}

	chirps := make([]Chirp, len(db_chirps))
	for i, chirp := range db_chirps {
		chirps[i] = Chirp(chirp)
	}

	respond_with_json(rw, http.StatusOK, chirps)
}