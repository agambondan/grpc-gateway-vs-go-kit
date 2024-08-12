package calculatorservice

import "context"

// Add implements service.CalculatorService.
func (c *calculatorService) Add(ctx context.Context, a float64, b float64) (float64, error) {
	c.mutex.Lock()

	c.value += a + b

	c.mutex.Unlock()

	return c.value, nil
}
