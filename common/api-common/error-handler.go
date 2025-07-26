package apicommon

import (
	"errors"
	"net/http"
	"vitals-guard/common/errs"
)

type errorHandler func(http.ResponseWriter, *http.Request) error

func ErrorHandler(f errorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {

			status := http.StatusInternalServerError
			var sc interface{ HTTPStatus() int }
			if errors.As(err, &sc) {
				status = sc.HTTPStatus()
			}

			var resp *errs.AppError
			if appErr, ok := err.(*errs.AppError); ok {
				resp = appErr
			} else {
				resp = errs.Wrap(err, errs.ErrInternalServer.Code, errs.ErrInternalServer.Status)
			}
			WriteError(w, resp, status)
		}
	}
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
