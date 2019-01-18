package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	sh_words "github.com/mattn/go-shellwords"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "ERROR: At least 1 argument is required!\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Usage: env.exe [envKey=envValue...] <binaryPath> [flags]\n")
		fmt.Fprintf(os.Stderr, "  Example: env.exe FOO=example BAR=\"hello world\" binary --flags\n")
		os.Exit(1)
	}

	var env []string
	var program strings.Builder

	isProgram := false
	for _, arg := range os.Args[1:] {
		if !isProgram && strings.Contains(arg, "=") {
			env = append(env, arg)
			continue
		}

		isProgram = true
		program.WriteString(arg)
		program.WriteString(" ")
	}

	parsed, err := sh_words.Parse(program.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to parse binary & arguments: %v", err)
		os.Exit(1)
	}
	// fmt.Printf("env %s %#v\n", strings.Join(env, " "), parsed)

	var args []string
	binary := parsed[0]

	if len(parsed) > 1 {
		args = parsed[1:]
	}

	cmd := exec.Command(binary, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	for _, v := range env {
		cmd.Env = append(cmd.Env, v)
	}

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
