package builder

import (
	"crypto/tls"
	"time"
)


// builder chain 每次配置一个对象就返回自身进行链式调用
// 缺点就是仍旧要多嵌入一个builder结构体
type Config struct {
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type Server struct {
	Addr string
	Port int
	Conf *Config
}

type ServerBuilder struct {
	Server
}

func (b *ServerBuilder) Create(addr string, port int) *ServerBuilder {
	b.Server.Addr = addr
	b.Server.Port = port
	b.Server.Conf = &Config{}
	return b
}

func (b *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	b.Server.Conf.Protocol = protocol
	return b
}

func (b *ServerBuilder) WithMaxConns(maxConns int) *ServerBuilder {
	
	b.Server.Conf.MaxConns = maxConns
	return b
}

func (b *ServerBuilder) WithTimeout(timeout time.Duration) *ServerBuilder {
	b.Server.Conf.Timeout = timeout
	return b
}

func (b *ServerBuilder) WithTLS(tls *tls.Config) *ServerBuilder {
	b.Server.Conf.TLS = tls
	return b
}

func (b* ServerBuilder) Build() *Server {
	return &b.Server	
}