package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/WillKopa/boot_dev_chirpy/constants"
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


func validate_chirp(chirp string) (string, error) {
	if len(chirp) > constants.MAX_CHIRP_LENGTH  {
		return "", fmt.Errorf("chirp is too long. must be less than 141 characters")
	}

	clean_chirp := remove_profane(chirp)
	return clean_chirp, nil

}

func remove_profane(text string) string {
	profane_words := map[string]struct{} {
		"kerfuffle": {},
		"sharbert": {},
		"fornax": {},
	}
	cleaned_text := []string{}
	for _, word := range(strings.Split(text, " ")) {
		_, is_profane := profane_words[strings.ToLower(word)]
		if is_profane {
			cleaned_text = append(cleaned_text, "****")
		} else {
			cleaned_text = append(cleaned_text, word)
		}
	}

	return strings.Join(cleaned_text, " ")
}
