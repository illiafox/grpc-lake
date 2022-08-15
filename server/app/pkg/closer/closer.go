package closer

import (
	"io"
	"server/app/pkg/log"
	"sync"
)

type Closer struct {
	io.Closer
	Comment string
}

// Closers is a list of closers.
//
// Can be created with Closer{}
type Closers struct {
	mutex   sync.Mutex
	closers []Closer
}

func (c *Closers) Add(closer io.Closer, comment string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.closers = append(c.closers, Closer{
		Closer: closer, Comment: comment,
	})
}

func (c *Closers) Close(logger log.Logger) {
	if c.closers == nil {
		panic("closer is nil")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i := len(c.closers) - 1; i >= 0; i-- {
		closer := c.closers[i]
		logger.Info(closer.Comment)
		if err := closer.Close(); err != nil {
			logger.Error(closer.Comment, log.Error(err))
		}
	}
}
