package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "pwman",
	Short: "pwman is a password manager for sane people",
	Long: `pwman is a password manager for sane people.
It allows you to securely store and manage your passwords with ease.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}