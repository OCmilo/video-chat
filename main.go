package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	roomUrl := uuid.New().String()
	http.Redirect(w, r, roomUrl, 301)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.Path("/{roomId}").Handler(http.FileServer(http.Dir("./front/public/index.html")))

	log.Fatal(http.ListenAndServe(":8081", r))
}
