package hashmap

// re-implemented after deep dive
type bucketNode struct {
	key   string
	value string
	next  *bucketNode
}

type bucket struct {
	head *bucketNode
}

type HashTable struct {
	buckets []*bucket
	size    int
}

func (h *HashTable) NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*bucket, size),
		size:    size,
	}
}

func (h *HashTable) hash(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % len(h.buckets)
}

func (h *HashTable) Set(key, value string) {

	// create the address value using the hash func
	address := h.hash(key)
	// if the value at the address is nil then create a bucket for it
	if h.buckets[address] == nil {
		h.buckets[address] = &bucket{} // slice of key value pairs [ [key, value], [key, value] ]
	}

	// just overwrite if the key is found

	h.buckets[address].insert(key)
}

func (b *bucket) insert(k string) {
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

// Un-guided implementation // little messy overall

type Pair [2]string

type Hashmap struct {
	data [][]Pair
}

func NewHashmap(size int) *Hashmap {
	return &Hashmap{data: make([][]Pair, size)}
}

func (h *Hashmap) hash(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % len(h.data)
}

func (h *Hashmap) Set(key, value string) {

	// create the address value using the hash func
	address := h.hash(key)
	// if the value at the address is nil then create a bucket for it
	if h.data[address] == nil {
		h.data[address] = []Pair{} // slice of key value pairs [ [key, value], [key, value] ]
	}

	// just overwrite if the key is found
	for i, pair := range h.data[address] {
		if pair[0] == key {
			h.data[address][i][1] = value // overwrite existing value
			return
		}
	}

	// add the data at address index
	h.data[address] = append(h.data[address], Pair{key, value})
}

func (h *Hashmap) Get(key string) string {
	address := h.hash(key)

	// get the current bucket / array from the key address
	currentBucket := h.data[address]

	if currentBucket != nil {
		for i := 0; i < len(currentBucket); i++ {
			if currentBucket[i][0] == key { // check the key inside the current bucket
				return currentBucket[i][1] // return the value
			}
		}
	}
	return ""
}

func (h *Hashmap) Keys() []string {
	var keys []string
	for i := 0; i < len(h.data); i++ {
		bucket := h.data[i]
		if bucket == nil {
			continue
		}
		for _, pair := range bucket {
			keys = append(keys, pair[0])
		}
	}
	return keys
}
