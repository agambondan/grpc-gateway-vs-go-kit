package calculatorservice

import "context"

// Divide implements service.CalculatorService.
func (c *calculatorService) Divide(ctx context.Context, a float64, b float64) (float64, error) {
	return a / b, nil
}
