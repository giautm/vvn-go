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

func encodeFaceRecognitionRequestJSON(
	req FaceIDRecognitionInput,
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
func encodeFaceRecognitionRequest(
	req FaceIDRecognitionInputForm,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	contentType string,
	rerr error,
) {
	request := req

	q := uri.NewQueryEncoder()
	{
		// Encode "request_id" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "request_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.RequestID))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	getBody, contentType := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := request.Image.WriteMultipart("image", w); err != nil {
			return errors.Wrap(err, "write \"image\"")
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	return getBody, contentType, nil
}
func encodeFaceRegisterRequestJSON(
	req FaceIDRegisterInput,
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
func encodeFaceRegisterRequest(
	req FaceIDRegisterInputForm,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	contentType string,
	rerr error,
) {
	request := req

	q := uri.NewQueryEncoder()
	{
		// Encode "unique_name" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "unique_name",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.UniqueName))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "force" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "force",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Force.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "person_name" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "person_name",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.PersonName.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	getBody, contentType := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := request.Image.WriteMultipart("image", w); err != nil {
			return errors.Wrap(err, "write \"image\"")
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	return getBody, contentType, nil
}
func encodeFaceUnregisterRequestJSON(
	req FaceUnregisterReq,
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
func encodeFaceVerificationRequestJSON(
	req FaceIDVerificationInput,
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
func encodeFaceVerificationRequest(
	req FaceIDVerificationInputForm,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	contentType string,
	rerr error,
) {
	request := req

	q := uri.NewQueryEncoder()
	{
		// Encode "check_3_random_pose" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "check_3_random_pose",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Check3RandomPose.Get(); ok {
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
			if val, ok := request.Check3StraightPose.Get(); ok {
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
			if val, ok := request.FakeThreshold.Get(); ok {
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
			if val, ok := request.MaskThreshold.Get(); ok {
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
			return e.EncodeValue(conv.StringToString(request.RequestID))
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
			if val, ok := request.ReturnFeature.Get(); ok {
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
			if val, ok := request.SimThresholdLevel1.Get(); ok {
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
			if val, ok := request.SimThresholdLevel2.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	getBody, contentType := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := request.ImageCard.WriteMultipart("image_card", w); err != nil {
			return errors.Wrap(err, "write \"image_card\"")
		}
		if err := request.ImageLive.WriteMultipart("image_live", w); err != nil {
			return errors.Wrap(err, "write \"image_live\"")
		}
		if val, ok := request.ImageLive1.Get(); ok {
			if err := val.WriteMultipart("image_live1", w); err != nil {
				return errors.Wrap(err, "write \"image_live1\"")
			}
		}
		if val, ok := request.ImageLive2.Get(); ok {
			if err := val.WriteMultipart("image_live2", w); err != nil {
				return errors.Wrap(err, "write \"image_live2\"")
			}
		}
		if val, ok := request.ImageLive3.Get(); ok {
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
func encodeOCRecognitionRequest(
	req OCRInputForm,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	contentType string,
	rerr error,
) {
	request := req

	q := uri.NewQueryEncoder()
	{
		// Encode "id_full_thr" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id_full_thr",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.IDFullThr.Get(); ok {
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
			return e.EncodeValue(conv.StringToString(request.RequestID))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	getBody, contentType := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := request.Image.WriteMultipart("image", w); err != nil {
			return errors.Wrap(err, "write \"image\"")
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	return getBody, contentType, nil
}
