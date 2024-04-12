/*
Copyright © 2024 chmikata <chmikata@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh-pkg-cli",
	Short: "gh-pkg-cli is a CLI tool to interact with GitHub Packages",
	Long: `gh-pkg-cli is a CLI tool to interact with GitHub Packages.

You can use this tool to list, search, and delete packages in GitHub Packages.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("org", "o", "", "Organization name")
	rootCmd.PersistentFlags().StringP("token", "t", "", "Token for authentication")
	rootCmd.PersistentFlags().StringP("matcher", "m", ".*", "Name of the container image to match")

	rootCmd.MarkPersistentFlagRequired("org")
	rootCmd.MarkPersistentFlagRequired("token")
}
