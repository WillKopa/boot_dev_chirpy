package api

import (
	"net/http"
	"sync/atomic"

	"github.com/WillKopa/boot_dev_chirpy/constants"
)

type apiConfig struct {
	file_server_hits atomic.Int32
}

func Get_mux() *http.ServeMux {
	cfg := apiConfig{
		file_server_hits: atomic.Int32{},
	}

	server_mux := http.NewServeMux()
	server_mux.Handle("/app/", http.StripPrefix("/app", cfg.middleware_metrics_inc(http.FileServer(http.Dir(constants.ROOT_PATH)))))
	server_mux.HandleFunc("GET /metrics", cfg.server_metrics)
	server_mux.HandleFunc("GET /healthz", is_service_available)
	server_mux.HandleFunc("POST /reset", cfg.reset_metrics)

	return server_mux
}
