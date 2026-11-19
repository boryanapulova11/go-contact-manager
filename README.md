# go-hashtable

A simple contact manager in Go built on the built-in `map`: name → phone number.

## Versions

The commit history shows how the solution was improved:

1. **First commit – hand-written hash table.** A `HashTable` type with 7 slots, a simple hash
   function (sum of character codes % 7) and separate chaining for collisions. It shows how a
   hash table works inside: Alice and Bob collide in slot 2.
2. **Second commit – Go's built-in `map`.** The same contact manager in a few lines. The map
   is a hash table built into the language: hashing, collisions and growing are handled by the
   Go runtime. It also adds update and delete.

To see the first version, open the first commit on GitHub or run `git log` and
`git show <commit>:main.go`.

## Task

Imagine you are building a simple contact manager application. You want to store people's phone
numbers so that you can quickly look them up by their unique names.

1. Design a hash table structure to store the contact information where:
   - the key is the person's name (a unique string);
   - the value is their phone number (a string of digits).
2. Explain how the hash function would work to place each contact into the hash table.
3. Insert the contacts: Alice: 555-1234, Bob: 555-5678, Carol: 555-8765.
4. Show how these key-value pairs might be stored in the hash table.
5. What would happen if two different names produce the same hash slot?

## 1. Structure

```go
phoneBook := make(map[string]string) // key: name, value: phone
```

| Operation | Code | Average time |
|---|---|---|
| Insert / update | `phoneBook["Alice"] = "555-1234"` | O(1) |
| Lookup | `phone, ok := phoneBook["Alice"]` | O(1) |
| Delete | `delete(phoneBook, "Alice")` | O(1) |
| Count | `len(phoneBook)` | O(1) |

Keys are unique: writing a second value for the same name replaces the old one.

## 2. Hash function

Go hashes the key for us. The runtime uses a fast hash function for strings, mixed with a
**random seed** chosen when the map is created. The hash decides where in the table the
pair is stored. Because of the random seed, the layout – and the order of `range` over a
map – changes between runs, so the program sorts the names before printing them.

## 3–4. Storing the contacts

```go
phoneBook["Alice"] = "555-1234"
phoneBook["Bob"] = "555-5678"
phoneBook["Carol"] = "555-8765"
```

The exact slots are hidden inside the runtime and differ between runs. Conceptually the table
looks like the one from the first version:

```
slot ?: Alice: 555-1234
slot ?: Bob:   555-5678
slot ?: Carol: 555-8765
```

## 5. Collisions

The `map` handles collisions on its own. Since Go 1.24, `map` is implemented as a
**Swiss Table**: slots are grouped by 8, and if the preferred group is full, the map
probes other groups (open addressing). The map stores the full key next to the value, so a
lookup compares keys and always finds the right contact. When the table gets too full, the
map **grows** and moves the entries into a larger table, which keeps lookups fast.

## Hand-written hash table vs map

| | First version (hand-written) | Second version (`map`) |
|---|---|---|
| Hash function | sum of character codes % 7 | runtime hash with random seed |
| Collisions | separate chaining (list per slot) | handled internally (open addressing) |
| Growing | fixed size (7 slots) | grows automatically |
| Code | ~70 lines | one line |
| Use | to learn how hash tables work | real programs |

## Things to remember

- `var m map[string]string` is a **nil map**: reading works, writing causes a panic. Create it
  with `make` or a map literal.
- Use the **comma ok** form `v, ok := m[key]` to tell "missing key" from "empty value".
- The order of `range` over a map is random. Sort the keys if you need a stable order.
- `delete` on a missing key does nothing and returns no error.
- A map is not safe for concurrent writes from several goroutines – use `sync.Mutex` or
  `sync.Map`.

## Output

```
All contacts:
  Alice  -> 555-1234
  Bob    -> 555-5678
  Carol  -> 555-8765
  total contacts: 3

Lookups:
  Alice  -> 555-1234
  Bob    -> 555-5678
  Carol  -> 555-8765
  Dave   -> not found

Update Bob's number:
  Bob    -> 555-0000

Delete Carol:
  Carol  -> not found

All contacts after the changes:
  Alice  -> 555-1234
  Bob    -> 555-0000
  total contacts: 2
```

## Run

```bash
go run .
```
