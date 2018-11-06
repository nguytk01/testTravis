package main

import (
	"net/http"
	"acs560_course_project/server/businesslogic/"
)

func main() {
	http.HandleFunc("/", Hey)
	http.HandleFunc("/newUnauthorizedSession", newUnauthorizedSessionHandler)
	http.HandleFunc("/newAuthorizedSession", newAuthorizedSessionHandler)
	http.HandleFunc("/new_user", newUserHandler)
	http.ListenAndServe(":8000", nil);
}