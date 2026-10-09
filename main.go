package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

func main() {
	serveMux := http.NewServeMux()
	apiCfg := apiConfig{}
	serveMux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	serveMux.HandleFunc("GET /api/healthz", handlerReadiness)
	serveMux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	serveMux.HandleFunc("POST /admin/reset", apiCfg.handlerHitsReset)
	myServer := http.Server{Addr: ":8080", Handler: serveMux}
	err := myServer.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
