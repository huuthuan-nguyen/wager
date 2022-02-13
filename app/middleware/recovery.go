package middleware

import (
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/wager/app/utils"
	"net/http"
)

// RecoveringMiddleware /**
var RecoveringMiddleware = mux.MiddlewareFunc(
	func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// recover when panic happening
			defer func() {
				if err := recover(); err != nil {
					// Handling the error
					if e, ok := err.(utils.Error); ok {
						statusCode := e.StatusCode
						w.WriteHeader(statusCode)
						w.Write([]byte(e.Error()))
					} else {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("Something went wrong."))
					}
				}
			}()

			next.ServeHTTP(w, r)
		})
	})
