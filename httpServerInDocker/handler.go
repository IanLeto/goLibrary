package httpServerInDocker

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"goLibrary/utils"
)

type RequestInfoDemo struct {
	Id  string `json:"id"`
	key string
}

type BaseResponseInfo struct {
	ErrCode int
	Data    interface{}
	Message string
}

type ResponseInfoDemo struct {
}

func Response(ctx *fasthttp.RequestCtx, data interface{}, err error) {
	if err != nil {
		data, err := json.Marshal(data)
		utils.NoErr(err)
		ctx.Response.SetBody(data)
	}
	ctx.Response.SetBody(nil)
}

func HandlerHelloWorld(ctx *fasthttp.RequestCtx) {
	ctx.Write([]byte(`hello world`))
}

func GetArgs(ctx *fasthttp.RequestCtx, info interface{}) error {
	fmt.Println(ctx.Request.Body())
	return json.Unmarshal(ctx.QueryArgs().QueryString(), info)
}

func HandlerGetDemo(ctx *fasthttp.RequestCtx) {
	//var info = &RequestInfoDemo{}

	//utils.NoErr(GetArgs(ctx, info))
	fmt.Println(string(ctx.QueryArgs().Peek("key")))
	res := ResponseInfoDemo{}
	Response(ctx, res, nil)
	//helper := utils.NewJsonHelperByte(ctx.Request.Body())
	//v := helper.Get("key")
	//
	//if v != "" {
	//	ctx.Response.SetStatusCode(http.StatusOK)
	//	v, err := json.Marshal(ResponseInfo{
	//		ErrCode: 200,
	//		Data:    nil,
	//		Message: "get demo",
	//	})
	//	utils.NoErr(err)
	//	ctx.Response.SetBody(v)
	//}

}
