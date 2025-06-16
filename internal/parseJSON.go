package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/aquasecurity/table"
)

func ParseJSON(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var data []map[string]any
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	keyCount := make(map[string]int)
	for _, obj := range data {
		keyCount[obj["level"].(string)]++
	}

	t := table.New(os.Stdout)
	t.SetHeaders("Log Level", "Frequency")

	for key, value := range keyCount {
		t.AddRow(key, strconv.Itoa(value))
	}

	t.Render()
}
