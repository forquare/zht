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
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"zht/history"
	"zht/utils"
)

var repeated, allRepeated, unique, count, sort bool
var h *history.History

// uniqCmd represents the uniq command
var uniqCmd = &cobra.Command{
	Use:     "uniq [flags] [file ...]",
	Aliases: []string{"unique"},
	Short:   "Report or filter out repeated commands",
	Long: `Parse all provided history files (or ~/.zsh_history by default) and prints out a uniq history.

By default will only match adjacent duplicates, if a command occurs more than once in another part of the history file it will not get filtered out (see --unique to stop this).`,
	Args: func(cmd *cobra.Command, args []string) error {
		return utils.CheckArgsHistoryFileExists(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		utils.ParseAllFiles(args)

		if sort {
			h = utils.SortByDate(*history.GetHistory(), false)
		} else {
			h = history.GetHistory()
		}

		if repeated {
			showRepeated()
		} else if allRepeated {
			showAllRepeated()
		} else if unique {
			showUniq()
		} else if count {
			countOccurrences()
		} else {
			defaultUniq()
		}
	},
}

func init() {
	rootCmd.AddCommand(uniqCmd)

	uniqCmd.Flags().BoolVarP(&repeated, "repeated", "d", false, "Print a single copy of each entry that is repeated - the last occurrence will be printed.")
	uniqCmd.Flags().BoolVarP(&allRepeated, "all-repeated", "D", false, "Print all entries that are repeated.")
	uniqCmd.Flags().BoolVarP(&unique, "unique", "u", false, "Only print entries that are not repeated.")
	uniqCmd.Flags().BoolVarP(&count, "count", "c", false, "Display the number of times each entry occurs - omits date and duration.")
	uniqCmd.Flags().BoolVarP(&sort, "sort", "s", false, "Sort input by date before making output uniq. When used with -c it will sort commands (ascending) by the number of times the occur.")
}

func showRepeated() {
	var entries = make(map[string]history.HistoryEntry)
	var h history.History

	for _, v := range *history.GetHistory() {
		s := utils.FormatCommand(v.CommandLines)
		entries[s] = v
	}

	for _, v := range entries {
		h = append(h, v)
	}

	utils.PrintHistory(utils.SortByDate(h, false))
}

func showAllRepeated() {
	var entries = make(map[string]int)
	var h history.History

	for _, v := range *history.GetHistory() {
		s := utils.FormatCommand(v.CommandLines)
		entries[s]++
	}

	for _, v := range *history.GetHistory() {
		s := utils.FormatCommand(v.CommandLines)
		if entries[s] > 1 {
			h = append(h, v)
		}
	}

	utils.PrintHistory(&h)
}

func showUniq() {
	var entries = make(map[string]int)
	var h history.History

	for _, v := range *history.GetHistory() {
		s := utils.FormatCommand(v.CommandLines)
		entries[s]++
	}

	for _, v := range *history.GetHistory() {
		s := utils.FormatCommand(v.CommandLines)
		if entries[s] == 1 {
			h = append(h, v)
		}
	}

	utils.PrintHistory(&h)
}

func countOccurrences() {
	var entries = make(map[string]int)
	var biggest int
	for _, v := range *history.GetHistory() {
		s := utils.FormatCommand(v.CommandLines)
		entries[s]++
		if entries[s] > biggest {
			biggest = entries[s]
		}
	}

	biggestLength := len(strconv.Itoa(biggest)) + 2

	var s string
	if sort {
		commands := utils.SortMapByValue(&entries)
		for i := 0; i < len(commands); i++ {
			command := commands[i]
			occurrences := entries[command]
			occurrencesLength := len(strconv.Itoa(occurrences))

			s += _formatCount(command, occurrences, biggestLength, occurrencesLength)
		}
	} else {
		for command, occurrences := range entries {
			occurrencesLength := len(strconv.Itoa(occurrences))

			s += _formatCount(command, occurrences, biggestLength, occurrencesLength)
		}
	}
	utils.PrintString(s)
}

func _formatCount(command string, occurrences int, biggestLength int, occurrencesLength int) string {
	var s string
	first := true
	i := 0
	l := len(strings.Split(command, "\n"))
	for _, c := range strings.Split(command, "\n") {
		if i == l-1 {
			break
		}
		i++
		if first {
			s += fmt.Sprintf("%d%*s\n", occurrences, biggestLength+len(c)-occurrencesLength, c)
			first = false
		} else {
			s += fmt.Sprintf("%*s\n", biggestLength+len(c), c)
		}
	}
	return s
}

func defaultUniq() {
	var temp string
	for _, v := range *history.GetHistory() {
		if v.HashedCommand != temp {
			utils.PrintHistoryEntry(&v)
			temp = v.HashedCommand
		}
	}
}
