package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/aquasecurity/table"
)

func ParseLog(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	logLevelCount := make(map[string]int)
	scanner := bufio.NewScanner(file)

	regex := regexp.MustCompile(`\[(\w+)\]`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindStringSubmatch(line)
		if len(matches) > 1 {
			level := matches[1]
			logLevelCount[level]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	t := table.New(os.Stdout)
	t.SetHeaders("Log Level", "Frequency")

	logLevels := make([]string, 0, len(logLevelCount))
	for level := range logLevelCount {
		logLevels = append(logLevels, level)
	}
	sort.Strings(logLevels)

	for _, level := range logLevels {
		t.AddRow(level, strconv.Itoa(logLevelCount[level]))
	}

	t.Render()
}
