/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/itsopenmiso/websitegen/docutils"

	"github.com/spf13/cobra"
)

// openmisoCmd represents the openmiso command
var openmisoCmd = &cobra.Command{
	Use:   "openmiso",
	Short: "Build openmiso documentation",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		docutils.PrecompileMarkdowns()
		// Here you have precompiled markdowns
		docutils.GrabJSONs()
		fmt.Println("websitegen finished Rendering markdowns.")
		fmt.Scanln()
		docutils.RemoveMarkdowns()
	},
}

func init() {
	rootCmd.AddCommand(openmisoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openmisoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openmisoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
