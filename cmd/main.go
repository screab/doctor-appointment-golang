package main

import (
    "log"
    "net/http"
    "screab/doctor-appointment-golang/api"
)

func main() {
    router := api.SetupRouter()
    log.Println("Server is running on port 8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}
