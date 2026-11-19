package main

import (
	"fmt"
	"sort"
)

func printContacts(phoneBook map[string]string) {
	names := make([]string, 0, len(phoneBook))
	for name := range phoneBook {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("  %-6s -> %s\n", name, phoneBook[name])
	}

	fmt.Println("  total contacts:", len(phoneBook))
}

func lookup(phoneBook map[string]string, name string) {
	if phone, ok := phoneBook[name]; ok {
		fmt.Printf("  %-6s -> %s\n", name, phone)
	} else {
		fmt.Printf("  %-6s -> not found\n", name)
	}
}

func main() {
	phoneBook := make(map[string]string)

	phoneBook["Alice"] = "555-1234"
	phoneBook["Bob"] = "555-5678"
	phoneBook["Carol"] = "555-8765"

	fmt.Println("All contacts:")
	printContacts(phoneBook)

	fmt.Println()
	fmt.Println("Lookups:")
	for _, name := range []string{"Alice", "Bob", "Carol", "Dave"} {
		lookup(phoneBook, name)
	}

	fmt.Println()
	fmt.Println("Update Bob's number:")
	phoneBook["Bob"] = "555-0000"
	lookup(phoneBook, "Bob")

	fmt.Println()
	fmt.Println("Delete Carol:")
	delete(phoneBook, "Carol")
	lookup(phoneBook, "Carol")

	fmt.Println()
	fmt.Println("All contacts after the changes:")
	printContacts(phoneBook)

	fmt.Println()
	fmt.Println("Map literal:")
	team := map[string]string{
		"Alice": "555-1234",
		"Bob":   "555-5678",
		"Carol": "555-8765",
	}
	printContacts(team)
}
