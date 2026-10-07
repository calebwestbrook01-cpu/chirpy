package main

import (
	"fmt"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	serveMux.HandleFunc("/healthz", handlerReadiness)
	myServer := http.Server{Addr: ":8080", Handler: serveMux}
	err := myServer.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
