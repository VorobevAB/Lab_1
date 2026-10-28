package main

import (
	"bufio"
	"fmt"
	"os"
)

func load(db *DB, path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			exec(db, line)
		}
	}
}

func save(path, cmd string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	fmt.Fprintln(f, cmd)
}

func main() {
	var file, query string

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--file":
			if i+1 < len(os.Args) {
				i++
				file = os.Args[i]
			}
		case "--query":
			if i+1 < len(os.Args) {
				i++
				query = os.Args[i]
			}
		}
	}

	db := NewDB()

	if file != "" {
		load(db, file)
	}

	if query != "" {
		res := exec(db, query)
		fmt.Println(res)
		if file != "" && (len(res) < 7 || res[:7] != "ОШИБКА") {
			save(file, query)
		}
	}
}
