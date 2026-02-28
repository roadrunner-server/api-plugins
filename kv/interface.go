package kv

import "context"

/*
  string key = 1;
  bytes value = 2;
  // RFC 3339
  string timeout = 3;
*/

// Item represents a single KV entry
type Item interface {
	Key() string
	Value() []byte
	Timeout() string
}

// Storage represents single abstract storage.
type Storage interface {
	// Has checks if value exists.
	Has(ctx context.Context, keys ...string) (map[string]bool, error)

	// Get loads value content into a byte slice.
	Get(ctx context.Context, key string) ([]byte, error)

	// MGet loads content of multiple values
	// Returns the map with existing keys and associated values
	MGet(ctx context.Context, keys ...string) (map[string][]byte, error)

	// Set used to upload item to KV with TTL
	// 0 value in TTL means no TTL
	Set(ctx context.Context, items ...Item) error

	// MExpire sets the TTL for multiply keys
	MExpire(ctx context.Context, items ...Item) error

	// TTL return the rest time to live for provided keys
	// Not supported for the memcached
	TTL(ctx context.Context, keys ...string) (map[string]string, error)

	// Clear clean the entire storage
	Clear(ctx context.Context) error

	// Delete one or multiple keys.
	Delete(ctx context.Context, keys ...string) error

	// Stop the storage driver
	Stop(ctx context.Context)
}

// Constructor provides storage based on the config
type Constructor interface {
	// KvFromConfig provides Storage based on the config key
	KvFromConfig(ctx context.Context, key string) (Storage, error)
	Name() string
}
