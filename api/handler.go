package api

import "net/http"

type Handler interface {
	Foo() http.HandlerFunc
}
