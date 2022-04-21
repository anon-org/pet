package impl

import (
	"github.com/anon-org/arithmetic/api"
	"net/http"
)

type Handler struct {
	Svc api.Service
}

func (h Handler) Foo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Svc.Foo()
	}
}
