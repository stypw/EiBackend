package mw

import (
	"net/http"
)

type Message struct {
	Code int
	Msg  string
}
type Extra map[string]interface{}

type Next func()

type Middleware interface {
	OnRequest(r *http.Request, msg *Message, extra Extra, next Next)
}
