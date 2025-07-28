package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	const root_path = "."

	server_mux := http.NewServeMux()
	server_mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(root_path))))
	server_mux.HandleFunc("/healthz", is_service_available)

	server := &http.Server{
		Handler: server_mux,
		Addr:    ":" + port,
	}
	log.Printf("Serving files from %s on port: %s\n", root_path, port)
	log.Fatal(server.ListenAndServe())
}

func is_service_available(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(http.StatusText(http.StatusOK)))
}
