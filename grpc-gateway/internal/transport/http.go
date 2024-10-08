package transport

import (
	"context"
	"net/http"

	"git.bluebird.id/firman.agam/grpc-gateway/internal/transport/middleware"
	"git.bluebird.id/firman.agam/grpc-gateway/internal/utils"
	"git.bluebird.id/promo/packages/zaplog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

type HTTPHandler interface {
	Register(ctx context.Context, gwmux *runtime.ServeMux, conn *grpc.ClientConn) error
}

type HTTPServer interface {
	RegisterGRPCGatewayHandler(service HTTPHandler) error
	RegisterHTTPHandler(endpoint string, handler http.Handler)
	RegisterHTTPHandleFunc(endpoint string, handler func(w http.ResponseWriter, r *http.Request))
	Start() error
	Stop(ctx context.Context) error
}

type httpServer struct {
	runtimeMux *runtime.ServeMux
	grpcPort   string
	httpPort   string
	httpMux    *http.ServeMux
	server     *http.Server
}

func NewHTTPServer(httpPort, grpcPort string) HTTPServer {
	httpMux := http.NewServeMux()
	// runtime MIMEWildcard is annotations in body in http proto
	// can custom MIME with any string but annotations in http proto
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.HTTPBodyMarshaler{
				Marshaler: &runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						Multiline:         true,
						Indent:            " ",
						UseProtoNames:     true,
						EmitUnpopulated:   true,
						EmitDefaultValues: true,
					},
					UnmarshalOptions: protojson.UnmarshalOptions{
						DiscardUnknown: true,
					},
				},
			},
		),
		// add more custom mux options if you needed
		runtime.WithErrorHandler(utils.CustomErrorHandler),
		runtime.WithIncomingHeaderMatcher(utils.CustomMatcher),
		runtime.WithForwardResponseOption(utils.CustomResponseHandler),
		runtime.WithMetadata(utils.CustomMetadata),
	)

	return &httpServer{
		grpcPort:   grpcPort,
		httpPort:   httpPort,
		runtimeMux: mux,
		httpMux:    httpMux,
		server: &http.Server{
			Addr:    ":" + httpPort,
			Handler: httpMux,
		},
	}
}

func (s *httpServer) RegisterGRPCGatewayHandler(service HTTPHandler) error {
	conn, err := grpc.NewClient(
		"0.0.0.0:"+s.grpcPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zaplog.WithContext(context.Background()).Info("failed to dial server", zap.Error(err))
		return err
	}

	return service.Register(context.Background(), s.runtimeMux, conn)
}

// RegisterHTTPHandler
// this function will add your custom handler if your proto not implement that thing
func (s *httpServer) RegisterHTTPHandler(endpoint string, handler http.Handler) {
	s.httpMux.Handle(endpoint, handler)
}

// RegisterHTTPHandler
// this function will add your custom handler if your proto not implement that thing
func (s *httpServer) RegisterHTTPHandleFunc(endpoint string, handler func(w http.ResponseWriter, r *http.Request)) {
	s.httpMux.HandleFunc(endpoint, handler)
}

func (s *httpServer) Start() error {
	s.httpMux.Handle("/api/v1/", http.StripPrefix("/api/v1", corsMiddleware(middleware.HTTPMiddleware(s.runtimeMux))))

	zaplog.WithContext(context.Background()).Info("serving grpc-gateway on port :" + s.httpPort)

	return s.server.ListenAndServe()
}

func (s *httpServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
