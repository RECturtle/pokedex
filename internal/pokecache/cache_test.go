package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestGetMissing(t *testing.T) {
	cache := NewCache(5 * time.Second)
	_, ok := cache.Get("nonexistent")
	if ok {
		t.Errorf("expected ok to be false for missing key")
	}
}

func TestOverwrite(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("key", []byte("original"))
	cache.Add("key", []byte("updated"))

	val, ok := cache.Get("key")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	if string(val) != "updated" {
		t.Errorf("expected updated value, got %s", string(val))
	}
}

func TestMultipleKeys(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("key1", []byte("val1"))
	cache.Add("key2", []byte("val2"))

	val1, ok := cache.Get("key1")
	if !ok || string(val1) != "val1" {
		t.Errorf("expected val1 for key1")
	}
	val2, ok := cache.Get("key2")
	if !ok || string(val2) != "val2" {
		t.Errorf("expected val2 for key2")
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
