// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	cfg       config
	requests  syncint64.Counter
	errors    syncint64.Counter
	duration  syncint64.Histogram
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...Option) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := &Client{
		cfg:       newConfig(opts...),
		sec:       sec,
		serverURL: u,
	}
	if c.requests, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientRequestCount); err != nil {
		return nil, err
	}
	if c.errors, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientErrorsCount); err != nil {
		return nil, err
	}
	if c.duration, err = c.cfg.Meter.SyncInt64().Histogram(otelogen.ClientDuration); err != nil {
		return nil, err
	}
	return c, nil
}

// NewFaceIDVerification invokes newFaceIDVerification operation.
//
// POST /faceid/verification
func (c *Client) NewFaceIDVerification(ctx context.Context, request VerificationInput) (res NewFaceIDVerificationRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("newFaceIDVerification"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "NewFaceIDVerification",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	var (
		contentType string
		reqBody     func() (io.ReadCloser, error)
	)
	contentType = "application/json"
	fn, err := encodeNewFaceIDVerificationRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	reqBody = fn

	u := uri.Clone(c.serverURL)
	u.Path += "/faceid/verification"

	body, err := reqBody()
	if err != nil {
		return res, errors.Wrap(err, "request body")
	}
	defer body.Close()

	r := ht.NewRequest(ctx, "POST", u, body)
	r.GetBody = reqBody

	r.Header.Set("Content-Type", contentType)

	if err := c.securityAPIKey(ctx, "NewFaceIDVerification", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeNewFaceIDVerificationResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// NewOCRRecognition invokes newOCRRecognition operation.
//
// Cung cấp phương thức để trích xuất thông tin trên các văn bản tài liệu như:
// Giấy phép lái xe (GPLX), Passport, CMND, Căn cước công dân (CCCD) ...
//
// POST /ocr/recognition
func (c *Client) NewOCRRecognition(ctx context.Context, request OCRInputForm) (res NewOCRRecognitionRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("newOCRRecognition"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "NewOCRRecognition",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	var (
		contentType string
		reqBody     func() (io.ReadCloser, error)
	)
	contentType = "multipart/form-data"
	fn, ct, err := encodeNewOCRRecognitionRequest(request, span)
	if err != nil {
		return res, err
	}
	reqBody = fn
	contentType = ct

	u := uri.Clone(c.serverURL)
	u.Path += "/ocr/recognition"

	body, err := reqBody()
	if err != nil {
		return res, errors.Wrap(err, "request body")
	}
	defer body.Close()

	r := ht.NewRequest(ctx, "POST", u, body)
	r.GetBody = reqBody

	r.Header.Set("Content-Type", contentType)

	if err := c.securityAPIKey(ctx, "NewOCRRecognition", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeNewOCRRecognitionResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
