package utils

import (
	"fmt"
	"log"
	"net/http"
)

func NewHttpService() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello world: port:8888")
	})
	log.Fatal(http.ListenAndServe(":8888", nil))
}
