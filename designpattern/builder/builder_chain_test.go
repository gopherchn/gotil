package builder

import (
	"testing"
)

func TestServerBuilder_Create(t *testing.T) {

	b := &ServerBuilder{}
	s1 := b.Create("127.0.1", 9999).Build()
	t.Log("s1", s1)

	server := b.Create("127.0.0.1", 9999).
		WithMaxConns(10).
		WithProtocol("tcp").
		Build()
	t.Logf("%+v, Conf: %+v\n", server, server.Conf)

}
