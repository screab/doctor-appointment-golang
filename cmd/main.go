package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=username dbname=appointments sslmode=test"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS appointment (id SERIAL PRIMARY KEY, title TEXT, date TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/appointments", appointmentsHandler)
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func appointmentsHandler(w http.ResponseWriter, r *http.Request) {
	// Logic to handle appointments goes here
	fmt.Fprintln(w, "Appointments Endpoint")
}