package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerServ(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodGet:

	case http.MethodPost:

	case http.MethodPut:

	case http.MethodPatch:

	case http.MethodDelete:

	default:
		println("default")
	}

}

func Server() {
	fmt.Println("Server start...")

	router := mux.NewRouter()

	router.HandleFunc("/", HandlerServ)
	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}
