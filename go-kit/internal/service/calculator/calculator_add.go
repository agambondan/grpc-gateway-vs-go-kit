package calculatorservice

import "context"

// Add implements service.CalculatorService.
func (c *calculatorService) Add(ctx context.Context, a float64, b float64) (float64, error) {
	return a + b, nil
}
