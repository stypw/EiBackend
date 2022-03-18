package rt

import (
	"net/http"

	"main/common/mw"
)

type router struct {
	middlewares []mw.Middleware
	middleIndex int
}

func (rter *router) callMiddleware(r *http.Request, msg *mw.Message, extra mw.Extra) {
	l := len(rter.middlewares)
	if rter.middleIndex >= l {
		return
	}
	m := rter.middlewares[rter.middleIndex]
	rter.middleIndex++
	m.OnRequest(r, msg, extra, func() {
		rter.callMiddleware(r, msg, extra)
	})
}
func (rter *router) OnRequest(r *http.Request, msg *mw.Message, extra mw.Extra) {
	rter.middleIndex = 0
	rter.callMiddleware(r, msg, extra)
}

func (rter *router) UseMiddleware(middleware mw.Middleware) {
	rter.middlewares = append(rter.middlewares, middleware)
}
