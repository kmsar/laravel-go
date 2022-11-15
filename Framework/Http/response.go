package Http

import (
	"errors"
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IHttp"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Logs"
	"go/types"
	"net/http"
	"os"
)

var (
	FileTypeError = errors.New("Wrong file parameter type")
)

// HandleResponse handles the controller function's response
func HandleResponse(response interface{}, ctx IHttp.IHttpRequest) {
	if response == nil {
		return
	}
	switch res := response.(type) {
	case error:
		Logs.WithError(ctx.String(http.StatusInternalServerError, res.Error())).Debug("response error")
	case string:
		Logs.WithError(ctx.String(http.StatusOK, res)).Debug("response error")
	case fmt.Stringer:
		Logs.WithError(ctx.String(http.StatusOK, res.String())).Debug("response error")
	case Support.Json:
		Logs.WithError(ctx.String(http.StatusOK, res.ToJson())).Debug("response error")
	case IHttp.IHttpResponse:
		Logs.WithError(res.Response(ctx)).Debug("response error")
	case types.Nil:
		return
	default:
		Logs.WithError(ctx.JSON(200, res)).Debug("response json error")
	}

}

type Response struct {
	status   int
	Json     interface{}
	String   string
	FilePath string
	File     *os.File
}

func StringResponse(str string, code ...int) IHttp.IHttpResponse {
	status := 200
	if len(code) > 0 {
		status = code[0]
	}
	return Response{
		status: status,
		String: str,
	}
}

func JsonResponse(json interface{}, code ...int) IHttp.IHttpResponse {
	status := 200
	if len(code) > 0 {
		status = code[0]
	}
	return Response{
		status: status,
		Json:   json,
	}
}

// FileResponse response file
func FileResponse(file interface{}) IHttp.IHttpResponse {
	switch f := file.(type) {
	case *os.File:
		return Response{File: f}
	case string:
		return Response{FilePath: f}
	default:
		panic(FileTypeError)
	}
}

func (res Response) Status() int {
	return res.status
}

func (res Response) Response(ctx IHttp.HttpContext) error {
	if res.Json != nil {
		return ctx.JSON(res.Status(), res.Json)
	}
	if res.FilePath != "" {
		return ctx.File(res.FilePath)
	}
	if res.File != nil {
		return ctx.File(res.File.Name())
	}

	return ctx.String(res.Status(), res.String)
}
