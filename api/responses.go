package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func respond_with_error(rw http.ResponseWriter, code int, msg string) {
	type respond_with_error struct {
		Error  string `json:"error"`
	}
	respond_with_json(rw, code, respond_with_error{
		Error: msg,
	} )
}

func respond_with_json(rw http.ResponseWriter, code int, payload any) {
	rw.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error creating respone: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(code)
	rw.Write(dat)
}