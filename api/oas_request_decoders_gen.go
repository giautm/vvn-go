// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeNewFaceIDVerificationRequest(r *http.Request, span trace.Span) (
	req NewFaceIDVerificationReq,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request VerificationInput
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return &request, close, nil
	case "multipart/form-data":
		var request VerificationInputForm

		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		form := url.Values(r.MultipartForm.Value)

		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "check_3_random_pose",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCheck3RandomPoseVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						requestDotCheck3RandomPoseVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Check3RandomPose.SetTo(requestDotCheck3RandomPoseVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"check_3_random_pose\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "check_3_straight_pose",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotCheck3StraightPoseVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						requestDotCheck3StraightPoseVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Check3StraightPose.SetTo(requestDotCheck3StraightPoseVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"check_3_straight_pose\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "fake_threshold",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotFakeThresholdVal float64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToFloat64(val)
						if err != nil {
							return err
						}

						requestDotFakeThresholdVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.FakeThreshold.SetTo(requestDotFakeThresholdVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"fake_threshold\"")
				}
				if err := func() error {
					if request.FakeThreshold.Set {
						if err := func() error {
							if err := (validate.Float{}).Validate(float64(request.FakeThreshold.Value)); err != nil {
								return errors.Wrap(err, "float")
							}
							return nil
						}(); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return req, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "mask_threshold",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotMaskThresholdVal float64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToFloat64(val)
						if err != nil {
							return err
						}

						requestDotMaskThresholdVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.MaskThreshold.SetTo(requestDotMaskThresholdVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"mask_threshold\"")
				}
				if err := func() error {
					if request.MaskThreshold.Set {
						if err := func() error {
							if err := (validate.Float{}).Validate(float64(request.MaskThreshold.Value)); err != nil {
								return errors.Wrap(err, "float")
							}
							return nil
						}(); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return req, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "request_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					request.RequestID = c
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"request_id\"")
				}
			} else {
				return req, close, errors.Wrap(err, "query")
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "return_feature",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotReturnFeatureVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						requestDotReturnFeatureVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ReturnFeature.SetTo(requestDotReturnFeatureVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"return_feature\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "sim_threshold_level1",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSimThresholdLevel1Val float64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToFloat64(val)
						if err != nil {
							return err
						}

						requestDotSimThresholdLevel1Val = c
						return nil
					}(); err != nil {
						return err
					}
					request.SimThresholdLevel1.SetTo(requestDotSimThresholdLevel1Val)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"sim_threshold_level1\"")
				}
				if err := func() error {
					if request.SimThresholdLevel1.Set {
						if err := func() error {
							if err := (validate.Float{}).Validate(float64(request.SimThresholdLevel1.Value)); err != nil {
								return errors.Wrap(err, "float")
							}
							return nil
						}(); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return req, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "sim_threshold_level2",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotSimThresholdLevel2Val float64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToFloat64(val)
						if err != nil {
							return err
						}

						requestDotSimThresholdLevel2Val = c
						return nil
					}(); err != nil {
						return err
					}
					request.SimThresholdLevel2.SetTo(requestDotSimThresholdLevel2Val)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"sim_threshold_level2\"")
				}
				if err := func() error {
					if request.SimThresholdLevel2.Set {
						if err := func() error {
							if err := (validate.Float{}).Validate(float64(request.SimThresholdLevel2.Value)); err != nil {
								return errors.Wrap(err, "float")
							}
							return nil
						}(); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return req, close, errors.Wrap(err, "validate")
				}
			}
		}
		if err := func() error {
			files, ok := r.MultipartForm.File["image_card"]
			if !ok || len(files) < 1 {
				return validate.ErrFieldRequired
			}
			fh := files[0]

			f, err := fh.Open()
			if err != nil {
				return errors.Wrap(err, "open")
			}
			closers = append(closers, f)
			request.ImageCard = ht.MultipartFile{
				Name:   fh.Filename,
				File:   f,
				Header: fh.Header,
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"image_card\"")
		}
		if err := func() error {
			files, ok := r.MultipartForm.File["image_live"]
			if !ok || len(files) < 1 {
				return validate.ErrFieldRequired
			}
			fh := files[0]

			f, err := fh.Open()
			if err != nil {
				return errors.Wrap(err, "open")
			}
			closers = append(closers, f)
			request.ImageLive = ht.MultipartFile{
				Name:   fh.Filename,
				File:   f,
				Header: fh.Header,
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"image_live\"")
		}
		if err := func() error {
			files, ok := r.MultipartForm.File["image_live1"]
			if !ok || len(files) < 1 {
				return nil
			}
			fh := files[0]

			f, err := fh.Open()
			if err != nil {
				return errors.Wrap(err, "open")
			}
			closers = append(closers, f)
			request.ImageLive1.SetTo(ht.MultipartFile{
				Name:   fh.Filename,
				File:   f,
				Header: fh.Header,
			})
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"image_live1\"")
		}
		if err := func() error {
			files, ok := r.MultipartForm.File["image_live2"]
			if !ok || len(files) < 1 {
				return nil
			}
			fh := files[0]

			f, err := fh.Open()
			if err != nil {
				return errors.Wrap(err, "open")
			}
			closers = append(closers, f)
			request.ImageLive2.SetTo(ht.MultipartFile{
				Name:   fh.Filename,
				File:   f,
				Header: fh.Header,
			})
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"image_live2\"")
		}
		if err := func() error {
			files, ok := r.MultipartForm.File["image_live3"]
			if !ok || len(files) < 1 {
				return nil
			}
			fh := files[0]

			f, err := fh.Open()
			if err != nil {
				return errors.Wrap(err, "open")
			}
			closers = append(closers, f)
			request.ImageLive3.SetTo(ht.MultipartFile{
				Name:   fh.Filename,
				File:   f,
				Header: fh.Header,
			})
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"image_live3\"")
		}

		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeNewOCRRecognitionRequest(r *http.Request, span trace.Span) (
	req OCRInputForm,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "multipart/form-data":
		var request OCRInputForm

		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		form := url.Values(r.MultipartForm.Value)

		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "id_full_thr",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIDFullThrVal float32
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToFloat32(val)
						if err != nil {
							return err
						}

						requestDotIDFullThrVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.IDFullThr.SetTo(requestDotIDFullThrVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"id_full_thr\"")
				}
				if err := func() error {
					if request.IDFullThr.Set {
						if err := func() error {
							if err := (validate.Float{}).Validate(float64(request.IDFullThr.Value)); err != nil {
								return errors.Wrap(err, "float")
							}
							return nil
						}(); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return req, close, errors.Wrap(err, "validate")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "request_id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					request.RequestID = c
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"request_id\"")
				}
			} else {
				return req, close, errors.Wrap(err, "query")
			}
		}
		if err := func() error {
			files, ok := r.MultipartForm.File["image"]
			if !ok || len(files) < 1 {
				return validate.ErrFieldRequired
			}
			fh := files[0]

			f, err := fh.Open()
			if err != nil {
				return errors.Wrap(err, "open")
			}
			closers = append(closers, f)
			request.Image = ht.MultipartFile{
				Name:   fh.Filename,
				File:   f,
				Header: fh.Header,
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"image\"")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}
