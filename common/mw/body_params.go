package mw

import (
	"net/http"

	"main/common/df"
	"main/common/kv"
)

var JsonBodyMiddleware = NewMiddleware(func(r *http.Request, msg *Message, extra Extra, next Next) {
	body := kv.FromStream(r.Body)
	if body == nil {
		msg.Code = df.HTTP_STATUS_PARAM_ERROR
		msg.Msg = "body为空"
		return
	}
	extra["body"] = body
	next()
})
