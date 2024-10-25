package ihandlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	// DefaultCORSMaxAge is the recommended TTL for a CORS cache.
	DefaultCORSMaxAge = time.Hour * 24

	// DefaultCORSMethods are the minimal methods CORS needs to support.
	DefaultCORSMethods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodDelete,
		http.MethodPut,
		http.MethodPatch,
		http.MethodOptions,
	}

	// DefaultCORSHeaders allows all headers for CORS middleware.
	DefaultCORSHeaders = []string{"*"}
	DefaultCORSOrigins = []string{"*"}
)

const (
	defaultReferrerPolicy = "no-referrer-when-downgrade"
)

func CORSMiddleware(maxAge time.Duration, headers, methods, origins []string) func(http.Handler) http.Handler {
	const sep = ","
	allowedHeaders := strings.Join(headers, sep)
	allowedMethods := strings.Join(methods, sep)
	allowedOrigins := strings.Join(origins, sep)
	maxAgeInSeconds := strconv.FormatInt(int64(maxAge/time.Second), 10)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
			w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
			w.Header().Set("Access-Control-Max-Age", maxAgeInSeconds)
			w.Header().Set("Referrer-Policy", defaultReferrerPolicy)
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func CORSMiddlewareWithDefaults() func(http.Handler) http.Handler {
	return CORSMiddleware(DefaultCORSMaxAge, DefaultCORSHeaders, DefaultCORSMethods, DefaultCORSOrigins)
}
