package builder

import "fmt"

var (
	defaultMaxTotal = 10
	defaultMaxIdle  = 9
	defaultMinIdle  = 1
)

type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// type ResourcePoolConfigBuilder struct {
// 	name     string
// 	maxTotal int
// 	maxIdle  int
// 	minIdle  int
// }
type ResourcePoolConfigBuilder struct {
	ResourcePoolConfig
}

func (c *ResourcePoolConfigBuilder) SetName(name string) error {
	if c.ResourcePoolConfig.name == "" {
		return fmt.Errorf("name cant't be empty")
	}
	c.ResourcePoolConfig.name = name
	return nil
}

func (c *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if c.ResourcePoolConfig.minIdle < 0 {
		return fmt.Errorf("minIdle must big than 0, current input %d", minIdle)
	}
	c.ResourcePoolConfig.maxIdle = minIdle
	return nil
}

func (c *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if c.ResourcePoolConfig.maxIdle < 0 {
		return fmt.Errorf("maxIdle must big than 0, current input %d", maxIdle)
	}
	c.ResourcePoolConfig.maxIdle = maxIdle
	return nil
}

func (c *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if c.ResourcePoolConfig.maxTotal <= 0 {
		return fmt.Errorf("maxTotal cant <= 0, input % d", maxTotal)
	}
	c.ResourcePoolConfig.maxIdle = maxTotal
	return nil
}

func (c *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if c.ResourcePoolConfig.name == "" {
		return nil, fmt.Errorf("name cant'be empty: %s ", c.name)
	}

	if c.ResourcePoolConfig.maxIdle == 0 {
		c.ResourcePoolConfig.maxIdle = defaultMaxIdle
	}

	if c.ResourcePoolConfig.minIdle == 0 {
		c.ResourcePoolConfig.minIdle = defaultMinIdle
	}

	if c.ResourcePoolConfig.maxTotal == 0 {
		c.ResourcePoolConfig.maxTotal = defaultMaxTotal
	}

	if c.ResourcePoolConfig.maxTotal < c.ResourcePoolConfig.maxIdle {
		return nil, fmt.Errorf("maxIdle  %d ,cant < maxTotal: %d,", c.maxIdle, c.maxTotal)
	}

	if c.ResourcePoolConfig.maxIdle < c.ResourcePoolConfig.minIdle {
		return nil, fmt.Errorf("maxIdle: %d cant't  < minIdle : %d ", c.maxIdle, c.minIdle)
	}

	return &ResourcePoolConfig{
		maxTotal: c.maxTotal,
		maxIdle:  c.maxIdle,
		minIdle:  c.minIdle,
		name:     c.name,
	}, nil
}
