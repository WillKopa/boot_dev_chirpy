package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func validate_chirp(rw http.ResponseWriter, req *http.Request) {
	type request_params struct {
		Chirp string `json:"body"`
	}

	type response_params struct {
		Err 	string 		`json:"error"`
	    Valid 	bool 		`json:"valid"`
	}
	response_body := response_params{}

	decoder := json.NewDecoder(req.Body)
	params := request_params{}
	err := decoder.Decode(&params)
	max_length := 140
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		response_body.Err = "Something went wrong"
		response_body.Valid = false
		rw.WriteHeader(http.StatusBadRequest)
	} else if len(params.Chirp) > max_length  {
		response_body.Err = "Chirp is too long"
		response_body.Valid = false
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		response_body.Valid = true
		rw.WriteHeader(http.StatusOK)
	}
	
	dat, err := json.Marshal(response_body)

	if err != nil {
		log.Printf("Error creating respone: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(dat)
}
