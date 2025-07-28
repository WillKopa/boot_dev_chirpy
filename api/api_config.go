package api

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/WillKopa/boot_dev_chirpy/constants"
	"github.com/WillKopa/boot_dev_chirpy/internal/database"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	file_server_hits atomic.Int32
	db_queries 		*database.Queries
}

func Get_mux() *http.ServeMux {
	db_queries := connect_db()
	cfg := apiConfig{
		file_server_hits: atomic.Int32{},
		db_queries: db_queries,
	}

	server_mux := http.NewServeMux()
	server_mux.Handle("/app/", http.StripPrefix("/app", cfg.middleware_metrics_inc(http.FileServer(http.Dir(constants.ROOT_PATH)))))

	// admin
	server_mux.HandleFunc("GET /admin/metrics", cfg.server_metrics)
	server_mux.HandleFunc("POST /admin/reset", cfg.reset_metrics)

	// api
	server_mux.HandleFunc("GET /api/healthz", is_service_available)
	server_mux.HandleFunc("POST /api/validate_chirp", validate_chirp)

	return server_mux
}


func connect_db() *database.Queries {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("error opening database: %s", err)
	}

	return database.New(db)
}
