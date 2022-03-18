package auth

import (
	"main/common/df"
	"main/common/mw"
	"main/common/tl"
	"net/http"
	"net/url"
)

const tokenHeaderKey string = "--open-the-door"

var authMiddleware = mw.NewMiddleware(func(r *http.Request, msg *mw.Message, extra mw.Extra, next mw.Next) {
	token := r.Header.Get(tokenHeaderKey)
	if token == "" {
		msg.Code = df.HTTP_STATUS_AUTH_ERROR
		msg.Msg = "请先登录"
		return
	}
	tn, _ := url.QueryUnescape(token)
	acc := tl.Decrypt(tl.Decrypt(tn))
	u, e := users[acc]
	if !e {
		msg.Code = df.HTTP_STATUS_AUTH_ERROR
		msg.Msg = "请先登录"
		return
	}
	if u.Token != token {
		msg.Code = df.HTTP_STATUS_AUTH_ERROR
		msg.Msg = "登录已失效"
		return
	}

	extra["user"] = u

	next()
})
