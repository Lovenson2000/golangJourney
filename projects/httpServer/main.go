package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Serving static files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/about.html")
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/products.html")
	})

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/services.html")
	})

	// Serve static files from the static folder for all other requests
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	port := ":3000"
	fmt.Println("Server is running on port" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
