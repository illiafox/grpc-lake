package entity

import (
	"errors"
	"time"
)

type Item struct {
	Name        string
	Data        []byte
	Created     time.Time
	Description string
}

var ErrItemNotFound = errors.New("item not found")
