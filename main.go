package main

import (
	"net/http"
)

func main() {

	//Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	//Star the server
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Word"))

}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pagina de Contact"))

}
