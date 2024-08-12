package calculatorservice

import "context"

// Multiply implements service.CalculatorService.
func (c *calculatorService) Multiply(ctx context.Context, a float64, b float64) (float64, error) {
	c.mutex.Lock()

	c.value = a * b

	c.mutex.Unlock()

	return c.value, nil
}
