package httpServerInDocker

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
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

// fastHttp 版本
// 服务端
func FastHttpDemo() {
	r := router.New()
	r.GET("/", HandlerHelloWorld)
	r.GET("/value", HandlerGetDemo)
	utils.NoErr(fasthttp.ListenAndServe(":8001", r.Handler))
}
