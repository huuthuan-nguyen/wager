package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
)

// CORSMiddleware /**
var CORSMiddleware = mux.MiddlewareFunc(
	func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, HEAD")
			if r.Method == http.MethodOptions {
				return
			}

			next.ServeHTTP(w, r)
		})
	})
