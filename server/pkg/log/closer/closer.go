package closer

import (
	"sync"

	"go.uber.org/zap"
)

type Closer struct {
	close   func() error
	comment string
}

// Closers is a list of closers.
//
// Can be created with Closer{}
type Closers struct {
	mutex   sync.Mutex
	closers []Closer
}

func (c *Closers) Add(closer func() error, comment string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.closers = append(c.closers, Closer{
		close: closer, comment: comment,
	})
}

func (c *Closers) Close(logger *zap.Logger) {
	if c.closers == nil {
		panic("close is nil")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i := len(c.closers) - 1; i >= 0; i-- {
		closer := c.closers[i]
		logger.Info(closer.comment)
		if err := closer.close(); err != nil {
			logger.Error(closer.comment, zap.Error(err))
		}
	}
}
