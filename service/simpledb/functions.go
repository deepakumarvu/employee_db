package simpledb

import orderedmap "github.com/wk8/go-ordered-map"

func (db *Database[I, T]) Init() *Database[I, T] {
	db.data = orderedmap.New()
	return db
}

func (db *Database[I, T]) GetItem(key I) (T, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	i, p := db.data.Get(key)
	var zeroVal T
	if i == nil {
		return zeroVal, p
	}
	return i.(T), p
}

func (db *Database[I, T]) SetItem(key I, value T) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, present := db.data.Get(key); present {
		return KeyAlreadyPresent
	}
	db.data.Set(key, value)
	return nil
}

func (db *Database[I, T]) UpdateItem(key I, value T) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, present := db.data.Get(key); !present {
		return KeyAbsent
	}
	db.data.Set(key, value)
	return nil
}

func (db *Database[I, T]) DeleteItem(key I) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, present := db.data.Get(key); !present {
		return KeyAbsent
	}
	db.data.Delete(key)
	return nil
}

func (db *Database[I, T]) GetItems(LastEvalKeyID I, numItems int) ([]T, I, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	var lastItem *orderedmap.Pair
	var items []T
	var i int
	var lastID I
	var zeroVal I

	// Default value for LastEvalKeyID if it is not provided
	if LastEvalKeyID == zeroVal {
		lastItem = db.data.Oldest()
		// If lastItem is not nil, set it as the first record to return
		if lastItem != nil {
			items = []T{lastItem.Value.(T)}
			i = 1
			lastID = lastItem.Key.(I)
		}
	} else {
		// Get the item with the given LastEvalKeyID
		lastItem = db.data.GetPair(LastEvalKeyID)
		if lastItem == nil {
			// If LastEvalKeyID is not present, return an error
			return nil, LastEvalKeyID, InvalidLastEvalKeyID
		}
	}
	// Get the next batch of items
	for ; i < numItems && lastItem != nil && lastItem.Next() != nil; i++ {
		lastItem = lastItem.Next()
		lastID = lastItem.Key.(I)
		items = append(items, lastItem.Value.(T))
	}

	if numItems > i {
		// If numItems is greater than the number of items returned, set the lastID to the zero value
		lastID = zeroVal
	}

	return items, lastID, nil
}
