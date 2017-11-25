// +build go1.8

package middleware

import (
	"errors"
	"net/http"
)

func (w *maybeCompressResponseWriter) Push(target string, opts *http.PushOptions) error {
	if ps, ok := w.w.(http.Pusher); ok {
		return ps.Push(target, opts)
	}
	return errors.New("tea/middleware: http.Pusher is unavailable on the writer")
}
