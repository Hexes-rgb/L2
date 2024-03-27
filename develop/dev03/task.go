package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// объявляем опции
var (
	column  = flag.Int("k", 0, "column to sort by")
	numeric = flag.Bool("n", false, "sort by numeric value")
	reverse = flag.Bool("r", false, "reverse sort order")
	unique  = flag.Bool("u", false, "output only the first of an equal run")
)

// byColumn реализует sort.Interface для сортировки по указанному столбцу
type byColumn struct {
	lines     []string
	column    int
	numeric   bool
	reverse   bool
	separator string
}

func (s byColumn) Len() int {
	return len(s.lines)
}

func (s byColumn) Swap(i, j int) {
	s.lines[i], s.lines[j] = s.lines[j], s.lines[i]
}

func (s byColumn) Less(i, j int) bool {
	icolumns := strings.Split(s.lines[i], s.separator)
	jcolumns := strings.Split(s.lines[j], s.separator)

	// убедимся, что сравниваем в определенных пределах
	if s.column >= len(icolumns) || s.column >= len(jcolumns) {
		return false
	}

	if s.numeric {
		inum, err1 := strconv.Atoi(icolumns[s.column])
		jnum, err2 := strconv.Atoi(jcolumns[s.column])
		if err1 == nil && err2 == nil {
			if s.reverse {
				return inum > jnum
			}
			return inum < jnum
		}
	}

	if s.reverse {
		return icolumns[s.column] > jcolumns[s.column]
	}
	return icolumns[s.column] < jcolumns[s.column]
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: sort [options] filename")
		os.Exit(1)
	}

	filename := flag.Arg(0)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	sort.Sort(byColumn{lines, *column, *numeric, *reverse, " "})

	lastLine := ""
	for _, line := range lines {
		if *unique && line == lastLine {
			continue
		}
		fmt.Println(line)
		lastLine = line
	}
}
