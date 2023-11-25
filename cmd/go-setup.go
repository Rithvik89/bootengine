package cmd

import (
	go_template "bootengine/pkg/template/go"

	"github.com/spf13/cobra"
)

var source string
var dest string

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "golang project setup",
	Long:  " TODO ",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		go_template.CloneProject(args[0], dest, source)
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
	goCmd.Flags().StringVarP(&source, "source", "s", "", "This is the source path from which bootengine to take reference. ")
	goCmd.Flags().StringVarP(&dest, "destination", "d", "", "Destination where the desired project setup is to be made. ")
}
