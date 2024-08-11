package healthhttp

import (
	"encoding/json"
	"net/http"

	"git.bluebird.id/promo/packages/zaplog"
	"go.uber.org/zap"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

type HealthHandler struct {
	healthServer *health.Server
}

func NewHealthHandler(server *health.Server) *HealthHandler {
	return &HealthHandler{
		healthServer: server,
	}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := zaplog.WithContext(nil)
	defer logger.Sync()

	healthStatus, err := h.healthServer.Check(r.Context(), &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		logger.Info("failed to check server", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := &HealthCheckResponse{Status: healthStatus.GetStatus().String()}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Info("failed to encode json", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
