package errs

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-errors/errors"
	"github.com/withbioma/interview-backend/lib/errs/clienterr"
	"github.com/withbioma/interview-backend/lib/errs/systemerr"
)

var IdempotencyError = systemerr.New("idempotency error")

func handleError(ctx context.Context, err error, code int) {
	if code >= 500 || code == 400 {
		SendToSentry(ctx, err, code)
	}
}

// SendToSentryWithoutCtx sends error to sentry whether or not the error comes from http endpoints.
func SendToSentryWithoutCtx(err error, code int) {
	SendToSentry(context.Background(), err, code)
}

// SendToSentry sends error to sentry with http request context.
func SendToSentry(ctx context.Context, err error, code int) {
	switch e := err.(type) {
	case clienterr.ClientError:
		// if there is no original error then there is nothing to report
		if e.OriginalError() != nil {
			log.Println(e.OriginalError())
		} else {
			log.Println(e)
		}
	default:
		log.Println(e)
	}
}

// APIErrorResponse canonical struct for API error responses.
type APIErrorResponse struct {
	Errors []clienterr.APIError `json:"errors"`
}

// APIError writes default error message and header for a given http status code.
func APIError(ctx context.Context, w http.ResponseWriter, code int, err error) {
	if err != nil {
		handleError(ctx, err, code)

		switch e := err.(type) {
		case clienterr.ClientError:
			w.WriteHeader(code)
			if encodeErr := json.NewEncoder(w).Encode(
				APIErrorResponse{Errors: []clienterr.APIError{e.Resolve()}},
			); encodeErr != nil {
				SendToSentry(ctx, encodeErr, -1)
			}
		case systemerr.SystemError:
			log.Println(errors.Wrap(err, 1))

			w.WriteHeader(code)
			http.Error(w, http.StatusText(code), code)
		default:
			log.Println(errors.Wrap(err, 1))

			w.WriteHeader(code)
			http.Error(w, http.StatusText(code), code)
		}
	} else {
		w.WriteHeader(code)
		http.Error(w, http.StatusText(code), code)
	}
}

func IsIdempotencyError(err error) bool {
	return err == IdempotencyError
}
