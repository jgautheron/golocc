package main

import (
	"flag"
	"log"
	"os"

	"github.com/jgautheron/golocc"
)

func main() {
	// Default to current working dir
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working dir: ", err.Error())
	}

	targetDir := flag.String("d", pwd, "target directory")
	outputFmt := flag.String("o", "text", "output format")
	ignore := flag.String("ignore", "", "ignore files matching the given regular expression")
	flag.Parse()

	var report golocc.ReportInterface
	switch *outputFmt {
	case "text":
		report = &golocc.TextReport{Writer: os.Stdout}
	case "json":
		report = &golocc.JSONReport{Writer: os.Stdout}
	}

	path := *targetDir
	pathLen := len(path)
	recursive := false
	if pathLen >= 5 && path[pathLen-3:] == "..." {
		recursive = true
		path = path[:pathLen-3]
	}

	parser := golocc.New(path, *ignore, recursive)
	res, err := parser.ParseTree()
	if err != nil {
		log.Fatal(err)
		return
	}
	report.Print(res)
}
