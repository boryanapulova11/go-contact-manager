package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Contact struct {
	Name    string
	Phone   string
	Address string
}

func (c Contact) String() string {
	return fmt.Sprintf("%-6s -> %s, %s", c.Name, c.Phone, c.Address)
}

type ContactManager struct {
	contacts map[string]Contact
}

var (
	ErrExists       = errors.New("contact already exists")
	ErrNotFound     = errors.New("contact not found")
	ErrInvalidPhone = errors.New("phone must contain only digits and '-'")
	ErrNoAddress    = errors.New("address must not be empty")
)

func NewContactManager() *ContactManager {
	return &ContactManager{contacts: make(map[string]Contact)}
}

func validPhone(phone string) bool {
	if phone == "" {
		return false
	}
	for _, r := range phone {
		if (r < '0' || r > '9') && r != '-' {
			return false
		}
	}
	return true
}

func (cm *ContactManager) Add(name, phone, address string) error {
	if !validPhone(phone) {
		return fmt.Errorf("add %q: %w", name, ErrInvalidPhone)
	}
	if strings.TrimSpace(address) == "" {
		return fmt.Errorf("add %q: %w", name, ErrNoAddress)
	}
	if _, ok := cm.contacts[name]; ok {
		return fmt.Errorf("add %q: %w", name, ErrExists)
	}
	cm.contacts[name] = Contact{Name: name, Phone: phone, Address: address}
	return nil
}

func (cm *ContactManager) Get(name string) (Contact, bool) {
	c, ok := cm.contacts[name]
	return c, ok
}

func (cm *ContactManager) UpdatePhone(name, phone string) error {
	if !validPhone(phone) {
		return fmt.Errorf("update %q: %w", name, ErrInvalidPhone)
	}
	c, ok := cm.contacts[name]
	if !ok {
		return fmt.Errorf("update %q: %w", name, ErrNotFound)
	}

	c.Phone = phone
	cm.contacts[name] = c
	return nil
}

func (cm *ContactManager) UpdateAddress(name, address string) error {
	if strings.TrimSpace(address) == "" {
		return fmt.Errorf("update %q: %w", name, ErrNoAddress)
	}
	c, ok := cm.contacts[name]
	if !ok {
		return fmt.Errorf("update %q: %w", name, ErrNotFound)
	}
	c.Address = address
	cm.contacts[name] = c
	return nil
}

func (cm *ContactManager) Delete(name string) error {
	if _, ok := cm.contacts[name]; !ok {
		return fmt.Errorf("delete %q: %w", name, ErrNotFound)
	}
	delete(cm.contacts, name)
	return nil
}

func (cm *ContactManager) List() []Contact {
	list := make([]Contact, 0, len(cm.contacts))
	for _, c := range cm.contacts {
		list = append(list, c)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	return list
}

func (cm *ContactManager) Count() int {
	return len(cm.contacts)
}

func printAll(cm *ContactManager) {
	for _, c := range cm.List() {
		fmt.Println(" ", c)
	}
	fmt.Println("  total contacts:", cm.Count())
}

func check(err error) {
	if err != nil {
		fmt.Println("  error:", err)
	}
}

func main() {
	cm := NewContactManager()

	check(cm.Add("Alice", "555-1234", "12 Oak Street, Sofia"))
	check(cm.Add("Bob", "555-5678", "7 Lake Road, Plovdiv"))
	check(cm.Add("Carol", "555-8765", "3 Hill Avenue, Varna"))

	fmt.Println("All contacts:")
	printAll(cm)

	fmt.Println()
	fmt.Println("Lookups:")
	for _, name := range []string{"Alice", "Bob", "Carol", "Dave"} {
		if c, ok := cm.Get(name); ok {
			fmt.Println(" ", c)
		} else {
			fmt.Printf("  %-6s -> not found\n", name)
		}
	}

	fmt.Println()
	fmt.Println("Invalid operations:")
	check(cm.Add("Alice", "555-9999", "1 Main Street")) // Alice already exists
	check(cm.Add("Dave", "abc-1234", "5 River Street")) // invalid phone number
	check(cm.Add("Dave", "555-4321", "   "))            // empty address
	check(cm.UpdatePhone("Eve", "555-1111"))            // Eve does not exist
	check(cm.Delete("Frank"))                           // Frank does not exist

	// errors.Is checks the kind of error, even when it is wrapped.
	if err := cm.Add("Bob", "555-0000", "7 Lake Road, Plovdiv"); errors.Is(err, ErrExists) {
		fmt.Println("  Bob is already in the list - use UpdatePhone instead")
	}

	fmt.Println()
	fmt.Println("Update Bob's phone, Alice's address, delete Carol:")
	check(cm.UpdatePhone("Bob", "555-0000"))
	check(cm.UpdateAddress("Alice", "20 Rose Street, Sofia"))
	check(cm.Delete("Carol"))
	printAll(cm)
}
