package rt

import "main/common/mw"

func NewRouter() Router {
	return &router{middlewares: make([]mw.Middleware, 0), middleIndex: 0}
}
