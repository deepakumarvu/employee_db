package simpledb

import (
	"sync"

	orderedmap "github.com/wk8/go-ordered-map"
)

type Database[I comparable, T any] struct {
	data  *orderedmap.OrderedMap
	mutex sync.RWMutex
}
