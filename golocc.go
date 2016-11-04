package golocc

import (
	"bufio"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Result - container for analysis results
type Result struct {
	LOC   int
	CLOC  int
	NCLOC int

	Struct    int
	Interface int

	Method           int
	ExportedMethod   int
	MethodLOC        int
	Function         int
	ExportedFunction int
	FunctionLOC      int

	Import int

	IfStatement     int
	SwitchStatement int
	GoStatement     int

	Test      int
	Assertion int

	Files int
}

// Parser - Code parser struct
type Parser struct {
	result *Result
	path   string
	ignore string
}

func New(path, ignore string) *Parser {
	return &Parser{
		result: &Result{},
		path:   path,
		ignore: ignore,
	}
}

func (p *Parser) ParseTree() (*Result, error) {
	pathLen := len(p.path)

	// Parse recursively the given path if the recursive notation is found
	if pathLen >= 5 && p.path[pathLen-3:] == "..." {
		filepath.Walk(p.path[:pathLen-3], func(fp string, f os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				return nil
			}

			if f.IsDir() {
				p.parseDir(fp, p.ignore)
			}
			return nil
		})
	} else {
		p.parseDir(p.path, p.ignore)
	}

	return p.result, nil
}

//parseDir - Parse all files within directory
func (p *Parser) parseDir(targetDir, ignore string) error {
	//create the file set
	fset := token.NewFileSet()
	d, err := parser.ParseDir(fset, targetDir, func(info os.FileInfo) bool {
		valid, name := true, info.Name()

		if len(ignore) != 0 {
			match, err := regexp.MatchString(ignore, targetDir+name)
			if err != nil {
				log.Fatal(err)
				return true
			}
			if match {
				valid = false
			}
		}

		return valid
	}, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	//count up lines
	fset.Iterate(func(file *token.File) bool {
		loc, cloc, assertions := p.countLOC(file.Name())
		p.result.LOC += loc
		p.result.CLOC += cloc
		p.result.NCLOC += (loc - cloc)
		p.result.Assertion += assertions
		p.result.Files++
		return true
	})

	//setup visitors
	var visitors []AstVisitor
	visitors = append(
		visitors,
		&TypeVisitor{res: p.result},
		&FuncVisitor{res: p.result, fset: fset},
		&ImportVisitor{res: p.result},
		&FlowControlVisitor{res: p.result})

	//count entities
	for _, pkg := range d {
		ast.Inspect(pkg, func(n ast.Node) bool {
			for _, vis := range visitors {
				vis.Visit(n)
			}
			return true
		})
	}

	return nil
}

//CountLOC - count lines of code, pull LOC, Comments, assertions
func (p *Parser) countLOC(filePath string) (int, int, int) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var loc int
	var cloc int
	var assertions int
	var inBlockComment bool

	assertionPrefixes := []string{
		"So(",
		"convey.So(",
		"assert.",
	}

	for {
		line, isPrefix, err := r.ReadLine()
		if err == io.EOF {
			return loc, cloc, assertions
		}
		if isPrefix == true {
			continue //incomplete line
		}
		if len(line) == 0 {
			continue //empty line
		}
		if strings.Index(strings.TrimSpace(string(line)), "//") == 0 {
			cloc++ //slash comment at start of line
			continue
		}
		for _, prefix := range assertionPrefixes {
			if strings.HasPrefix(strings.TrimSpace(string(line)), prefix) {
				assertions++
			}
		}

		blockCommentStartPos := strings.Index(strings.TrimSpace(string(line)), "/*")
		blockCommentEndPos := strings.LastIndex(strings.TrimSpace(string(line)), "*/")

		if blockCommentStartPos > -1 {
			//block was started and not terminated
			if blockCommentEndPos == -1 || blockCommentStartPos > blockCommentEndPos {
				inBlockComment = true
			}
		}
		if blockCommentEndPos > -1 {
			//end of block is found and no new block was started
			if blockCommentStartPos == -1 || blockCommentEndPos > blockCommentStartPos {
				inBlockComment = false
				cloc++ //end of block counts as a comment line but we're already out of the block
			}
		}

		loc++
		if inBlockComment {
			cloc++
		}
	}
}

func (p *Parser) GetResult() *Result {
	return p.result
}
