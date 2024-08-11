package calculatorservice

import "context"

// Add implements service.CalculatorService.
func (c *calculatorService) Add(ctx context.Context, a float64, b float64) (float64, error) {
	userID := ctx.Value("userID").(string)

	c.mutex.Lock()

	c.value[userID] += a + b

	c.mutex.Unlock()

	return c.value[userID], nil
}
