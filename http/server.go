package main


import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})

	fmt.Println("Server is running at http://localhost:6767")
	http.ListenAndServe(":6767", nil)
}