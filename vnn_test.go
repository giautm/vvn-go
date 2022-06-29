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

	res, err := c.OCRecognition(context.Background(), api.OCRInputForm{
		RequestID: uuid.NewString(),
		Image: http.MultipartFile{
			Name: "random_name_abc.jpeg",
			File: f,
		},
		IDFullThr: api.NewOptFloat32(0.8),
	})
	r.NoError(err)

	switch t := res.(type) {
	case *api.OCRResult:
		r.NotEmpty(t.ID.Value)
		r.Equal(t.IDCheck.Value, api.IDCheckEnumREAL)
	case *api.GatewayError:
		r.Fail(t.Message)
	}
}

func TestFaceIDVerification_integration(t *testing.T) {
	apiKey := os.Getenv("VVN_API_KEY")
	if apiKey == "" {
		t.Skip("VVN_API_KEY is not set")
	}

	r := require.New(t)
	c, err := api.NewClient(vvn.ServerProduction, vvn.StaticKey(apiKey))
	r.NoError(err)

	res, err := c.FaceVerification(context.Background(), &api.FaceIDVerificationInput{
		RequestID: uuid.NewString(),
		ImageCard: "https://cdnimg.vietnamplus.vn/t660/uploaded/hotnnz/2022_03_14/obama.jpg",
		ImageLive: "http://c.files.bbci.co.uk/9E1D/production/_111777404_obamaendorsebiden.jpg",
	})
	r.NoError(err)

	switch t := res.(type) {
	case *api.FaceIDVerificationResult:
		r.Equal(t.Message.ErrorCode, "000")
		r.Equal(t.Message.ErrorMessage, "OK")

		r.Equal(t.VerifyResultText.Value, "Same person")
		r.Equal(t.WearingMask.Value, api.MaskResultEnumNO)
	case *api.GatewayError:
		r.Fail(t.Message)
	}
}

func TestFaceRecognition_integration(t *testing.T) {
	apiKey := os.Getenv("VVN_API_KEY")
	if apiKey == "" {
		t.Skip("VVN_API_KEY is not set")
	}

	r := require.New(t)
	c, err := api.NewClient(vvn.ServerProduction, vvn.StaticKey(apiKey))
	r.NoError(err)

	res, err := c.FaceRecognition(context.Background(), &api.FaceIDRecognitionInput{
		RequestID: uuid.NewString(),
		Image:     "https://cdnimg.vietnamplus.vn/t660/uploaded/hotnnz/2022_03_14/obama.jpg",
	})
	r.NoError(err)

	switch t := res.(type) {
	case *api.FaceIDRecognitionResult:
		r.Equal(t.Message.ErrorCode, "000")
		r.Equal(t.Message.ErrorMessage, "OK")

		r.Len(t.RecognitionResult, 1)
		r.Len(t.RecognitionResult[0].TopK, 5)
		r.Equal(t.RecognitionResult[0].TopK[0].UniqueName, "FUNDIIN-BLACKLIST-3043544A51EC68E1202B840FE2C30E98")
	case *api.GatewayError:
		r.Fail(t.Message)
	}
}
