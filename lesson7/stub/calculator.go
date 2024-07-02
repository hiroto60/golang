package stub

import (
	"fmt"
	"golang/lesson7/calc"
)

type Calculator struct {
	result int
	logger Logger
}

func NewCalculator(logger Logger) *Calculator {
	return &Calculator{logger: logger}
}

func (c *Calculator) Add(a int) {
	c.result = calc.Add(c.result, a)
	c.logger.Log(fmt.Sprintf("Add(%d): result = %d", a, c.result))
}

func (c *Calculator) Subtract(a int) {
	c.result = calc.Subtract(c.result, a)
	c.logger.Log(fmt.Sprintf("Subtract(%d): result = %d", a, c.result))
}

func (c *Calculator) Multiply(a int) {
	c.result = calc.Multiply(c.result, a)
	c.logger.Log(fmt.Sprintf("Multiply(%d): result = %d", a, c.result))
}

func (c *Calculator) Divide(a int) error {
	result, err := calc.Divide(c.result, a)
	if err != nil {
		c.logger.Log(fmt.Sprintf("Divide(%d): error = %v", a, err))
		return fmt.Errorf("division error: %w", err)
	}
	c.result = result
	c.logger.Log(fmt.Sprintf("Divide(%d): result = %d", a, c.result))
	return nil
}

func (c *Calculator) Result() int {
	return c.result
}

func (c *Calculator) Reset() {
	c.result = 0
	c.logger.Log("Reset(): result = 0")
}
