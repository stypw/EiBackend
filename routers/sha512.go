package routers

import (
	"main/common/df"
	"main/common/mw"
	"main/common/rt"
	"main/common/tl"
	"net/http"
)

type sha512Router struct{}

func (router *sha512Router) OnRequest(r *http.Request, msg *mw.Message, extra mw.Extra, next mw.Next) {
	v := r.URL.Query().Get("value")
	if v == "" {
		msg.Code = df.HTTP_STATUS_PARAM_ERROR
		msg.Msg = "参数错误"
		return
	}
	msg.Code = 200
	msg.Msg = tl.ToSha512String(v)
}

func init() {
	rt.RegisterRouter("/api/sha512", rt.NewFactory(func() rt.Router {
		rt := rt.NewRouter()
		rt.UseMiddleware(&sha512Router{})
		return rt
	}))
}
