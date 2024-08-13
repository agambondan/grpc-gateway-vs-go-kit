package env

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"git.bluebird.id/promo/packages/zaplog"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

const (
	grpcPortEnv = "GRPC_PORT"
	httpPortEnv = "HTTP_PORT"
)

var (
	GRPCPort string
	HTTPPort string
)

func LoadEnv() {
	logger := zaplog.WithContext(context.Background())
	var err error

	GRPCPort, err = getEnvString(grpcPortEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	HTTPPort, err = getEnvString(httpPortEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}
}

func getEnvString(env string) (string, error) {
	res := os.Getenv(env)
	if res == "" {
		return res, fmt.Errorf("env %s is empty", env)
	}
	return res, nil
}

func getEnvInt(env string) (int, error) {
	res := os.Getenv(env)
	if res == "" {
		return 0, fmt.Errorf("env %s is empty", env)
	}

	resInt, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}

	return resInt, nil
}

func getEnvInt64(env string) (int64, error) {
	res := os.Getenv(env)
	if res == "" {
		return 0, fmt.Errorf("env %s is empty", env)
	}

	resInt, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}

	return int64(resInt), nil
}

func getEnvFloat64(env string) (float64, error) {
	res := os.Getenv(env)
	if res == "" {
		return 0, fmt.Errorf("env %s is empty", env)
	}

	resFloat64, err := strconv.ParseFloat(res, 64)
	if err != nil {
		return 0, err
	}

	return resFloat64, nil
}

func getEnvBool(env string) (bool, error) {
	res := os.Getenv(env)
	if res == "" {
		return false, fmt.Errorf("env %s is empty", env)
	}

	resBool, err := strconv.ParseBool(res)
	if err != nil {
		return false, err
	}

	return resBool, nil
}
