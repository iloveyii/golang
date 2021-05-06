package main

import(
	"fmt"
	"net/http"
)

// Render index
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my site")
}
// Render about
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about page")
}


func main() {
	fileServer := http.FileServer(http.Dir("./static")) 
	http.Handle("/*", fileServer) 

	http.HandleFunc("/index", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server listening on port 3300")
	http.ListenAndServe(":3300", fileServer) 
}