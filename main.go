package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
	"database/sql"
	_ "github.com/lib/pq"
)

type person struct {
	firstName string
	lastName  string
}

func (p person) speak(phrase string) {
	fmt.Println(p.firstName, p.lastName, "say", phrase)
}

func HomeHandler(w http.ResponseWriter, r*http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Category: %v\n", vars["category"])
	w.Write([]byte("Gorilla!\n"))
}

func HealthCheckHandler(w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// w.WriteString(w, `{"alive": true}`)
	json.NewEncoder(w).Encode(map[string]bool{"alive": true})
}

func GetBookingListHandler(w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	connStr := "user=booyokkk dbname=first sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	fmt.Println("Server runing on port 8000")
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/health-check", HealthCheckHandler).Methods("GET")
	r.HandleFunc("/bookings", GetBookingListHandler).Methods("GET")
	// http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
