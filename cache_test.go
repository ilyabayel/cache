package cache_test

import (
	"testing"

	"github.com/ilyabayel/cache"
)

func TestSetAndGet(t *testing.T) {
	c := cache.New()

	c.Set("key", "value")

	if v := c.Get("key"); v != "value" {
		t.Error("Set failed")
	}
}

func TestDelete(t *testing.T) {
	c := cache.New()

	c.Set("key", "value")

	if v := c.Get("key"); v != "value" {
		t.Error("Set failed")
	}

	c.Delete("key")

	if v := c.Get("key"); v != nil {
		t.Error("Delete failed")
	}
}

func TestWithDifferentTypes(t *testing.T) {
	c := cache.New()

	c.Set("string", "42")
	c.Set("int", 42)

	if v := c.Get("string"); v != "42" {
		t.Error("Set string failed")
	}

	if v := c.Get("int"); v != 42 {
		t.Error("Set int failed")
	}
}
