package httpc

import (
	"context"
	"github.com/bellingham07/gozero-tools/errorx"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// Response 基本响应结构
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// RespSuccess 正常响应的返回
func RespSuccess(ctx context.Context, w http.ResponseWriter, resp interface{}) {
	var body Response
	body.Code = 200
	body.Msg = "success"
	body.Data = resp
	// 将预定好的结构重新封装给，然后交给go-zero处理
	httpx.OkJsonCtx(ctx, w, body)
}

// RespError 错误响应
func RespError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	var (
		code     = http.StatusInternalServerError
		res      = Response{Code: code, Msg: "服务器开小差啦，稍后再来试一试"}
		metadata any
		appType  string
	)
	switch err.(type) {
	case *errorx.CodeError:
		customErr := errorx.From(err)
		res.Code = customErr.Code
		res.Msg = customErr.Msg
		code = customErr.Code
		appType = customErr.BizType
		metadata = customErr.Metadata
	}

	logc.Errorw(ctx, "",
		logc.Field("err", err),
		logc.Field("code", code),
		logc.Field("type", appType),
		logc.Field("metadata", metadata),
		logc.Field("method", r.Method),
		logc.Field("path", r.URL.Path),
	)

	// 业务错误，http响应状态码为200，业务错误码在body中
	httpx.OkJsonCtx(ctx, w, res)
}
func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusUnauthorized, &Response{http.StatusUnauthorized, "鉴权失败", nil})
}
