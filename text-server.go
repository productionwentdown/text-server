package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "%s", os.Getenv("TEXT"))
}

func main() {
	fmt.Println("Listening on port 8080...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
