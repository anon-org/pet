package api

import "encoding/json"

type Users []*User

func (u Users) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		return "[]"
	}

	return string(b)
}

type User struct {
	ID   string
	Name string
}

func (u User) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		return "{}"
	}

	return string(b)
}
