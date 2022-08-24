package utils

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"zht/history"
)

const CUSTOM_STDIN = "[[STDIN}}"

func CheckArgsHistoryFileExists(args []string) error {
	args = GetDefaultZshHistoryFileIfNeeded(args)

	if args[0] == CUSTOM_STDIN {
		return nil
	}

	for i := 0; i < len(args); i++ {
		historyFile := args[i]

		if historyFileInfo, err := os.Stat(historyFile); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("history file '%s' does not exist", historyFile)
		} else if historyFileInfo.IsDir() {
			return fmt.Errorf("history file '%s' is a directory, not a file", historyFile)
		}
	}
	return nil
}

func GetDefaultZshHistoryFileIfNeeded(args []string) []string {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		if len(args) == 0 {
			userHomeDir, err := os.UserHomeDir()
			if err != nil {
				panic(err)
			}
			defaultHistoryFile := userHomeDir + "/.zsh_history"
			args = append(args, defaultHistoryFile)
		}
	} else {
		args = append(args, CUSTOM_STDIN)
	}
	return args
}

func ParseAllFiles(args []string) {
	args = GetDefaultZshHistoryFileIfNeeded(args)
	for i := 0; i < len(args); i++ {
		ParseFile(args[i])
	}
}

func SortByDate(h history.History, reverse bool) *history.History {
	sort.Slice(h, func(i, j int) bool {
		if reverse {
			return h[i].Date > h[j].Date
		} else {
			return h[i].Date < h[j].Date
		}
	})
	return &h
}

type Pair struct {
	Key   string
	Value int
}

/*
From: https://medium.com/@kdnotes/how-to-sort-golang-maps-by-value-and-key-eedc1199d944#af09
*/
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func SortMapByValue(m *map[string]int) []string {
	var s []string
	p := make(PairList, len(*m))

	i := 0
	for k, v := range *m {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	//p is sorted

	for _, k := range p {
		s = append(s, fmt.Sprintf("%v", k.Key))
	}

	return s
}
