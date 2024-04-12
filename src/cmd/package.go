/*
Copyright © 2024 chmikata <chmikata@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/chmikata/gh-pkg-cli/internal/application"
	"github.com/spf13/cobra"
)

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Display package",
	Long:  "Display package information",
	RunE: func(cmd *cobra.Command, args []string) error {
		org, _ := rootCmd.PersistentFlags().GetString("org")
		token, _ := rootCmd.PersistentFlags().GetString("token")
		matcher, _ := rootCmd.PersistentFlags().GetString("matcher")
		registry := application.NewRegistry(org, token)
		packages, err := registry.GetPackages(matcher)
		if err != nil {
			fmt.Println(err)
			return err
		}
		v, err := json.Marshal(packages)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(string(v))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(packageCmd)
}
