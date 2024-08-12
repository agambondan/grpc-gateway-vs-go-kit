package calculatorservice

import (
	"sync"

	"git.bluebird.id/firman.agam/go-kit/internal/service"
)

type calculatorService struct {
	value float64
	mutex sync.Mutex
}

func NewCalculatorService() service.CalculatorService {
	return &calculatorService{}
}
