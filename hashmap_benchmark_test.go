package toyhashmap

import (
	"strconv"
	"testing"
)

// 組み込みの map と自作 HashMap の比較対象データ数
const numEntries = 1_000_000
const initCap = 1_000

// ランダムなキーを作成するヘルパー関数
func generateKeys(n int) []string {
	keys := make([]string, n)
	for i := 0; i < n; i++ {
		keys[i] = "key" + strconv.Itoa(i)
	}
	return keys
}

// **1. Put のベンチマーク**
func BenchmarkHashMapPut(b *testing.B) {
	keys := generateKeys(numEntries)
	hm := NewHashMap[int](initCap)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm.Put(keys[i%numEntries], i)
	}
}

func BenchmarkBuiltinMapPut(b *testing.B) {
	keys := generateKeys(numEntries)
	m := make(map[string]int, initCap)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[keys[i%numEntries]] = i
	}
}

// **2. Get のベンチマーク**
func BenchmarkHashMapGet(b *testing.B) {
	keys := generateKeys(numEntries)
	hm := NewHashMap[int](initCap)
	for i := 0; i < numEntries; i++ {
		hm.Put(keys[i], i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = hm.Get(keys[i%numEntries])
	}
}

func BenchmarkBuiltinMapGet(b *testing.B) {
	keys := generateKeys(numEntries)
	m := make(map[string]int, initCap)
	for i := 0; i < numEntries; i++ {
		m[keys[i]] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m[keys[i%numEntries]]
	}
}

// **3. Delete のベンチマーク**
func BenchmarkHashMapDelete(b *testing.B) {
	keys := generateKeys(numEntries)
	hm := NewHashMap[int](initCap)
	for i := 0; i < numEntries; i++ {
		hm.Put(keys[i], i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm.Delete(keys[i%numEntries])
	}
}

func BenchmarkBuiltinMapDelete(b *testing.B) {
	keys := generateKeys(numEntries)
	m := make(map[string]int, initCap)
	for i := 0; i < numEntries; i++ {
		m[keys[i]] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		delete(m, keys[i%numEntries])
	}
}
