package main

import (
	"context"
	"fmt"
	"github.com/anon-org/pet/api"
	"github.com/anon-org/pet/rpc"
	. "github.com/dustinkirkland/golang-petname"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wire := rpc.Wire("http://localhost:8000", 60*time.Second)

	store(wire)
	fetch(wire)
}

func store(rpc api.Service) {
	_, err := rpc.Store(context.Background(), Name())
	if err != nil {
		panic(err)
	}
}

func fetch(rpc api.Service) {
	result, err := rpc.Fetch(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
