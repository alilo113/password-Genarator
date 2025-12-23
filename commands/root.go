package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	length int
	gen    bool
	add    bool
    get  bool
    list bool
    search string
)

var rootCmd = &cobra.Command{
	Use:   "pwman",
	Short: "CLI password manager",
    Run: func(cmd *cobra.Command, args []string) {

        var generatedPassword string

        // 1. Generate only if user explicitly asked
        if gen {
            generatedPassword = generatePassword(length)
            fmt.Println("Generated password:", generatedPassword)
        }

        // generate if -l is specified without -g
        if length != 16 && !gen {
            generatedPassword = generatePassword(length)
            fmt.Println("Generated password:", generatedPassword)
        }

        // 2. Add uses generated password if present
        if add {
            runAdd(generatedPassword)
            return
        }

        // 3. Get passwords
        if get {
            runGet()
            return
        }

        // List passwords
        if list {
            listPasswords()
            return
        }

        // 4. Nothing requested
        if !gen && !add && !get {
            fmt.Println("No action specified. Use -g, -a, or -r.")
        }

        // Search passwords
        if search != "" {
            runSearch(search)
            return
        },
}

func init() {
    rootCmd.Flags().BoolVarP(&gen, "generate", "g", false, "Generate a password")
    rootCmd.Flags().IntVarP(&length, "length", "l", 16, "Generate a password with a specified length") // use 'l' for length
    rootCmd.Flags().BoolVarP(&add, "add", "a", false, "Store a new password on your machine")
    rootCmd.Flags().BoolVarP(&get, "get", "r", false, "Retrieve a stored password from your machine")
    rootCmd.Flags().BoolVarP(&list, "list", "s", false, "List all stored passwords")
    rootCmd.Flags().StringVarP(&search, "search", "f", "", "Search for a password by site name or username")
}

func Execute() {
	rootCmd.Execute()
}