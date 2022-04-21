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
	pet := rpc.Wire("http://localhost:8000", 60*time.Second)

	store(pet)
	fetch(pet)
}

func store(pet api.Service) {
	_, err := pet.Store(context.Background(), Name())
	if err != nil {
		panic(err)
	}
}

func fetch(pet api.Service) {
	result, err := pet.Fetch(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
