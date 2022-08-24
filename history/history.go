package history

import (
	"sync"
)

var once sync.Once

type History []HistoryEntry

type HistoryEntry struct {
	Date              int    `regroup:"date"`
	ExecutionDuration int    `regroup:"duration"`
	TempCommand       string `regroup:"entry"`
	CommandLines      []string
	HashedCommand     string
}

var (
	instance History
)

func GetHistory() *History {
	once.Do(func() {
		instance = make(History, 0)
	})
	return &instance
}

func AppendHistory(entry *HistoryEntry) *History {
	instance = append(*GetHistory(), *entry)
	return GetHistory()
}
