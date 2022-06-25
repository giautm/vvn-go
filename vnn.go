package vvn

import (
	"context"

	"giautm.dev/vvn/api"
)

const (
	// ServerProduction is the base URL for the production server.
	ServerProduction = "https://api.cloudekyc.com/v3.2"
)

// StaticKey is a static API key.
type StaticKey string

var (
	_ api.SecuritySource = (*StaticKey)(nil)
)

// APIKey implement api.SecuritySource.
func (s StaticKey) APIKey(context.Context, string) (api.APIKey, error) {
	return api.APIKey{APIKey: string(s)}, nil
}
