package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filepacker",
	Short: "A cli tool to pack a directory into a single .go file",
	Long:  `File packer allows to generate a single .go file wrapping the full content of the directory and providing their content through a getter.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
