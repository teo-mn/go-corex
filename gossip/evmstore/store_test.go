package evmstore

import (
	"time"

	"github.com/corex-mn/corex-base/kvdb"
	"github.com/corex-mn/corex-base/kvdb/memorydb"
)

func cachedStore() *Store {
	cfg := LiteStoreConfig()

	return NewStore(memorydb.New(), cfg)
}

func nonCachedStore() *Store {
	cfg := StoreConfig{}

	return NewStore(memorydb.New(), cfg)
}

func withDelay(db kvdb.DropableStore) kvdb.DropableStore {
	mem, ok := db.(*memorydb.Database)
	if ok {
		mem.SetDelay(time.Millisecond)

	}

	return db
}
