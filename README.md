# go-hashtable

A simple contact manager built on a hash table in Go: name → phone number.

## Task

Imagine you are building a simple contact manager application. You want to store people's phone
numbers so that you can quickly look them up by their unique names.

1. Design a hash table structure to store the contact information where:
   - the key is the person's name (a unique string);
   - the value is their phone number (a string of digits).
2. Explain how the hash function would work to place each contact into the hash table.
3. Suppose you have the following contacts to insert:
   - Alice: 555-1234
   - Bob: 555-5678
   - Carol: 555-8765
4. Show an example of how these key-value pairs might be stored in the hash table slots after hashing.
5. What would happen if two different names produce the same hash slot? How does the hash table
   handle this situation?

## 1. Structure

```go
type entry struct {
	name  string
	phone string
}

type HashTable struct {
	buckets [][]entry
}
```

The table is an array of **slots** (buckets). Each slot holds a list of `entry` values, so
more than one contact can live in the same slot. The table has 7 slots.

Operations:

- `Put(name, phone)` – adds a contact, or updates the phone if the name already exists;
- `Get(name)` – returns the phone and `true`, or `""` and `false` if the name is missing.

## 2. Hash function

The hash function turns the name into a slot number:

1. Add up the character codes (ASCII) of all letters in the name.
2. Take the remainder of the division by the number of slots (`% 7`).

```go
func (h *HashTable) hash(name string) int {
	sum := 0
	for i := 0; i < len(name); i++ {
		sum += int(name[i])
	}
	return sum % len(h.buckets)
}
```

The same name always gives the same slot, so a lookup goes straight to that slot instead of
checking every contact. On average, insert and lookup take O(1) time.

This hash function is simple, which makes it easy to follow by hand. Real hash tables use
better functions (for example FNV or SipHash) that spread keys more evenly – with a sum of
characters, anagrams like "Amy" and "May" always collide.

## 3–4. Inserting the contacts

| Name | Character codes | Sum | Sum % 7 | Slot |
|---|---|---|---|---|
| Alice | 65+108+105+99+101 | 478 | 478 % 7 | **2** |
| Bob | 66+111+98 | 275 | 275 % 7 | **2** |
| Carol | 67+97+114+111+108 | 497 | 497 % 7 | **0** |

The table after the inserts:

```
slot 0: Carol: 555-8765
slot 1: (empty)
slot 2: Alice: 555-1234  ->  Bob: 555-5678
slot 3: (empty)
slot 4: (empty)
slot 5: (empty)
slot 6: (empty)
```

## 5. Collisions

A **collision** happens when two different keys get the same slot. Here Alice and Bob both
land in slot 2.

This table handles collisions with **separate chaining**: each slot holds a list, and a new
contact is appended to the list of its slot. To find Bob, the table computes the slot (2) and
then compares the names in that list until it finds "Bob". Because the full name is stored
with the phone number, the right contact is always found.

Another common method is **open addressing** (for example linear probing): if a slot is
taken, the table tries the next slot (3, 4, …) until it finds a free one.

Too many collisions make lookups slower, up to O(n) in the worst case. To keep them rare,
hash tables use a good hash function and **resize** (usually double the number of slots and
re-insert all keys) when they get too full.

## Go's built-in map

In real Go code you would use the built-in `map`, which is a hash table:

```go
phoneBook := map[string]string{}
phoneBook["Alice"] = "555-1234"
phone, ok := phoneBook["Alice"]
```

## Output

```
Hash function: sum of the character codes of the name % table size
  Alice: 65+108+105+99+101 = 478, 478 % 7 = slot 2
  Bob:   66+111+98 = 275, 275 % 7 = slot 2
  Carol: 67+97+114+111+108 = 497, 497 % 7 = slot 0

Hash table with 7 slots:
  slot 0: Carol: 555-8765
  slot 1: (empty)
  slot 2: Alice: 555-1234  ->  Bob: 555-5678
  ...

Lookups:
  Alice  -> 555-1234
  Bob    -> 555-5678
  Carol  -> 555-8765
  Dave   -> not found
```

## Run

```bash
go run .
```
