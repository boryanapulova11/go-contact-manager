# go-contact-manager

A simple contact manager in Go, implemented in three different ways.
Each implementation lives in its own branch:

| Branch | Implementation | What it shows |
|---|---|---|
| [`hashTable`](../../tree/hashTable) | Hand-written hash table | How a hash table works inside: hash function, slots, collisions with separate chaining |
| [`map`](../../tree/map) | Go's built-in `map` | The same in a few lines – hashing and collisions are handled by the Go runtime |
| [`struct`](../../tree/struct) | `Contact` struct + `ContactManager` methods | Name, phone and address; encapsulation, validation and error handling |

The `main` branch contains only this overview. Switch to a branch to see its code.

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

## How to run a version

```bash
git clone <repository-url>
cd go-contact-manager
git switch hashTable   # or: map, struct
go run .
```
