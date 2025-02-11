package toyhashmap

import (
	"strconv"
	"testing"
)

func TestPutAndGet(t *testing.T) {
	hm := NewHashMap[int]()

	hm.Put("key1", 10)
	hm.Put("key2", 20)
	hm.Put("key3", 30)

	if val, ok := hm.Get("key1"); !ok || val != 10 {
		t.Errorf("Expected 10, got %v", val)
	}
	if val, ok := hm.Get("key2"); !ok || val != 20 {
		t.Errorf("Expected 20, got %v", val)
	}
	if val, ok := hm.Get("key3"); !ok || val != 30 {
		t.Errorf("Expected 30, got %v", val)
	}

	if _, ok := hm.Get("unknown"); ok {
		t.Errorf("Expected Get() to return false for missing key")
	}
}

func TestPutOverwrite(t *testing.T) {
	hm := NewHashMap[int]()
	hm.Put("key1", 10)
	hm.Put("key1", 50)

	if val, ok := hm.Get("key1"); !ok || val != 50 {
		t.Errorf("Expected 50, got %v", val)
	}
}

func TestDelete(t *testing.T) {
	hm := NewHashMap[int]()
	hm.Put("key1", 100)
	hm.Put("key2", 200)

	hm.Delete("key1")

	if _, ok := hm.Get("key1"); ok {
		t.Errorf("Expected Get() to return false for deleted key")
	}

	if val, ok := hm.Get("key2"); !ok || val != 200 {
		t.Errorf("Expected 200, got %v", val)
	}
}

func TestResize(t *testing.T) {
	hm := NewHashMap[int](1)

	for i := 1; i <= 100; i++ {
		hm.Put("key"+strconv.Itoa(i), i)
	}

	if val, ok := hm.Get("key1"); !ok || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	if val, ok := hm.Get("key2"); !ok || val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}
	if val, ok := hm.Get("key3"); !ok || val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}
	if val, ok := hm.Get("key4"); !ok || val != 4 {
		t.Errorf("Expected 4, got %v", val)
	}
}

func TestCollision(t *testing.T) {
	hm := NewHashMap[int](1)

	for i := 1; i <= 100; i++ {
		hm.Put("key"+strconv.Itoa(i), i)
	}

	if val, ok := hm.Get("key1"); !ok || val != 1 {
		t.Errorf("Expected 1, got %v", 1)
	}
	if val, ok := hm.Get("key2"); !ok || val != 2 {
		t.Errorf("Expected 2, got %v", 2)
	}
}
