/*
Copyright © 2022 Ben Lavery-Griffiths <ben@lavery-griffiths.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"zht/history"
	"zht/utils"

	"github.com/spf13/cobra"
)

var reverseFlag bool

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort [flags] [file ...]",
	Short: "Sort one or more history files",
	Long: `Parse all provided history files (or ~/.zsh_history by default) and prints them out sorted.

By default history will be sorted by timestamp.`,
	Args: func(cmd *cobra.Command, args []string) error {
		return utils.CheckArgsHistoryFileExists(args)
	},

	Run: func(cmd *cobra.Command, args []string) {
		utils.ParseAllFiles(args)
		sortHistory()
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sortCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sortCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sortCmd.Flags().BoolVarP(&reverseFlag, "reverse", "r", false, "Reverse sort")
}

func sortHistory() {
	utils.PrintHistory(utils.SortByDate(*history.GetHistory(), reverseFlag))
}
