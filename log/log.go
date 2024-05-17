package log

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pattern struct {
	Expression string
	Value      int
}

func LogParse() {
	var parsedPatterns []Pattern

	f, _ := os.Open("./logger.log")
	buf, _ := io.ReadAll(f)
	data := string(buf)
	// fmt.Println(data)

MAINLOOP:
	for _, p := range strings.Split(data, "\n") {
		parts := strings.Split(p, ": ")
		if len(parts) != 2 {
			continue
		}
		bcnt := 0

		for _, ss := range parts[0] {
			if ss == '[' {
				bcnt += 1
			}
			if bcnt > 1 {
				continue MAINLOOP
			}
		}

		value, _ := strconv.Atoi(parts[1])

		parsedPatterns = append(parsedPatterns, Pattern{Expression: parts[0], Value: value})
	}

	sort.Slice(parsedPatterns, func(i, j int) bool {
		return parsedPatterns[i].Value < parsedPatterns[j].Value
	})

	// 結果をファイルに書き出し
	outputFile, err := os.Create("sorted_patterns.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer outputFile.Close()

	for _, p := range parsedPatterns {
		_, err := fmt.Fprintf(outputFile, "%s: %d\n", p.Expression, p.Value)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
	}
}
