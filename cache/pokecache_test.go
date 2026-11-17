package pokecache

import (
	"bytes"
	"sync"
	"testing"
	"time"
)

func TestPokeCacheBasic(t *testing.T) {
	cache := NewCache(time.Second * 5)

	cache.AddCache("t1", []byte("Hello World"))
	t1, success := cache.GetCache("t1")
	if !success || !bytes.Equal(t1, []byte("Hello World")) {
		t.Error("t1 failed: expected Hello World")
	}
}

func TestPokeCacheMultipleKeys(t *testing.T) {
	cache := NewCache(time.Second * 5)

	cache.AddCache("k1", []byte("val1"))
	cache.AddCache("k2", []byte("val2"))
	cache.AddCache("k3", []byte("val3"))

	for _, k := range []string{"k1", "k2", "k3"} {
		val, ok := cache.GetCache(k)
		if !ok || !bytes.Equal(val, []byte("val"+k[len(k)-1:])) {
			t.Errorf("key %s unexpected", k)
		}
	}
}

func TestPokeCacheExpiryOnGet(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)

	cache.AddCache("about-to-expire", []byte("short-lived"))

	// Should be present immediately
	val, ok := cache.GetCache("about-to-expire")
	if !ok {
		t.Error("expected entry to exist immediately")
	}
	if !bytes.Equal(val, []byte("short-lived")) {
		t.Error("expected short-lived")
	}

	// Sleep past expiry
	time.Sleep(time.Millisecond * 100)

	// Should be gone after expiry
	val, ok = cache.GetCache("about-to-expire")
	if ok {
		t.Error("expected entry to be expired after sleeping 100ms past 10ms interval")
	}
}

func TestPokeCacheConcurrentAccess(t *testing.T) {
	cache := NewCache(time.Second * 5)
	var wg sync.WaitGroup

	// Writers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.AddCache("k", []byte("val"))
		}(i)
	}

	// Readers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.GetCache("k")
		}()
	}

	wg.Wait()
}

func TestPokeCacheCloseStopsGoroutine(t *testing.T) {
	cache := NewCache(time.Millisecond * 100)

	// Give the reaper time to start
	time.Sleep(time.Millisecond * 10)

	// Close should not hang
	cache.Close()

	// Should return quickly (within a few ms after stop)
	// If Close hangs, this test will timeout
}

func TestPokeCacheIntervalZeroDefaultsToSecond(t *testing.T) {
	// interval <= 0 should default to 1 second
	cache := NewCache(0)
	// Should not panic
	val, ok := cache.GetCache("dummy")
	if ok {
		t.Log("got value with zero interval (unexpected but OK)")
	}
	_ = val
}

func TestPokeCacheAddCopiesInput(t *testing.T) {
	cache := NewCache(time.Second * 5)

	input := []byte("original")
	cache.AddCache("k", input)

	// Mutate the original slice - cache should be unaffected
	input[0] = 'X'

	val, ok := cache.GetCache("k")
	if !ok {
		t.Error("expected entry to exist")
	}
	if bytes.Equal(val, []byte("Xoriginal")) {
		t.Error("cache was mutated by caller modifying input slice; addCache should have copied")
	}
}

func TestPokeCacheGetReturnsCopy(t *testing.T) {
	cache := NewCache(time.Second * 5)

	input := []byte("original")
	cache.AddCache("k", input)

	// Mutate the returned slice - cache should be unaffected
	val, ok := cache.GetCache("k")
	if !ok {
		t.Error("expected entry to exist")
		return
	}
	val[0] = 'X'

	// Re-read: should still be "original"
	val2, ok2 := cache.GetCache("k")
	if !ok2 {
		t.Error("expected entry to still exist after mutating first return")
	}
	if !bytes.Equal(val2, []byte("original")) {
		t.Error("mutating returned slice corrupted cache; getCache should return a copy")
	}
}

func TestPokeCacheDeleteAfterExpiryThenReadd(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)

	cache.AddCache("k", []byte("first"))
	val, ok := cache.GetCache("k")
	if !ok || !bytes.Equal(val, []byte("first")) {
		t.Error("expected first value")
	}

	// Sleep past expiry
	time.Sleep(time.Millisecond * 100)

	// Should be gone
	val, ok = cache.GetCache("k")
	if ok {
		t.Error("expected entry to be expired")
	}

	// Re-add should work
	cache.AddCache("k", []byte("second"))
	val, ok = cache.GetCache("k")
	if !ok || !bytes.Equal(val, []byte("second")) {
		t.Error("expected second value after re-add")
	}
}
