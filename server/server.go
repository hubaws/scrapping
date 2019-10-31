package server

import (
	"log"

	"github.com/gorilla/mux"
)

func StartServer(){
	log.Println("Start App Server")
	r := mux.NewRouter()
	log.Println(r)
}