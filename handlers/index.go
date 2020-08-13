package handlers

import (
	"github.com/peimengc/chitchat/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads()

	if err == nil {
		_, err = session(w, r)

		if err != nil {
			generateHTML(w, threads, "layout", "navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "auth.navbar", "index")
		}
	}
}

func Err(w http.ResponseWriter, r *http.Request) {
	queryVal := r.URL.Query()

	msg := queryVal.Get("msg")

	if _, err := session(w, r); err != nil {
		generateHTML(w, msg, "layout", "navbar", "error")
	} else {
		generateHTML(w, msg, "layout", "auth.navbar", "error")
	}
}
