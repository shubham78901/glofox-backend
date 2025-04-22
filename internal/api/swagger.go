// File: internal/api/swagger.go

package api

import (
	_ "glofox-backend/docs" // This is required for swagger to find the docs

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupSwagger adds swagger documentation routes to the router
func SetupSwagger(router *mux.Router) {
	// Serve Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
