package service

import "context"

// CalculatorService defines the interface for the calculator operations.
type CalculatorService interface {
	Add(ctx context.Context, a, b float64) (float64, error)
	Subtract(ctx context.Context, a, b float64) (float64, error)
	Multiply(ctx context.Context, a, b float64) (float64, error)
	Divide(ctx context.Context, a, b float64) (float64, error)
}
