package commands

import "github.com/spf13/cobra"

var (
    length int
    gen    bool
)

var rootCmd = &cobra.Command{
    Use:   "pwman",
    Short: "CLI password manager",
    Run: func(cmd *cobra.Command, args []string) {
        if gen {
            runGenerate() // call the generate logic from generate.go
            return
        }
        println("No action specified. Use -g to generate a password.")
    },
}

func init() {
    rootCmd.Flags().BoolVarP(&gen, "generate", "g", false, "Generate a password")
    rootCmd.Flags().IntVarP(&length, "length", "l", 16, "Length of the password")
}

func Execute() {
    rootCmd.Execute()
}