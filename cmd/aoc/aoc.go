package main

import (
	"flag"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	CREATE_COMMAND = "create"
)

func main() {
	createCmd := flag.NewFlagSet(CREATE_COMMAND, flag.ExitOnError)
	year := createCmd.String("y", "", "Year of the challenge")
	day := createCmd.String("d", "", "Day of the challenge")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case CREATE_COMMAND:
		createCmd.Parse(os.Args[2:])
		create(*year, *day)
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func create(year, day string) {
	if year == "" {
		year = strconv.Itoa(time.Now().Year())
	}
	if day == "" {
		day = strconv.Itoa(time.Now().Day())
	}
	if len(day) == 1 {
		day = "0" + day
	}

	dir := path.Join(year, day)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}

	files := []string{
		"solve.go",
		"input.txt",
		"sample.txt",
	}
	for _, file := range files {
		filePath := path.Join(dir, file)
		_, err := os.Stat(filePath)
		if err == nil {
			log.Printf("warn: file already exists, skipping... (%s)", filePath)
			continue
		}
		_, err = os.Create(filePath)
		if err != nil {
			panic(err)
		}
		if file == "solve.go" {
			f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			template, err := os.ReadFile("template.txt")
			if err != nil {
				panic(err)
			}
			f.Write(template)
		}
	}

	log.Printf("Created directory %s/%s\n", year, day)
}
