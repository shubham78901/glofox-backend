package api

import (
	_ "glofox-backend/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupSwagger(router *mux.Router) {

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
