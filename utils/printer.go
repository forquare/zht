package utils

import (
	"fmt"
	"strings"
	"time"
	"zht/history"
)

var HumanTime bool
var HumanTimeFormat string

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
	var s string
	if HumanTime || HumanTimeFormat != "" {
		if HumanTimeFormat == "" {
			HumanTimeFormat = time.RFC1123
		}
		date := time.Unix(int64(entry.Date), 0)
		s = fmt.Sprintf(": %s:%d;", date.Format(HumanTimeFormat), entry.ExecutionDuration)
	} else {
		s = fmt.Sprintf(": %d:%d;", entry.Date, entry.ExecutionDuration)
	}
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
