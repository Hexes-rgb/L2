package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fieldsFlag := flag.String("f", "", "fields - select only these fields")
	delimiterFlag := flag.String("d", "\t", "delimiter - use a different delimiter")
	separatedFlag := flag.Bool("s", false, "separated - only lines with the delimiter")

	flag.Parse()

	// Проверяем, что флаг fields задан
	if *fieldsFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: flag -f is required")
		os.Exit(1)
	}

	// Разбиваем fieldsFlag на отдельные номера полей
	fields := strings.Split(*fieldsFlag, ",")
	fieldIndexes := make([]int, 0, len(fields))
	for _, field := range fields {
		var fieldIndex int
		_, err := fmt.Sscanf(field, "%d", &fieldIndex)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid field number: %s\n", field)
			os.Exit(1)
		}
		fieldIndexes = append(fieldIndexes, fieldIndex-1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, *delimiterFlag)

		if *separatedFlag && len(columns) <= 1 {
			continue
		}

		for i, fieldIndex := range fieldIndexes {
			if fieldIndex < 0 || fieldIndex >= len(columns) {
				continue
			}
			if i > 0 {
				fmt.Print(*delimiterFlag)
			}
			fmt.Print(columns[fieldIndex])
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input:", err)
	}
}
