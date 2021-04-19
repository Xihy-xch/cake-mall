package client

import (
	"net/http"

	"github.com/dghubble/sling"
)

type HClient struct {
	Sling *sling.Sling
}

func NewClient(host string) *HClient {

	doer := LogDoer{doer: http.DefaultClient}
	return &HClient{
		Sling: sling.New().Base(host).Doer(doer),
	}
}
