package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(logger *slog.Logger, pool *pgxpool.Pool) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request)  {
		logger.Info("Request received")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		payload, _ := json.Marshal("Hello World")
		
		w.Write(append(payload, '\n'))
		logger.Info("Request closed")
	})
	
	router.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request)  {
		logger.Info("Request received")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		
		err := pool.Ping(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		payload, _ := json.Marshal("Pong")
		
		w.WriteHeader(http.StatusOK)
		w.Write(append(payload, '\n'))
		logger.Info("Request closed")
	})

	return router
}