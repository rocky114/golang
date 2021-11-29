package cache

import "testing"

func TestGet(t *testing.T) {
	cache := NewCache(200)
	cache.Set("name", "rocky")

	if cache.Get("name") != "rocky" {
		t.Error("cache error")
	}
}
