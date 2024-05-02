package main

import (
	"io"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// output some minimal but helpful text in simple html
	io.WriteString(w, "Auth service home.")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// output some minimal but helpful text in simple html
	io.WriteString(w, "Health of auth service is OK")
}

//func postHttpBasicHandler(w http.ResponseWriter, r *http.Request) {
//	rUser, rPass, ok := r.BasicAuth()
//	if !ok {
//		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
//		http.Error(w, "Unauthorized", http.StatusUnauthorized)
//		return
//	}
//
//	// Validate the user and pass.
//	// Note: You should have your own ValidateUser function which may likely use a database.
//	validUser, err := ValidateUser(rUser, rPass)
//	if err != nil || !validUser {
//		http.Error(w, "Unauthorized", http.StatusUnauthorized)
//		return
//	}
//
//	io.WriteString(w, "Authenticated with HTTP Basic")
//}
