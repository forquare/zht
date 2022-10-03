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
	"github.com/gobwas/glob"
	"regexp"
	"strings"
	"zht/history"
	"zht/utils"

	"github.com/spf13/cobra"
)

var invert, icase, searchcount bool
var regex, posixregex, globbing, pattern string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:     "search [flags] <pattern> [file ...]",
	Aliases: []string{"find"},
	Short:   "Search history using regular expressions",
	Long: `Search through ZSH history using regular expressions.

By default, no regex is used and the pattern is taken to be exactly as is.
When using regex mode, by default zsh_history will use RE2 (https://github.com/google/re2/wiki/Syntax).
POSIX compatible Extended Regular Expressions (ERE) can also be used.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if _flagProvidesPattern() {
			return utils.CheckArgsHistoryFileExists(args)
		} else {
			return utils.CheckArgsHistoryFileExists(args[1:])
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if _flagProvidesPattern() {
			utils.ParseAllFiles(args)
		} else {
			pattern = args[0]
			utils.ParseAllFiles(args[1:])
		}

		search()
		/*
			search with regex
			search without regex (like grep -F)
			fuzzy search
			add date.range - fuzzy date
			search between dates (filter)
			invert selection (like grep -v)
			case-insensitive search
			count times search matches
		*/
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&regex, "regex", "r", "", "Search using regular expressions")
	searchCmd.Flags().StringVarP(&posixregex, "posix", "R", "", "Search using POSIX ERE regular expressions")
	//searchCmd.Flags().StringVarP(&date, "date", "d", "", "Search between date ranges")
	searchCmd.Flags().StringVarP(&globbing, "glob", "g", "", "Search using globbing")

	searchCmd.Flags().BoolVarP(&invert, "invert-match", "v", false, "Print lines that do not match your pattern")
	searchCmd.Flags().BoolVarP(&count, "count", "c", false, "Print the number of times your pattern matches")
	searchCmd.Flags().BoolVarP(&icase, "ignore-case", "i", false, "Ignore case while searching")
	//searchCmd.Flags().BoolVarP(&newlines, "collapse-newlines", "n", false, "When searching remove newlines in command history so patterns match across them")

}

func search() {
	h = history.GetHistory()

	if regex != "" {
		r := regexp.MustCompile(regex)
		for _, entry := range *h {
			if r.MatchString(strings.Join(entry.CommandLines, "\n")) {
				utils.PrintHistoryEntry(&entry)
			}
		}
	} else if posixregex != "" {
		r := regexp.MustCompilePOSIX(posixregex)
		for _, entry := range *h {
			if r.MatchString(strings.Join(entry.CommandLines, "\n")) {
				utils.PrintHistoryEntry(&entry)
			}
		}
	} else if globbing != "" {
		g := glob.MustCompile(globbing)
		for _, entry := range *h {
			if g.Match(strings.Join(entry.CommandLines, "\n")) {
				utils.PrintHistoryEntry(&entry)
			}
		}
	} else {
		for _, entry := range *h {
			if strings.Contains(strings.Join(entry.CommandLines, "\n"), pattern) {
				utils.PrintHistoryEntry(&entry)
			}
		}
	}
}

func _flagProvidesPattern() bool {
	if regex != "" || posixregex != "" || globbing != "" {
		return true
	}
	return false
}
