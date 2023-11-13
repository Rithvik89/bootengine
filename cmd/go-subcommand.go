package cmd

import (
	template "bootengine/pkg/template"

	"github.com/spf13/cobra"
)

var dest string
var source string
var name string

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "golang project",
	Long:  "TODO",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		template.CloneProject(args[0], dest, source)
	},
}

func init() {
	goCmd.AddCommand(projectCmd)
	goCmd.Flags().StringVarP(&source, "source", "s", "", "This is the source path from which bootengine to take reference. ")
	goCmd.Flags().StringVarP(&dest, "destination", "d", "", "Destination where the desired project setup is to be made. ")
}
