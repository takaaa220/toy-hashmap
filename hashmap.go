package toyhashmap

func fnv1a32(data string) uint32 {
	bb := []byte(data)

	var hash uint32 = 0x811c9dc5
	const fnvPrime uint32 = 0x01000193

	for _, b := range bb {
		hash ^= uint32(b)
		hash *= fnvPrime
	}
	return hash
}

func nextPowerOf2(n uint32) uint32 {
	if n < 8 {
		return 8
	}

	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++
	return n
}

type HashMap[V any] struct {
	capacity uint32
	size     uint32
	buckets  []*bucket[V]
}

type bucket[V any] struct {
	entries []entry[V]
}

type entry[V any] struct {
	key   string
	value V
}

func NewHashMap[V any](capacity ...uint32) *HashMap[V] {
	cap := uint32(8)
	if len(capacity) == 1 {
		cap = nextPowerOf2(capacity[0])
	}

	buckets := make([]*bucket[V], cap)
	for i := range buckets {
		buckets[i] = &bucket[V]{}
	}

	return &HashMap[V]{
		capacity: cap,
		size:     0,
		buckets:  buckets,
	}
}

func (h *HashMap[V]) Put(key string, value V) {
	if h.shouldRehash() {
		h.resize()
	}

	i := h.bucketIndex(key)
	b := h.buckets[i]

	// Update the value if the key already exists
	for i, e := range b.entries {
		if e.key == key {
			b.entries[i].value = value
			return
		}
	}

	// Append the entry to the bucket
	b.entries = append(b.entries, entry[V]{key: key, value: value})
	h.size += 1
}

func (h *HashMap[V]) Delete(key string) {
	i := h.bucketIndex(key)
	b := h.buckets[i]

	// Remove the entry from the bucket
	for j, e := range b.entries {
		if e.key == key {
			b.entries[j] = b.entries[len(b.entries)-1]
			b.entries = b.entries[:len(b.entries)-1]
			h.size -= 1

			if len(b.entries) == 0 {
				b.entries = nil
			}

			return
		}
	}
}

func (h *HashMap[V]) Get(key string) (value V, ok bool) {
	var zero V

	i := h.bucketIndex(key)
	b := h.buckets[i]

	// Find the entry in the bucket
	for _, e := range b.entries {
		if e.key == key {
			return e.value, true
		}
	}

	return zero, false
}

func (h *HashMap[V]) bucketIndex(key string) uint32 {
	return fnv1a32(key) % h.capacity
}

func (h *HashMap[V]) shouldRehash() bool {
	return float64(h.size)/float64(h.capacity) > 0.75
}

func (h *HashMap[V]) resize() {
	newCapacity := h.capacity * 2

	newBuckets := make([]*bucket[V], newCapacity)
	for i := range newBuckets {
		newBuckets[i] = &bucket[V]{}
	}

	for _, b := range h.buckets {
		for _, e := range b.entries {
			i := fnv1a32(e.key) % newCapacity
			newBuckets[i].entries = append(newBuckets[i].entries, e)
		}
	}

	h.buckets = newBuckets
	h.capacity = newCapacity
}
