package commands

import "github.com/spf13/cobra"

var (
    length int
    gen    bool
    add bool
)

var rootCmd = &cobra.Command{
    Use:   "pwman",
    Short: "CLI password manager",
    Run: func(cmd *cobra.Command, args []string) {
        // Read flags from the command to make sure values are updated
        cmd.Flags().GetInt("length") // ensures `length` is parsed correctly
        cmd.Flags().GetBool("generate")
        cmd.Flags().GetBool("add")

        if add {
            runAdd() // <-- call a function to handle adding a password
            return
        }

		if gen || length != 16 {
			runGenerate(length) // pass length explicitly
			return
		}
        println("No action specified. Use -g to generate a password.")
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