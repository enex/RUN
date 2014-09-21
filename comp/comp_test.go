package comp_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/rthornton128/calc/comp"
)

var ext string

func init() {
	ext = ""
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}
}

func TestSimpleExpression(t *testing.T) {
	test_handler(t, "(decl main int (+ 5 3))", "8")
}

func TestSimpleExpressionWithComments(t *testing.T) {
	test_handler(t, ";comment 1\n(decl main int (* 5 3)); comment 2", "15")
}

func TestComplexExpression(t *testing.T) {
	test_handler(t, "(decl main int (- (* 9 (+ 2 3)) (+ (/ 20 (% 15 10)) 1)))",
		"40")
}

func TestVarExpression(t *testing.T) {
	test_handler(t, "(decl main int ((var (= a 5)) a))", "5")
}

func test_handler(t *testing.T, src, expected string) {
	defer tearDown()

	err := ioutil.WriteFile("test.calc", []byte(src), os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	comp.CompileFile("test.calc")
	os.Remove("test.calc")

	runpath, _ := filepath.Abs("../runtime")
	runlib := filepath.Join(runpath, "runtime.a")
	out, err := exec.Command("gcc"+ext, "-Wall", "-Wextra", "-std=c99",
		"-I", runpath, "--output=test"+ext,
		"test.c", runlib).CombinedOutput()
	if err != nil {
		t.Log(string(out))
		t.Fatal(err)
	}
	var output []byte

	switch runtime.GOOS {
	case "windows":
		output, err = exec.Command("test" + ext).Output()
	default:
		output, err = exec.Command("./test").Output()
	}
	output = []byte(strings.TrimSpace(string(output)))
	t.Log("len output:", len(output))
	t.Log("len expected:", len(expected))

	if string(output) != expected {
		t.Fatal("For " + src + " expected " + expected + " got " + string(output))
	}
}

func tearDown() {
	os.Remove("test.c")
	os.Remove("test" + ext)
}
