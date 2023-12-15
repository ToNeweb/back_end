package main

import (
	"log"
	"net/http"
	"server04/config"
	"server04/middleware"
	"time"

	"server04/router"

	"github.com/gorilla/mux"
)

func main() {

	//initiate Ent Client
	client, err := config.NewEntClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}

	// //set the client to the variable defined in package config
	// //this will enable the client intance to be accessed anywhere through the accessor which is a function
	// //named GetClient
	config.SetClient(client)

	// //initiate router and register all the route
	r := mux.NewRouter()
	r.Use(middleware.Header)
	router.RegisterRouter(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8019",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port 8019")
	log.Fatal(srv.ListenAndServe())
}

//http://127.0.0.1:8019/user/
