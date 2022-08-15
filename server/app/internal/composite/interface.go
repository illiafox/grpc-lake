package composite

import "io"

// Composite represents a Client and its Close method.
type Composite[T any] interface {
	io.Closer
	Client() T
}
