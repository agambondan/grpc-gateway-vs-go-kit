package transport

import (
	"git.bluebird.id/firman.agam/go-kit/internal/endpoint"
	calculatorTransport "git.bluebird.id/firman.agam/go-kit/internal/transport/calculator"
	"git.bluebird.id/firman.agam/go-kit/internal/utils"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(r *mux.Router, endpoint endpoint.Endpoints) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(utils.EncodeLegacyError),
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	}

	calculatorSubrouter := r.PathPrefix("/calculator").Subrouter()

	newCalculatorHTTPHandler := calculatorTransport.NewCalculatorHTTPHandler(endpoint, opts...)
	newCalculatorHTTPHandler.SetupRoutes(calculatorSubrouter)

	//// add more if you have more service transport
	// healthSubrouter := r.PathPrefix("/health").Subrouter()
}
