package hashtable

// HashTable has 10 size
type HashTable struct {
	data [7]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	value string
	next  *bucketNode
}

// Insert value
func (h *HashTable) Insert(value string) {
	index := hash(value)
	h.data[index].insert(value)
}

// Search value
func (h *HashTable) Search(value string) bool {
	index := hash(value)
	return h.data[index].search(value)
}

// Remove
func (h *HashTable) Remove(value string) {
	index := hash(value)
	h.data[index].remove(value)
}

// insert
func (b *bucket) insert(value string) {
	if b.head == nil {
		b.head = &bucketNode{value: value}
	} else {
		curr := b.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &bucketNode{value: value}
	}
}

// search
func (b *bucket) search(value string) bool {
	node := b.head

	for node != nil {
		if node.value == value {
			return true
		}
		node = node.next
	}

	return false
}

// remove
func (b *bucket) remove(value string) {
	if b.head.value == value {
		b.head = b.head.next
		return
	}

	prev := b.head

	for prev.next != nil {
		if prev.next.value == value {
			prev.next = prev.next.next
		} else {
			prev = prev.next
		}
	}
}

// hash
func hash(value string) int {
	sum := 0

	for _, v := range value {
		sum += int(v)
	}

	return sum % 7
}

// MakeHashTable returns HashTable
func MakeHashTable() *HashTable {
	hashTable := &HashTable{}

	for i := range hashTable.data {
		hashTable.data[i] = &bucket{}
	}

	return hashTable
}
