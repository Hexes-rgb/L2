package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Флаги командной строки
var (
	afterFlag      = flag.Int("A", 0, "print N lines after match")
	beforeFlag     = flag.Int("B", 0, "print N lines before match")
	contextFlag    = flag.Int("C", 0, "print N lines around match")
	countFlag      = flag.Bool("c", false, "print the count of lines")
	ignoreCaseFlag = flag.Bool("i", false, "ignore case")
	invertFlag     = flag.Bool("v", false, "select non-matching lines")
	fixedFlag      = flag.Bool("F", false, "interpret pattern as a fixed string")
	lineNumFlag    = flag.Bool("n", false, "print line number with output lines")
)

// Функция для поиска совпадений по заданным параметрам
func grep(scanner *bufio.Scanner, pattern string, after, before int,
	count, ignoreCase, invert, fixed, lineNum bool, writer *bufio.Writer) {

	var matchCount int
	var beforeLines []string
	var afterLines []int
	currentLineNum := 1

	// Компиляция регулярного выражения с учётом флагов
	var re *regexp.Regexp
	var err error
	if fixed {
		pattern = regexp.QuoteMeta(pattern)
	}
	if ignoreCase {
		re, err = regexp.Compile("(?i)" + pattern)
	} else {
		re, err = regexp.Compile(pattern)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Чтение строк из входного потока
	for scanner.Scan() {
		line := scanner.Text()

		// Проверка совпадения с учетом флагов
		var match bool
		if fixed {
			if ignoreCase {
				match = strings.EqualFold(line, pattern)
			} else {
				match = line == pattern
			}
		} else {
			match = re.MatchString(line)
		}
		if invert {
			match = !match
		}

		if match {
			if count {
				matchCount++
			} else {
				if before > 0 && len(beforeLines) > 0 {
					for _, l := range beforeLines {
						if lineNum {
							writer.WriteString(fmt.Sprintf("%d-", currentLineNum-len(beforeLines)))
						}
						writer.WriteString(fmt.Sprintf("%s\n", l))
					}
					beforeLines = nil
				}

				if lineNum {
					writer.WriteString(fmt.Sprintf("%d:", currentLineNum))
				}
				writer.WriteString(fmt.Sprintf("%s\n", line))

				if after > 0 {
					afterLines = make([]int, after)
				}
			}
		} else {
			if before > 0 {
				if len(beforeLines) == before {
					beforeLines = beforeLines[1:]
				}
				beforeLines = append(beforeLines, line)
			}
		}

		if len(afterLines) > 0 {
			afterLines = afterLines[1:]
			if len(afterLines) == 0 && !match {
				continue
			}
			if lineNum && !match {
				writer.WriteString(fmt.Sprintf("%d-", currentLineNum))
			}
			writer.WriteString(fmt.Sprintf("%s\n", line))
		}

		currentLineNum++
	}

	if count {
		writer.WriteString(fmt.Sprintf("%d\n", matchCount))
	}

	if err := scanner.Err(); err != nil {
		writer.Flush()
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	writer.Flush()
}

func main() {
	// Парсинг флагов командной строки
	flag.Parse()

	pattern := flag.Arg(0)
	files := flag.Args()[1:]

	if *contextFlag > 0 {
		*afterFlag = *contextFlag
		*beforeFlag = *contextFlag
	}

	writer := bufio.NewWriter(os.Stdout)

	if len(files) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		grep(scanner, pattern, *afterFlag, *beforeFlag, *countFlag,
			*ignoreCaseFlag, *invertFlag, *fixedFlag, *lineNumFlag, writer)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
				continue
			}
			scanner := bufio.NewScanner(f)
			grep(scanner, pattern, *afterFlag, *beforeFlag, *countFlag,
				*ignoreCaseFlag, *invertFlag, *fixedFlag, *lineNumFlag, writer)
			f.Close()
		}
	}

	writer.Flush()
}
