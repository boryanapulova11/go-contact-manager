# go-contact-manager – branch `struct`

A contact manager in Go built with structs and methods. Each contact has a **name**, a
**phone number** and an **address**. The contacts are stored in a `map` that is hidden inside
a `ContactManager` type.

Other implementations of the same task are in the `hashTable` and `map` branches.

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

## Solution

```go
type Contact struct {
	Name    string
	Phone   string
	Address string
}

type ContactManager struct {
	contacts map[string]Contact // hidden: only the methods can change it
}
```

| Method | What it does |
|---|---|
| `NewContactManager()` | creates an empty manager (with `make`, so the map is not nil) |
| `Add(name, phone, address)` | adds a new contact; error if it exists, the phone is invalid or the address is empty |
| `Get(name)` | returns the contact and `true`, or `false` if it is missing |
| `UpdatePhone(name, phone)` | changes the phone of an existing contact |
| `UpdateAddress(name, address)` | changes the address of an existing contact |
| `Delete(name)` | removes a contact; error if it is missing |
| `List()` | returns all contacts sorted by name |
| `Count()` | returns the number of contacts |

What this version adds:

- **Encapsulation** – the map field starts with a lower-case letter, so code outside the
  package cannot change it directly. Every change goes through a method.
- **Validation and errors** – methods return `error`. Errors are defined once
  (`ErrExists`, `ErrNotFound`, `ErrInvalidPhone`, `ErrNoAddress`), wrapped with
  `fmt.Errorf("...: %w", err)` and checked with `errors.Is`.
- **`fmt.Stringer`** – `Contact` has a `String()` method, so `fmt.Println(contact)` prints
  it in a readable way.
- **A struct value in a map is a copy** – to change a field, read the struct, change it and
  write it back: `c := m[name]; c.Phone = phone; m[name] = c`.

### Output

```
All contacts:
  Alice  -> 555-1234, 12 Oak Street, Sofia
  Bob    -> 555-5678, 7 Lake Road, Plovdiv
  Carol  -> 555-8765, 3 Hill Avenue, Varna
  total contacts: 3

Lookups:
  Alice  -> 555-1234, 12 Oak Street, Sofia
  Bob    -> 555-5678, 7 Lake Road, Plovdiv
  Carol  -> 555-8765, 3 Hill Avenue, Varna
  Dave   -> not found

Invalid operations:
  error: add "Alice": contact already exists
  error: add "Dave": phone must contain only digits and '-'
  error: add "Dave": address must not be empty
  error: update "Eve": contact not found
  error: delete "Frank": contact not found
  Bob is already in the list - use UpdatePhone instead

Update Bob's phone, Alice's address, delete Carol:
  Alice  -> 555-1234, 20 Rose Street, Sofia
  Bob    -> 555-0000, 7 Lake Road, Plovdiv
  total contacts: 2
```

## Run

```bash
go run .
```
