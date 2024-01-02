//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LenderAudit struct {
	checkout time.Time
	checkin  time.Time
}
type Member struct {
	name  Name
	books map[Title]LenderAudit
}

type BookEntry struct {
	totel  int
	lended int
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkin.IsZero() {
			returnTime = "not returned yet"
		} else {
			returnTime = audit.checkin.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkout.String(), ":", returnTime)
	}
}

func printLibraryAudit(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	for title, entry := range library.books {
		fmt.Println(title, ":", entry.totel, "/", entry.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, member *Member, title Title) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not in library collection!")
		return false
	}
	if book.lended == book.totel {
		fmt.Println("Book not available")
		return false
	}
	book.lended++
	library.books[title] = book // Why do we need to do this? what happens if we don't?

	member.books[title] = LenderAudit{checkout: time.Now()}
	return true
}

func checkinBooks(library *Library, member *Member, title Title) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("Book not in library collection!")
		return false
	}
	book.lended--
	library.books[title] = book

	audit, found := member.books[title]
	if !found {
		fmt.Println("Book not checked out by the user!")
		return false
	}

	audit.checkin = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		members: make(map[Name]Member),
		books:   make(map[Title]BookEntry),
	}
	library.books["The Hobbit"] = BookEntry{totel: 4, lended: 0}
	library.books["The Fellowship of the Ring"] = BookEntry{totel: 1, lended: 0}
	library.books["The Two Towers"] = BookEntry{totel: 2, lended: 0}

	library.members["Casey"] = Member{"Casey", make(map[Title]LenderAudit)}
	library.members["Steve"] = Member{"Steve", make(map[Title]LenderAudit)}
	library.members["Bill"] = Member{"Bill", make(map[Title]LenderAudit)}

	fmt.Println("\nInit: ")
	printLibraryBooks(&library)
	printLibraryAudit(&library)

	member := library.members["Casey"]
	checkout := checkoutBook(&library, &member, "The Hobbit")
	fmt.Println("\nCheckout on book: ")
	if checkout {
		printLibraryBooks(&library)
		printLibraryAudit(&library)
	}
	returned := checkinBooks(&library, &member, "The Hobbit")
	fmt.Println("\nChecked in book: ")
	if returned {
		printLibraryBooks(&library)
		printLibraryAudit(&library)
	}

}
