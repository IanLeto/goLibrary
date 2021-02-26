package httpServerInDocker

import (
	"goLibrary/utils"
	"net/http"
)

type HttpServiceConfig struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Port string `json:"port"`
}

type HttpService struct {
	Config HttpServiceConfig
}

func NewHttpService() *HttpService {

	return nil
}

func Run() {
	http.HandleFunc("/ping", func(responseWriter http.ResponseWriter, request *http.Request) {
		_, err := responseWriter.Write([]byte(`pong`))
		utils.CheckPanicError(err)
	})
}
