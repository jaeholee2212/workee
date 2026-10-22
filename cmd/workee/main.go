package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("workee", flag.ContinueOnError)
	flags.SetOutput(stderr)
	flags.Usage = func() {
		fmt.Fprintln(stderr, "Usage: workee [options]")
		fmt.Fprintln(stderr, "\nOptions:")
		fmt.Fprintln(stderr, "  -h, --help  Show help")
	}
	if err := flags.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		return 2
	}
	if flags.NArg() > 0 {
		fmt.Fprintf(stderr, "workee: unexpected argument %q\n", flags.Arg(0))
		return 2
	}

	fmt.Fprintln(stdout, "workee")
	return 0
}
