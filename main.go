package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening at 8080...")
	log.Fatal(http.ListenAndServe(":8080", logger(http.FileServer(http.Dir("app")))))
}

func logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s %s\n", r.RemoteAddr, r.Method, r.RequestURI)
		handler.ServeHTTP(w, r)
	})
}
