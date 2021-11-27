package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "awscreds",
	Short: "aws creds tool",
	Long:  "aws creds tool by using mfa device",
}

// Execute is function for execute rootCmd
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
