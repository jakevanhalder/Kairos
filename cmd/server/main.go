package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Kairos is running")
}

func main() {
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}