/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:                   "gen",
	DisableFlagsInUseLine: true,
	SilenceUsage:          true,
	Args:                  cobra.NoArgs,
	Short:                 "A collection of several useful generators",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
