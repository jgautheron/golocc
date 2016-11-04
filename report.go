package golocc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

//ReportInterface - reports that parse results and print out a report
type ReportInterface interface {
	Print(*Result)
}

//JSONReport json structure for LOC report
type JSONReport struct {
	Writer io.Writer
}

//Print - print out parsed report in json format
func (j *JSONReport) Print(res *Result) {
	jsonOutput, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(j.Writer, string(jsonOutput))
}

//TextReport - plaintext report output
type TextReport struct {
	Writer io.Writer
}

//Print - print out plaintext report
func (t *TextReport) Print(res *Result) {
	fmt.Fprintf(t.Writer, "\n")
	fmt.Fprintln(t.Writer, strings.Repeat("-", 80))
	fmt.Fprintf(t.Writer, "Number of Files: %v\n", res.Files)
	fmt.Fprintf(t.Writer, "Lines of Code: %v (%v CLOC, %v NCLOC)\n", res.LOC, res.CLOC, res.NCLOC)
	fmt.Fprintf(t.Writer, "Imports:       %v\n", res.Import)
	fmt.Fprintf(t.Writer, "Structs:       %v\n", res.Struct)
	fmt.Fprintf(t.Writer, "Interfaces:    %v\n", res.Interface)
	fmt.Fprintf(t.Writer, "Methods:       %v (%v Exported, %v LOC, %v LOC Avg.)\n", res.Method, res.ExportedMethod, res.MethodLOC, t.divideBy(res.MethodLOC, res.Method))
	fmt.Fprintf(t.Writer, "Functions:     %v (%v Exported, %v LOC, %v LOC Avg.)\n", res.Function, res.ExportedFunction, res.FunctionLOC, t.divideBy(res.FunctionLOC, res.Function))
	fmt.Fprintln(t.Writer, strings.Repeat("-", 80))
	fmt.Fprintf(t.Writer, "Ifs:           %v \n", res.IfStatement)
	fmt.Fprintf(t.Writer, "Switches:      %v \n", res.SwitchStatement)
	fmt.Fprintf(t.Writer, "Go Routines:   %v \n", res.GoStatement)
	fmt.Fprintln(t.Writer, strings.Repeat("-", 80))
	fmt.Fprintf(t.Writer, "Tests:         %v \n", res.Test)
	fmt.Fprintf(t.Writer, "Assertions:    %v \n", res.Assertion)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Fprintf(t.Writer, "\n")
}

//Safely divide where vivisor might be zero
func (t *TextReport) divideBy(num, dividedBy int) int {
	if dividedBy == 0 {
		return 0
	}
	return num / dividedBy
}
