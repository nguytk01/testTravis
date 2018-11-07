package main

import (
	"net/http"
	"testTravis/server/businesslogic"
)

func main() {
	http.HandleFunc("/", businesslogic.Hey)
	http.HandleFunc("/newUnauthorizedSession", businesslogic.NewUnauthorizedSessionHandler)
	http.HandleFunc("/newAuthorizedSession", businesslogic.NewAuthorizedSessionHandler)
	http.HandleFunc("/new_user", businesslogic.NewUserHandler)
  http.HandleFunc("/Login", businesslogic.LogUserIn)
	http.ListenAndServe(":8000", nil);
}