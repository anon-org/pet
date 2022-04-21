package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/anon-org/arithmetic/api"
	"net/http"
	"time"
)

type rpc struct {
	BaseURL string
	Timeout time.Duration
}

func (r rpc) Fetch(ctx context.Context) (api.Users, error) {
	url := fmt.Sprintf("%s/", r.BaseURL)

	ctx, cancel := context.WithTimeout(ctx, r.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result api.Users
	return result, json.NewDecoder(resp.Body).Decode(&result)
}

func (r rpc) Store(ctx context.Context, name string) (*api.User, error) {
	url := fmt.Sprintf("%s/", r.BaseURL)

	ctx, cancel := context.WithTimeout(ctx, r.Timeout)
	defer cancel()

	b, err := json.Marshal(api.Request{Name: name})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result *api.User
	return result, json.NewDecoder(resp.Body).Decode(&result)
}
