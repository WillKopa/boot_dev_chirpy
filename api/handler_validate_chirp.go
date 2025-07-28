package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func validate_chirp(rw http.ResponseWriter, req *http.Request) {
	type request_params struct {
		Chirp string `json:"body"`
	}

	type response_params struct {
	    Cleaned_body 	string 		`json:"cleaned_body"`
	}
	response_body := response_params{}

	decoder := json.NewDecoder(req.Body)
	params := request_params{}
	err := decoder.Decode(&params)
	max_length := 140
	
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		respond_with_error(rw, http.StatusBadRequest, "parameters not valid")
		return
	} else if len(params.Chirp) > max_length  {
		respond_with_error(rw, http.StatusBadRequest, "Chirp is too long")
		return
	}

	response_body.Cleaned_body = remove_profane(params.Chirp)
	respond_with_json(rw, http.StatusOK, response_body)
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
