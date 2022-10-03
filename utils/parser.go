package utils

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"github.com/oriser/regroup"
	"os"
	"strings"
	"zht/history"
)

var mainEntryRE = regroup.MustCompile(`^: (?P<date>\d+):(?P<duration>\d+);(?P<entry>.*)`)

func ParseFile(path string) {
	var file *os.File
	if path == CUSTOM_STDIN {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(path)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	}

	reader := bufio.NewReader(file)

	for {
		entry := &history.HistoryEntry{}

		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if err := mainEntryRE.MatchToTarget(line, entry); err != nil {
			panic(err)
		}

		entry.CommandLines = append(entry.CommandLines, entry.TempCommand)
		entry.TempCommand = ""

		for {
			if strings.HasSuffix(strings.TrimSpace(entry.CommandLines[len(entry.CommandLines)-1]), "\\") {
				command, err := reader.ReadString('\n')
				if err != nil {
					panic(err)
				}
				entry.CommandLines = append(entry.CommandLines, command)
			} else {
				break
			}
		}
		entry.HashedCommand = fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(entry.CommandLines, "\\\n"))))
		history.AppendHistory(entry)
	}
}
