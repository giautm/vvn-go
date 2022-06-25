// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// HandleNewFaceIDVerificationRequest handles newFaceIDVerification operation.
//
// POST /faceid/verification
func (s *Server) handleNewFaceIDVerificationRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("newFaceIDVerification"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "NewFaceIDVerification",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	ctx, err = s.securityAPIKey(ctx, "NewFaceIDVerification", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			Operation: "NewFaceIDVerification",
			Security:  "APIKey",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, close, err := s.decodeNewFaceIDVerificationRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			Operation: "NewFaceIDVerification",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.NewFaceIDVerification(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeNewFaceIDVerificationResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleNewOCRRecognitionRequest handles newOCRRecognition operation.
//
// POST /ocr/recognition
func (s *Server) handleNewOCRRecognitionRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("newOCRRecognition"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "NewOCRRecognition",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	ctx, err = s.securityAPIKey(ctx, "NewOCRRecognition", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			Operation: "NewOCRRecognition",
			Security:  "APIKey",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, close, err := s.decodeNewOCRRecognitionRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			Operation: "NewOCRRecognition",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.NewOCRRecognition(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeNewOCRRecognitionResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

func (s *Server) badRequest(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
	span trace.Span,
	otelAttrs []attribute.KeyValue,
	err error,
) {
	span.RecordError(err)
	span.SetStatus(codes.Error, "BadRequest")
	s.errors.Add(ctx, 1, otelAttrs...)
	s.cfg.ErrorHandler(ctx, w, r, err)
}