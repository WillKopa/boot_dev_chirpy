package api

import (
	"fmt"
	"net/http"
)

func is_service_available(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(http.StatusText(http.StatusOK)))
}

func (cfg *apiConfig) server_metrics(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "text/html; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	metrics_text := 
	`<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
	</html>`
	text := fmt.Sprintf(metrics_text, cfg.file_server_hits.Load())
	rw.Write([]byte(text))
}

func (cfg *apiConfig) reset_everything(rw http.ResponseWriter, req *http.Request) {
	if cfg.platform != "dev" {
		respond_with_error(rw, http.StatusForbidden, "You shall not pass")
		return
	}
	cfg.db_queries.DeleteUsers(req.Context())
	rw.Header().Add("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	cfg.file_server_hits.Store(0)
	text := fmt.Sprintf("Hits Reset to: %v\nAll users deleted", cfg.file_server_hits.Load())
	rw.Write([]byte(text))
}

func (cfg *apiConfig) middleware_metrics_inc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		cfg.file_server_hits.Add(1)
		next.ServeHTTP(wr, req)
	})
}
