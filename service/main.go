package main

import (
	"github.com/anon-org/arithmetic/impl"
	"net/http"
)

func main() {

	http.ListenAndServe(":8000", impl.Wire().Foo())
}
