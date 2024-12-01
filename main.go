package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const MULTIPLE_ENTRIES = "├──"
const LAST_ENTRY = "└──"
const BAR = "│"
const IDENTATION = 3

var (
	d = flag.String("d", ".", "directory")
)

var ignorables = []string{".git"}

func main() {
	flag.Parse()
	splited_dir := strings.Split(*d, "/")
	fmt.Println(splited_dir[len(splited_dir)-1])
	build_directory_tree(0, *d)
}

func build_directory_tree(depth int, dir string) {
	entries, err := os.ReadDir(dir)

	if err != nil {
		log.Fatalln("Error reading directory " + dir)
	}

	entries_len := len(entries)
	last_entry := false

	for i, e := range entries {
		suffix := ""
		if slices.Contains(ignorables, e.Name()) {
			continue
		}

		if i == entries_len-1 {
			last_entry = true
		}

		if e.IsDir() {
			suffix = "/"
		}

		fmt.Println(build_line(depth, e.Name()+suffix, last_entry))

		if e.IsDir() {
			dir_path := filepath.Join(dir, e.Name())
			build_directory_tree(depth+1, dir_path)
		}
	}
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
