package commands

import (
	"fmt"
	"strings"
)

func runSearch(query string) {
	query = strings.TrimSpace(strings.ToLower(query))
	if query == "" {
		fmt.Println("Please provide a search term.")
		return
	}

	entries, err := loadStorage()
	if err != nil {
		fmt.Println("Failed to read storage:", err)
		return
	}

	found := false

	fmt.Printf("%-25s %s\n", "SITE", "USERNAME")
	fmt.Println(strings.Repeat("-", 40))

	for _, entry := range entries {
		site := strings.ToLower(entry.Site)
		user := strings.ToLower(entry.Username)

		if strings.Contains(site, query) || strings.Contains(user, query) {
			fmt.Printf("%-25s %s\n", entry.Site, entry.Username)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching entries found.")
	}
}