package golocc

import (
	"testing"
)

func TestCLOC(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.CLOC != 8 {
		t.Error("expected 8 comment lines")
	}
}

func TestStructCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.Struct != 2 {
		t.Error("expected 2 structs")
	}
}

func TestInterfaceCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.Interface != 1 {
		t.Error("expected 1 interface")
	}
}

func TestMethodCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.Method != 2 {
		t.Error("expected 2 methods")
	}

	if res.ExportedMethod != 1 {
		t.Error("expected 1 exported method")
	}
}

func TestMethodLineCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.MethodLOC != 15 {
		t.Error("expected 15 method lines got", res.MethodLOC)
	}
}

func TestFunctionCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.Function != 5 {
		t.Error("expected 5 functions got", res.Function)
	}

	if res.ExportedFunction != 4 {
		t.Error("expected 4 exported functions")
	}
}

func TestFunctionLineCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.FunctionLOC != 10 {
		t.Error("expected 10 function lines got", res.FunctionLOC)
	}
}

func TestImportCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.Import != 5 {
		t.Error("expected 5 imports got", res.Import)
	}
}

func TestTestCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.Test != 3 {
		t.Error("expected 3 tests got", res.Test)
	}
}

func TestAssertCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.Assertion != 2 {
		t.Error("expected 2 tests got", res.Test)
	}
}

func TestIfStatementCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.IfStatement != 2 {
		t.Error("expected 2 if statements got", res.IfStatement)
	}
}

func TestSwitchStatementCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.SwitchStatement != 1 {
		t.Error("expected 1 switch statement got", res.SwitchStatement)
	}
}

func TestGoStatementCount(t *testing.T) {
	parser := Parser{result: &Result{}}
	parser.parseDir("./fixture", "")
	res := parser.GetResult()

	if res.GoStatement != 2 {
		t.Error("expected 2 switch statement got", res.GoStatement)
	}
}
