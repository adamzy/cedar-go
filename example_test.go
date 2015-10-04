package cedar_test

import (
	"fmt"

	"github.com/adamzy/cedar-go"
)

var trie *cedar.Cedar

func Example() {
	trie = cedar.New()

	// Insert key-value pairs.
	// The order of insertion is not important.
	trie.Insert([]byte("How many"), 0)
	trie.Insert([]byte("How many loved"), 1)
	trie.Insert([]byte("How many loved your moments"), 2)
	trie.Insert([]byte("How many loved your moments of glad grace"), 3)
	trie.Insert([]byte("姑苏"), 4)
	trie.Insert([]byte("姑苏城外"), 5)
	trie.Insert([]byte("姑苏城外寒山寺"), 6)

	// Get the associated value of a key directly.
	value, _ := trie.Get([]byte("How many loved your moments of glad grace"))
	fmt.Println(value)

	// Or, use `jump` to get the id of the trie node fist,
	id, _ := trie.Jump([]byte("How many loved your moments"), 0)
	// then get the key and the value.
	key, _ := trie.Key(id)
	value, _ = trie.Value(id)
	fmt.Printf("%d\t%s:%v\n", id, key, value)

	// Output:
	// 3
	// 281	How many loved your moments:2
}

func Example_prefixMatch() {
	fmt.Println("id\tkey:value")
	for _, id := range trie.PrefixMatch([]byte("How many loved your moments of glad grace"), 0) {
		key, _ := trie.Key(id)
		value, _ := trie.Value(id)
		fmt.Printf("%d\t%s:%v\n", id, key, value)
	}
	// Output:
	// id	key:value
	// 262	How many:0
	// 268	How many loved:1
	// 281	How many loved your moments:2
	// 296	How many loved your moments of glad grace:3
}

func Example_prefixPredict() {
	fmt.Println("id\tkey:value")
	for _, id := range trie.PrefixPredict([]byte("姑苏"), 0) {
		key, _ := trie.Key(id)
		value, _ := trie.Value(id)
		fmt.Printf("%d\t%s:%v\n", id, key, value)
	}
	// Output:
	// id	key:value
	// 303	姑苏:4
	// 309	姑苏城外:5
	// 318	姑苏城外寒山寺:6
}

func Example_saveAndLoad() {
	trie.SaveToFile("cedar.gob", "gob")
	trie.SaveToFile("cedar.json", "json")

	trie.LoadFromFile("cedar.gob", "gob")
	trie.LoadFromFile("cedar.json", "json")
}
