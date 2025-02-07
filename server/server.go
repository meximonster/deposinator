package server

import (
	"net/http"
	"time"

	"github.com/deposinator/controller"
	"github.com/gorilla/mux"
)

func Run() error {

	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         ":5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/signup", controller.Signup).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("POST")

	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
