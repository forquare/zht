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

var firstLines int

// headCmd represents the head command
var headCmd = &cobra.Command{
	Use:   "head",
	Short: "Print lines from the top of history",
	Long:  `By default, print the first 10 lines of a history file`,
	Args: func(cmd *cobra.Command, args []string) error {
		return utils.CheckArgsHistoryFileExists(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		utils.ParseAllFiles(args)
		head()
	},
}

func init() {
	rootCmd.AddCommand(headCmd)

	headCmd.Flags().IntVarP(&firstLines, "lines", "n", 10, "The number of lines from the top to print")
}

func head() {
	h := *history.GetHistory()

	if len(h) < firstLines {
		firstLines = len(h)
	}
	h = h[:firstLines]
	utils.PrintHistory(&h)
}
