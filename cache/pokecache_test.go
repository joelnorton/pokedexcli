package pokecache

import (
	"bytes"
	"testing"
)

func TestPokeCache(t *testing.T) {

	cache := NewCache()
	cache.addCache("t1", []byte("Hello World"))
	t1, success := cache.getCache("t1")
	if !success || !bytes.Equal(t1, []byte("Hello World")) {
		t.Error("t1 failed")
	}
	cache.addCache("t2", []byte("Who?"))
	t2, success := cache.getCache("t2")
	if !success || !bytes.Equal(t2, []byte("Who?")) {
		t.Error("t1 failed")
	}
	cache.addCache("t3", []byte("Goodbye World"))
	t3, success := cache.getCache("t3")
	if !success || !bytes.Equal(t3, []byte("Goodbye World")) {
		t.Error("t3 failed")
	}
	cache.addCache("t4", []byte("Blah blah blah"))
	t1, success = cache.getCache("t1")
	if !success || !bytes.Equal(t1, []byte("Hello World")) {
		t.Error("t4 failed")
	}
}
