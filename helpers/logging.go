package helpers

import (
	"net/http"
	_ "time"

	"github.com/sirupsen/logrus"
)

func WithLogging(h http.Handler) http.Handler {
	logFn := func(rw http.ResponseWriter, r *http.Request) {
		// start := time.Now()

		uri := r.RequestURI
		// method := r.Method
		h.ServeHTTP(rw, r) // serve the original request

		// duration := time.Since(start)

		// log request details
		logrus.WithField("Test", uri)
		// log.WithFields(log.Fields{
		// 		"uri":      uri,
		// 		"method":   method,
		// 		"duration": duration,
		// })
	}
	return http.HandlerFunc(logFn)
}
