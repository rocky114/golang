package cache

import (
	"testing"
)

func TestGet(t *testing.T) {
	cache := NewCache(20)
	cache.Set("name", "rocky")
	cache.Set("age", 1)

	t.Log(cache.Len(), 3)

	if cache.Get("name") != "rocky" {
		t.Error("cache error")
	}
}
