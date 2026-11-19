package main

import (
	"fmt"
	"strings"
)

type entry struct {
	name  string
	phone string
}

type HashTable struct {
	buckets [][]entry
}

func NewHashTable(size int) *HashTable {
	return &HashTable{buckets: make([][]entry, size)}
}

func (h *HashTable) hash(name string) int {
	sum := 0
	for i := 0; i < len(name); i++ {
		sum += int(name[i])
	}
	return sum % len(h.buckets)
}

func (h *HashTable) Put(name, phone string) {
	i := h.hash(name)

	for j := range h.buckets[i] {
		if h.buckets[i][j].name == name {
			h.buckets[i][j].phone = phone
			return
		}
	}

	h.buckets[i] = append(h.buckets[i], entry{name, phone})
}

func (h *HashTable) Get(name string) (string, bool) {
	i := h.hash(name)

	for _, e := range h.buckets[i] {
		if e.name == name {
			return e.phone, true
		}
	}

	return "", false
}

func (h *HashTable) Print() {
	for i, bucket := range h.buckets {
		parts := make([]string, 0, len(bucket))
		for _, e := range bucket {
			parts = append(parts, e.name+": "+e.phone)
		}

		if len(parts) == 0 {
			fmt.Printf("  slot %d: (empty)\n", i)
		} else {
			fmt.Printf("  slot %d: %s\n", i, strings.Join(parts, "  ->  "))
		}
	}
}

func explainHash(name string, size int) {
	codes := make([]string, 0, len(name))
	sum := 0
	for i := 0; i < len(name); i++ {
		codes = append(codes, fmt.Sprintf("%d", name[i]))
		sum += int(name[i])
	}

	fmt.Printf("  %-6s %s = %d, %d %% %d = slot %d\n",
		name+":", strings.Join(codes, "+"), sum, sum, size, sum%size)
}

func main() {
	const size = 7

	contacts := []entry{
		{"Alice", "555-1234"},
		{"Bob", "555-5678"},
		{"Carol", "555-8765"},
	}

	table := NewHashTable(size)

	fmt.Println("Hash function: sum of the character codes of the name % table size")
	for _, c := range contacts {
		explainHash(c.name, size)
		table.Put(c.name, c.phone)
	}

	fmt.Println()
	fmt.Println("Hash table with", size, "slots:")
	table.Print()

	fmt.Println()
	fmt.Println("Lookups:")
	for _, name := range []string{"Alice", "Bob", "Carol", "Dave"} {
		if phone, ok := table.Get(name); ok {
			fmt.Printf("  %-6s -> %s\n", name, phone)
		} else {
			fmt.Printf("  %-6s -> not found\n", name)
		}
	}

	fmt.Println()
	fmt.Println("Collision: Alice and Bob both land in slot 2.")
	fmt.Println("The table uses separate chaining: each slot holds a list of entries,")
	fmt.Println("and a lookup compares the names in that list to find the right one.")

	fmt.Println()
	fmt.Println("The same with Go's built-in map (a hash table):")
	phoneBook := map[string]string{}
	for _, c := range contacts {
		phoneBook[c.name] = c.phone
	}
	fmt.Println("  phoneBook[\"Bob\"] =", phoneBook["Bob"])
}
