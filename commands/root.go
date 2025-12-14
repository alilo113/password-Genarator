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
        // Read flags from the command to make sure values are updated
        cmd.Flags().GetInt("length") // ensures `length` is parsed correctly
        cmd.Flags().GetBool("generate")

		if gen || length != 16 {
			runGenerate(length) // pass length explicitly
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