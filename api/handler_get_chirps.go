package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/WillKopa/boot_dev_chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) get_chirps(rw http.ResponseWriter, req *http.Request) {
	author_id, _ := uuid.Parse(req.URL.Query().Get("author_id"))
	sort_by := req.URL.Query().Get("sort")
	var db_chirps []database.Chirp
	var err error
	if strings.ToLower(sort_by) == "desc" {
		db_chirps, err = cfg.db_queries.GetChirpsDESC(req.Context(), author_id)
	} else {
		db_chirps, err = cfg.db_queries.GetChirpsASC(req.Context(), author_id)
	}

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