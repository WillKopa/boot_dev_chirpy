package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/WillKopa/boot_dev_chirpy/internal/database"
)


func (cfg *apiConfig) create_chirp(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	chirp := Chirp{}
	err := decoder.Decode(&chirp)
	
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		respond_with_error(rw, http.StatusInternalServerError, "error parsing json from request")
	}

	chirp.Body, err = validate_chirp(chirp.Body)
	if err != nil {
		respond_with_error(rw, http.StatusBadRequest, err.Error())
	}

	db_params := database.CreateChirpParams{}
	db_params.Body = chirp.Body
	db_params.UserID = chirp.UserID

	db_chirp, err := cfg.db_queries.CreateChirp(req.Context(), db_params)
	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "Sever error")
		return
	}

	respond_with_json(rw, http.StatusCreated, Chirp(db_chirp))
	
}