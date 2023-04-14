package cmd

import (
	"github.com/sebps/filepacker/pkg"
	"github.com/spf13/cobra"
)

var source string
var target string
var pck string

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack a directory",
	Long: `
	Pack the target directory into a single .go blob file exposing each file content through a getter function. 

	For example:

	filepacker -s resources -t filestorage -p main

	Flags:
	-s, --source     	source directory to pack
	-t, --target   		generated go file name
	-p, --package		go package name of the generated go file 
`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Pack(source, target, pck)
	},
}

func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.PersistentFlags().StringVarP(&source, "source", "s", ".", "source directory to pack ( default is '.' )")
	packCmd.PersistentFlags().StringVarP(&target, "target", "t", "filepackage", "generated file name - without .go extension ( default is 'filestorage' )")
	packCmd.PersistentFlags().StringVarP(&pck, "package", "p", "main", "go package of the generated file")
}
