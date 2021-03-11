package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeHashTable(t *testing.T) {
	hashTable := MakeHashTable()

	assert.True(t, cap(hashTable.data) == 7)
	assert.True(t, len(hashTable.data) == 7)

	hashTable.Insert("baby")
	assert.True(t, hashTable.data[1].head.value == "baby")
	hashTable.Insert("rocket")
	assert.True(t, hashTable.data[4].head.value == "rocket")
	hashTable.Insert("zsd")
	assert.True(t, hashTable.data[1].head.next.value == "zsd")

	assert.True(t, hashTable.Search("baby"))
	assert.True(t, hashTable.Search("rocket"))
	assert.True(t, hashTable.Search("zsd"))
	assert.False(t, hashTable.Search("rer"))
	assert.False(t, hashTable.Search("ewx"))
	assert.False(t, hashTable.Search("caay"))

	hashTable.Remove("baby")
	assert.False(t, hashTable.Search("baby"))
	hashTable.Remove("rocket")
	assert.False(t, hashTable.Search("rocket"))
	hashTable.Remove("zsd")
	assert.False(t, hashTable.Search("zsd"))
}
