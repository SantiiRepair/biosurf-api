package auth

import (
	mux "github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", HandleRegister).Methods("POST")
	r.HandleFunc("/login", HandleLogin).Methods("POST")

	http.ListenAndServe(":8080", r)
}
