# vvn-go

> Unofficial client for VVN AI EKYC

This package uses [ogen](https://github.com/ogen-go/ogen) to generate the client from the VVN AI APIs specification.

Install:

```
go get -u giautm.dev/vvn@main
```

See the example from test:
```go
package vvn_test

import (
	"context"
	"os"
	"testing"

	"giautm.dev/vvn"
	"giautm.dev/vvn/api"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/http"
	"github.com/stretchr/testify/require"
)

func TestNewOCRRecognition_integration(t *testing.T) {
	apiKey := os.Getenv("VVN_API_KEY")
	if apiKey == "" {
		t.Skip("VVN_API_KEY is not set")
	}
	documentFile := os.Getenv("VVN_DOCUMENT_FILE")
	if documentFile == "" {
		t.Skip("VVN_DOCUMENT_FILE is not set")
	}

	r := require.New(t)
	c, err := api.NewClient(vvn.ServerProduction, vvn.StaticKey(apiKey))
	r.NoError(err)

	f, err := os.OpenFile(documentFile, os.O_RDONLY, 0666)
	r.NoError(err)

	res, err := c.NewOCRRecognition(context.Background(), api.OCRInputForm{
		RequestID: uuid.NewString(),
		Image: http.MultipartFile{
			Name: "random_name_abc.jpeg",
			File: f,
		},
		IDFullThr: api.NewOptFloat32(0.8),
	})
	r.NoError(err)

	result, ok := res.(*api.OCRResult)
	r.Equal(ok, true)
	r.NotEmpty(result.ID.Value)
	r.Equal(result.IDCheck.Value, api.OCRResultIDCheckREAL)
}
```

