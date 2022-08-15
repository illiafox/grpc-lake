package encode

import (
	"server/app/internal/domain/entity"
	"time"
)

//go:generate msgp

var _ = entity.Item(Item{})

type Item struct {
	Name        string    `msgp:"name"`
	Data        []byte    `msgp:"data"`
	Created     time.Time `msgp:"created"`
	Description string    `msgp:"desc"`
}

func (i Item) ToEntity() entity.Item {
	return entity.Item(i)
}
