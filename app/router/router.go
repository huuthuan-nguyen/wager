package router

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/wager/app/handler"
	"github.com/huuthuan-nguyen/wager/app/middleware"
	"github.com/huuthuan-nguyen/wager/config"
	"net/http"
)

// NewRouter /**
func NewRouter(config *config.Config, handler *handler.Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	recoveryMiddleware := middleware.RecoveringMiddleware       // recovering from panic
	loggingMiddleware := middleware.LoggingMiddleware           // logging
	compressMiddleware := handlers.CompressHandler              // compressing body for saving bandwidth
	rateLimitingMiddleware := middleware.RateLimitingMiddleware // rate limiting
	corsMiddleware := middleware.CORSMiddleware                 // allow cors for browser request
	router.Use(
		compressMiddleware,
		loggingMiddleware,
		recoveryMiddleware,
		corsMiddleware,
		rateLimitingMiddleware,
	)

	router.HandleFunc("/wagers", handler.WagerPlace).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/buy/{id:[0-9]+}", handler.WagerBuy).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/wagers", handler.WagerIndex).Methods(http.MethodGet, http.MethodOptions)

	return router
}
