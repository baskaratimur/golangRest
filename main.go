package main

import (
	"fmt"
	"net/http"
	"project_baskara/request"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/register", request.RegisFunc).Methods("POST")
	r.HandleFunc("/findall", request.Findall).Methods("GET")
	r.HandleFunc("/findid", request.Findid).Methods("GET")
	r.HandleFunc("/update", request.Update).Methods("POST")
	r.HandleFunc("/delete", request.Delete).Methods("DELETE")
	fmt.Println("server start:8080")
	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", nil)
}
