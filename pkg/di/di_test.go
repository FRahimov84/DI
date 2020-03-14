package di

import (
	"testing"
)

func Test_container_Provide(t *testing.T) {
	type DSN string
	container := NewContainer()
	container.Provide(
		func() DSN { return DSN("Hello") },
		func(dsn DSN) string {return "world"},

	)

}