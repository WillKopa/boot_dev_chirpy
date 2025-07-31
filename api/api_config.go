package api

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/WillKopa/boot_dev_chirpy/constants"
	"github.com/WillKopa/boot_dev_chirpy/internal/database"
)

type apiConfig struct {
	file_server_hits 	atomic.Int32
	db_queries 			*database.Queries
	platform			string
	secret				string
}

func Get_mux() *http.ServeMux {
	db_queries := connect_db()
	cfg := apiConfig{
		file_server_hits: 	atomic.Int32{},
		db_queries:		 	db_queries,
		platform: 			os.Getenv("PLATFORM"),
		secret: 			os.Getenv("SECRET"),
	}

	server_mux := http.NewServeMux()
	server_mux.Handle("/app/", http.StripPrefix("/app", cfg.middleware_metrics_inc(http.FileServer(http.Dir(constants.ROOT_PATH)))))

	// admin
	server_mux.HandleFunc("GET /admin/metrics", cfg.server_metrics)
	server_mux.HandleFunc("POST /admin/reset", cfg.reset_everything)

	// api
	server_mux.HandleFunc("DELETE /api/chirps/{chirpID}", cfg.delete_chirp)
	server_mux.HandleFunc("GET /api/chirps", cfg.get_chirps)
	server_mux.HandleFunc("GET /api/chirps/{chirpID}", cfg.get_single_chirp)
	server_mux.HandleFunc("POST /api/chirps", cfg.create_chirp)
	server_mux.HandleFunc("GET /api/healthz", is_service_available)
	server_mux.HandleFunc("POST /api/login", cfg.login)
	server_mux.HandleFunc("POST /api/refresh", cfg.refresh)
	server_mux.HandleFunc("POST /api/revoke", cfg.revoke)
	server_mux.HandleFunc("POST /api/users", cfg.create_user)
	server_mux.HandleFunc("PUT /api/users", cfg.update_user)


	// webhooks
	server_mux.HandleFunc("POST /api/polka/webhooks", cfg.make_chirpy_red)

	return server_mux
}


func connect_db() *database.Queries {
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("error opening database: %s", err)
	}

	return database.New(db)
}
