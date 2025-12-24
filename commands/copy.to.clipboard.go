package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func runCopy() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Site: ")
	site, _ := reader.ReadString('\n')
	site = strings.TrimSpace(site)

	if site == "" {
		fmt.Println("Site name cannot be empty.")
		return
	}

	fmt.Print("Master password: ")
	master, _ := reader.ReadString('\n')
	master = strings.TrimSpace(master)

	entries, err := loadStorage()
	if err != nil {
		fmt.Println("Failed to read storage:", err)
		return
	}

	for _, entry := range entries {
		if strings.EqualFold(entry.Site, site) {
			password, err := decryptPassword(master, entry.Password)
			if err != nil {
				fmt.Println("Wrong master password.")
				return
			}

			if err := clipboard.WriteAll(password); err != nil {
				fmt.Println("Failed to copy to clipboard:", err)
				return
			}

			fmt.Println("Password copied to clipboard (clears on overwrite).")
			return
		}
	}

	fmt.Println("No entry found for site:", site)
}