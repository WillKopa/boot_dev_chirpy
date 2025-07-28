package main

import (
	"log"
	"net/http"

	"github.com/WillKopa/boot_dev_chirpy/api"
	"github.com/WillKopa/boot_dev_chirpy/constants"
)

func main() {
	server_mux := api.Get_mux()
	server := &http.Server{
		Handler: server_mux,
		Addr:    "localhost:" + constants.PORT,
	}
	log.Printf("Serving files from %s on port: %s\n", constants.ROOT_PATH, constants.PORT)
	log.Fatal(server.ListenAndServe())
}


