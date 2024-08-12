package calculatorservice

import "context"

// Add implements service.CalculatorService.
func (c *calculatorService) Fibonacci(ctx context.Context, n int32) (int32, error) {
	return c.fibonacci(ctx, n), nil
}

// fibonacci is a helper function that calculates the nth Fibonacci number.
func (c *calculatorService) fibonacci(ctx context.Context, n int32) int32 {
	if n <= 1 {
		return n
	}

	var a, b int32 = 0, 1
	for i := int32(2); i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
