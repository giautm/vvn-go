// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeNewFaceIDVerificationRequestJSON(
	req VerificationInput,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
func encodeNewFaceIDVerificationRequest(
	req VerificationInputForm,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	contentType string,
	rerr error,
) {
	q := uri.NewQueryEncoder()
	{
		// Encode "check_3_random_pose" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "check_3_random_pose",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.Check3RandomPose.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "check_3_straight_pose" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "check_3_straight_pose",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.Check3StraightPose.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "fake_threshold" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "fake_threshold",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.FakeThreshold.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "mask_threshold" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "mask_threshold",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.MaskThreshold.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "request_id" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "request_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(req.RequestID))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "return_feature" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "return_feature",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.ReturnFeature.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "sim_threshold_level1" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "sim_threshold_level1",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.SimThresholdLevel1.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "sim_threshold_level2" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "sim_threshold_level2",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.SimThresholdLevel2.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	getBody, contentType := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := req.ImageCard.WriteMultipart("image_card", w); err != nil {
			return errors.Wrap(err, "write \"image_card\"")
		}
		if err := req.ImageLive.WriteMultipart("image_live", w); err != nil {
			return errors.Wrap(err, "write \"image_live\"")
		}
		if val, ok := req.ImageLive1.Get(); ok {
			if err := val.WriteMultipart("image_live1", w); err != nil {
				return errors.Wrap(err, "write \"image_live1\"")
			}
		}
		if val, ok := req.ImageLive2.Get(); ok {
			if err := val.WriteMultipart("image_live2", w); err != nil {
				return errors.Wrap(err, "write \"image_live2\"")
			}
		}
		if val, ok := req.ImageLive3.Get(); ok {
			if err := val.WriteMultipart("image_live3", w); err != nil {
				return errors.Wrap(err, "write \"image_live3\"")
			}
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	return getBody, contentType, nil
}
func encodeNewOCRRecognitionRequest(
	req OCRInputForm,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	contentType string,
	rerr error,
) {
	q := uri.NewQueryEncoder()
	{
		// Encode "id_full_thr" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id_full_thr",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := req.IDFullThr.Get(); ok {
				return e.EncodeValue(conv.Float32ToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "request_id" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "request_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(req.RequestID))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	getBody, contentType := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := req.Image.WriteMultipart("image", w); err != nil {
			return errors.Wrap(err, "write \"image\"")
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	return getBody, contentType, nil
}
