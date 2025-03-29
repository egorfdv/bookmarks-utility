package main

import "fmt"

type bookmarkMap = map[string]string

func main() {
	bookmarks := bookmarkMap{}
	fmt.Println("Bookmarking app")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			bookmarks = addBookmark(bookmarks)
		case 3:
			bookmarks = deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Select an option")
	fmt.Println("1. View bookmarks")
	fmt.Println("2. Add a bookmark")
	fmt.Println("3. Delete bookmark")
	fmt.Println("4. Exit")
	fmt.Scan(&variant)
	return variant
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("You have no bookmarks")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	}
}

func addBookmark(bookmarks bookmarkMap) bookmarkMap {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Enter a name: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Enter link: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
	return bookmarks
}

func deleteBookmark(bookmarks map[string]string) bookmarkMap {
	var bookmarkKeyToDelete string
	fmt.Print("Enter a name: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
	return bookmarks
}
