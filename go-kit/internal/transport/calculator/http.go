package calculator

import (
	"net/http"

	"git.bluebird.id/firman.agam/go-kit/internal/endpoint"
	"git.bluebird.id/firman.agam/go-kit/internal/transport/calculator/decoder"
	"git.bluebird.id/firman.agam/go-kit/internal/transport/middleware"
	"git.bluebird.id/firman.agam/go-kit/internal/utils"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// CalculatorHTTPHandler is the concrete implementation of HTTPHandler for the calculator service.
type CalculatorHTTPHandler struct {
	endpoints endpoint.Endpoints
	options   []kithttp.ServerOption
}

// NewCalculatorHTTPHandler initializes a new CalculatorHTTPHandler with injected dependencies.
func NewCalculatorHTTPHandler(endpoints endpoint.Endpoints, opts ...kithttp.ServerOption) *CalculatorHTTPHandler {
	return &CalculatorHTTPHandler{
		endpoints: endpoints,
		options:   opts,
	}
}

// SetupRoutes sets up the HTTP routes for the calculator service.
func (h *CalculatorHTTPHandler) SetupRoutes(r *mux.Router) {
	addEndpoint := h.endpoints.AddEndpoint()
	addEndpoint = middleware.TransportLogging("add_endpoint")(addEndpoint)
	addEndpoint = middleware.BasicAuth()(addEndpoint)
	addHandler := kithttp.NewServer(
		addEndpoint,
		decoder.HTTPAdd,
		utils.EncodeHTTPResponseWithData,
		h.options...,
	)
	subtractEndpoint := h.endpoints.SubtractEndpoint()
	subtractEndpoint = middleware.TransportLogging("subtract_endpoint")(subtractEndpoint)
	subtractEndpoint = middleware.BasicAuth()(subtractEndpoint)
	subtractHandler := kithttp.NewServer(
		subtractEndpoint,
		decoder.HTTPSubtract,
		utils.EncodeHTTPResponseWithData,
		h.options...,
	)

	multiplyEndpoint := h.endpoints.MultiplyEndpoint()
	multiplyEndpoint = middleware.TransportLogging("multiply_endpoint")(multiplyEndpoint)
	multiplyEndpoint = middleware.BasicAuth()(multiplyEndpoint)
	multiplyHandler := kithttp.NewServer(
		multiplyEndpoint,
		decoder.HTTPMultiply,
		utils.EncodeHTTPResponseWithData,
		h.options...,
	)

	divideEndpoint := h.endpoints.DivideEndpoint()
	divideEndpoint = middleware.TransportLogging("divide_endpoint")(divideEndpoint)
	divideEndpoint = middleware.BasicAuth()(divideEndpoint)
	divideHandler := kithttp.NewServer(
		divideEndpoint,
		decoder.HTTPDivide,
		utils.EncodeHTTPResponseWithData,
		h.options...,
	)

	fibonacciEndpoint := h.endpoints.FibonacciEndpoint()
	fibonacciEndpoint = middleware.TransportLogging("fibonacci_endpoint")(fibonacciEndpoint)
	fibonacciEndpoint = middleware.BasicAuth()(fibonacciEndpoint)
	fibonacciHandler := kithttp.NewServer(
		fibonacciEndpoint,
		decoder.HTTPFibonacci,
		utils.EncodeHTTPResponseWithData,
		h.options...,
	)

	r.Handle("/add", addHandler).Methods(http.MethodPost)
	r.Handle("/subtract", subtractHandler).Methods(http.MethodPost)
	r.Handle("/multiply", multiplyHandler).Methods(http.MethodPost)
	r.Handle("/divide", divideHandler).Methods(http.MethodPost)
	r.Handle("/fibonacci/{n}", fibonacciHandler).Methods(http.MethodGet)
}
