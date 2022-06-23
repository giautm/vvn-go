package vvn

import (
	"context"
	"net/http"
)

// WithAPIKey sets the key header value to the given value.
func WithAPIKey(apiKey string) ClientOption {
	editor := func(ctx context.Context, req *http.Request) error {
		req.Header.Add("key", apiKey)
		return nil
	}
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, editor)
		return nil
	}
}
