package builder

// Option解决的问题不需要包一个额外包装一个builder
// Option
type ResourcePoolConfigOption func(*ResourcePoolConfig)

func WithMaxIdle(maxIdle int) ResourcePoolConfigOption {
	return func(o *ResourcePoolConfig) {
		o.maxIdle = maxIdle
	}
}

func WithMinIdle(minIdle int) ResourcePoolConfigOption {
	return func(o *ResourcePoolConfig) {
		o.minIdle = minIdle
	}
}

func WithMaxPool(maxTotal int) ResourcePoolConfigOption {
	return func(o *ResourcePoolConfig) {
		o.maxTotal = maxTotal
	}
}

func NewResourcePoolConfig(options ...func(*ResourcePoolConfig)) (*ResourcePoolConfig, error) {
	cfg := &ResourcePoolConfig{}
	for _, option := range options {
		option(cfg)
	}

	// 基础数据初始化完成, 可以进行额外判断
	if cfg.maxIdle == 0 {
		cfg.maxIdle = defaultMaxIdle
	}
	if cfg.minIdle == 0 {
		cfg.minIdle = defaultMinIdle
	}
	if cfg.maxTotal == 0 {
		cfg.maxTotal = defaultMaxTotal
	}
	// maxTotal和minTotal判断

	return cfg, nil
}
