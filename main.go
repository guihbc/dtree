package main

import (
	"fmt"
	"strings"
)

const MULTIPLE_ENTRIES = "├──"
const LAST_ENTRY = "└──"
const BAR = "│"
const IDENTATION = 3

func main() {
	line := build_line(0, "main/", false)
	fmt.Println(line)
}

func build_line(depth int, entry_name string, last_entry bool) string {
	base_line := BAR + strings.Repeat(" ", IDENTATION)
	line := strings.Repeat(base_line, depth)

	symbol := MULTIPLE_ENTRIES

	if last_entry {
		symbol = LAST_ENTRY
	}

	return line + symbol + " " + entry_name
}
