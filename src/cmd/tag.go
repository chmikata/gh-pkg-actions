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

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Display container image tags",
	Long:  "Display container image tags matching container name and depth",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		pattern, _ := cmd.Flags().GetString("pattern")
		if pattern != "sem" && pattern != "sha" {
			fmt.Println("Pattern must be either 'sem' or 'sha'")
			return fmt.Errorf("invalid pattern")
		}
		depth, _ := cmd.Flags().GetInt("depth")
		if depth < 0 {
			fmt.Println("Depth must be a positive integer")
			return fmt.Errorf("invalid depth")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		org, _ := rootCmd.PersistentFlags().GetString("org")
		token, _ := rootCmd.PersistentFlags().GetString("token")
		matcher, _ := rootCmd.PersistentFlags().GetString("matcher")
		pattern, _ := cmd.Flags().GetString("pattern")
		depth, _ := cmd.Flags().GetInt("depth")
		registry := application.NewRegistry(org, token)
		tags, err := registry.GetTags(matcher, pattern, depth)
		if err != nil {
			fmt.Println(err)
			return err
		}
		v, err := json.Marshal(tags)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(string(v))
		return nil
	},
}

func init() {
	tagCmd.Flags().StringP("pattern", "p", "", "Pattern to <sem> or <sha> match image to tags")
	tagCmd.Flags().IntP("depth", "d", 0, "Depth of tags to display")

	tagCmd.MarkFlagRequired("pattern")

	rootCmd.AddCommand(tagCmd)
}
