/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	mcobra "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"
)

// manCmd represents the man command
var manCmd = &cobra.Command{
	Use:   "man",
	Short: "Generate man pages",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		manPage, err := mcobra.NewManPage(1, rootCmd)
		if err != nil {
			panic(err)
		}

		manPage = manPage.WithSection("Copyright", "(C) 2022 Christian Muehlhaeuser.\n"+
			"Released under MIT license.")

		f, err2 := os.Create("zht.1")
		if err2 != nil {
			fmt.Println(err)
		}
		defer f.Close()
		f.WriteString(manPage.Build(roff.NewDocument()))
	},
}

func init() {
	genCmd.AddCommand(manCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// manCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// manCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
