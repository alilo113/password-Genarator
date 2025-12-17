package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	length int
	gen    bool
	add    bool
)

var rootCmd = &cobra.Command{
	Use:   "pwman",
	Short: "CLI password manager",
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Flags().GetInt("length")
        cmd.Flags().GetBool("generate")
        cmd.Flags().GetBool("add")

        if add {
            runAdd("") // pass empty string if no generated password
            return
        }

        var generatedPassword string
        if gen || length != 16 {
            generatedPassword = runGenerate(length)
        }

        if generatedPassword != "" {
            return
        }

        fmt.Println("No action specified. Use -g to generate or -a to add a password.")
    },

}

func init() {
	rootCmd.Flags().BoolVarP(&gen, "generate", "g", false, "Generate a password")
	rootCmd.Flags().IntVarP(&length, "length", "l", 16, "Generate a password with a specified length")
	rootCmd.Flags().BoolVarP(&add, "add", "a", false, "Store a new password on your machine")
}

func Execute() {
	rootCmd.Execute()
}