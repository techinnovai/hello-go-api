package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message": "Hello World"}`)
	})

	fmt.Println("Running on http://localhost:8080/hello")
	http.ListenAndServe(":8080", nil)
}
