package commands

import (
    "fmt"
    "os"
    "encoding/json"
    "path/filepath"
    "sort"
    "strings"
)

func listPasswords() error {
	// Get user home directory safely
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	storagePath := filepath.Join(home, ".pwman", "store.json")

	// If storage file does not exist
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		fmt.Println("No passwords stored yet.")
		return nil
	}

	// Read storage file
	data, err := os.ReadFile(storagePath)
	if err != nil {
		return err
	}

	// Handle empty file
	if len(data) == 0 {
		fmt.Println("No passwords stored yet.")
		return nil
	}

	// Parse JSON
	var entries []Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return err
	}

	if len(entries) == 0 {
		fmt.Println("No passwords stored yet.")
		return nil
	}

	// Sort entries by site name
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Site < entries[j].Site
	})

	// Print output
	fmt.Printf("%-20s %s\n", "NAME", "USERNAME")
	fmt.Println(strings.Repeat("-", 30))

	for _, entry := range entries {
		fmt.Printf("%-20s %s\n", entry.Site, entry.Username)
	}

	return nil
}
