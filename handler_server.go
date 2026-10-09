package main

import (
	"fmt"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	x := cfg.fileserverHits.Load()
	message := fmt.Sprintf(
		`<html>
  		<body>
    		<h1>Welcome, Chirpy Admin</h1>
    		<p>Chirpy has been visited %d times!</p>
  		</body>
	</html>`,
		x)
	w.Write([]byte(message))
}

func (cfg *apiConfig) handlerHitsReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
}
