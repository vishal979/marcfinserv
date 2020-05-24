package main

import (
	"log"
	"marcfinserv/filehandler"
	"marcfinserv/routes"
	"marcfinserv/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func serve(r *mux.Router) {
	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}
	log.Println("Server running on port 8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Println("Error in setting up server with the error", err)
	}

}

func main() {
	filehandler.Open()
	defer filehandler.CloseFile()
	r := routes.Init()
	routes.Handle(r)
	utils.LoadTemplates("templates/*.html")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	serve(r)
}
