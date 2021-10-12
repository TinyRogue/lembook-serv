package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ":(")
}

func main() {
	http.HandleFunc("/", handler)
	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8080"
	}
	log.Fatal(http.ListenAndServe(":" + PORT, nil))
}
