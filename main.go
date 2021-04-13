package main

import (
	"goLibrary/utils"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", func(responseWriter http.ResponseWriter, request *http.Request) {
		_, err := responseWriter.Write([]byte(`pong`))
		utils.CheckPanicError(err)
	})

}
