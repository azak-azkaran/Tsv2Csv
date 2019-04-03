package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func TsvToCsv(text []string) string {
	var builder strings.Builder
	for _, line := range text {
		tmp := strings.ReplaceAll(line, "\t", "\",\"")
		builder.WriteString("\"")
		builder.WriteString(tmp)
		builder.WriteString("\"\n")
	}
	return builder.String()
}

func main() {
	Init(os.Stdout, os.Stdout, os.Stderr)

	Info.Println("starting")
	argsWithProg := os.Args

	i := len(argsWithProg)
	switch i {
	case 1:
		Error.Println("no file to read")
		panic("cant run without file")
	}

	for index, arg := range argsWithProg {
		if index == 0 {
			continue
		}
		Info.Println("file: ", arg)
		dat, err := ioutil.ReadFile(arg)
		if err != nil {
			Error.Fatalln("Could not read the file")
		}

		filename := strings.ReplaceAll(arg, "tsv", "csv")
		text := strings.Split(string(dat), "\n")

		Info.Println("Writing to: ", filename)
		err = ioutil.WriteFile(filename, []byte(TsvToCsv(text)), 0644)
		if err != nil {
			Error.Fatalln("Could not read the file")
		}
	}
}
