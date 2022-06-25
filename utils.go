package vvn

import (
	"context"

	"giautm.dev/vvn/api"
)

// StaticKey is a static API key.
type StaticKey string

var (
	_ api.SecuritySource = (*StaticKey)(nil)
)

func (s StaticKey) APIKey(context.Context, string) (api.APIKey, error) {
	return api.APIKey{APIKey: string(s)}, nil
}
