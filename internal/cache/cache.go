package cache

import (
	"time"

	"github.com/realnighthawk/bucky/cache"
	"github.com/realnighthawk/bucky/cache/inmem"
)

var (
	AdminTokenKey = "admin_token"
	AdminToken    = ""
)

// New creates a new cache instance
func New() (cache.Handler, error) {

	handler, err := inmem.New(inmem.Options{
		Expiration:      5 * time.Minute,
		CleanupInterval: 10 * time.Minute,
	})
	if err != nil {
		return nil, err
	}

	// Seed cache
	// TODO: Add seed data that are useful
	err = handler.Set(AdminTokenKey, AdminToken)
	if err != nil {
		return nil, err
	}

	return handler, nil
}
