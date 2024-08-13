package calculatorservice

import "context"

// Subtract implements service.CalculatorService.
func (c *calculatorService) Subtract(ctx context.Context, a float64, b float64) (float64, error) {
	return a - b, nil
}
