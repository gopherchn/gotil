# 依赖
- go get github.com/stretchr/testify


### singleton(单例模式)
## 饿汉式
- 类加载的时候 Instance实例已经创建好了
- 创建过程是线程安全的
- 不支持延加载，初始化的时间会比较长

## 懒汉式
- 获取Instance的时候才创建实例
- 创建过程需要加锁
- 支持延迟加载

### benchmark

- go test --benchmem -bench="." -v

```
pkg: github.com/gopherchn/gotil/designpattern/singleton
BenchmarkInstanceParallel
BenchmarkInstanceParallel-8             1000000000               0.1690 ns/op          0 B/op          0 allocs/op
BenchmarkLazyInstanceParallel
BenchmarkLazyInstanceParallel-8         1000000000               0.6700 ns/op          0 B/op          0 allocs/op
```