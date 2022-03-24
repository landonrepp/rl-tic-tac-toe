package api

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
