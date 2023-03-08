package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func handler() {
	muxRoutes := mux.NewRouter().StrictSlash(true)

	// WeekDays
	muxRoutes.HandleFunc("/weekdays/{USER_ID}", getWeekDays).Methods("GET")
	muxRoutes.HandleFunc("/weekdays", createWeekDays).Methods("POST")
	muxRoutes.HandleFunc("/weekdays/{USER_ID}", deleteWeekDays).Methods("DELETE")

	// Weekend
	muxRoutes.HandleFunc("/weekend/{USER_ID}", getWeekEnd).Methods("GET")
	muxRoutes.HandleFunc("/weekend", createWeekEnd).Methods("POST")
	muxRoutes.HandleFunc("/weekend/{USER_ID}", deleteWeekEnd).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"DELETE", "GET", "POST"},
	})

	handler := c.Handler(muxRoutes)
	log.Println("server started on 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
