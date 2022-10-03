package utils

import (
	"fmt"
	"strings"
	"zht/history"
)

func FormatCommand(commandLines []string) string {
	var s string
	for _, command := range commandLines {
		if strings.HasSuffix(command, "\n") {
			s += fmt.Sprintf("%s", command)
		} else {
			s += fmt.Sprintf("%s\n", command)
		}
	}
	return s
}

func FormatHistoryEntryOutput(entry *history.HistoryEntry) string {
	s := fmt.Sprintf(": %d:%d;", entry.Date, entry.ExecutionDuration)
	s += FormatCommand(entry.CommandLines)
	return s
}

func PrintHistory(h *history.History) {
	for _, entry := range *h {
		PrintString(FormatHistoryEntryOutput(&entry))
	}
}

func PrintHistoryEntry(h *history.HistoryEntry) {
	var list history.History
	list = append(list, *h)
	PrintHistory(&list)
}

func PrintString(s string) {
	fmt.Printf(s)
}
