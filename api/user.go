package api

import (
	"bytes"
	"encoding/json"
)

type Users []*User

func (u Users) String() string {
	var b bytes.Buffer
	e := json.NewEncoder(&b)
	e.SetIndent("", "  ")
	if err := e.Encode(u); err != nil {
		return "[]"
	}

	return b.String()
}

type User struct {
	ID   string
	Name string
}
