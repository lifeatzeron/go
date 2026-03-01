package main

// extension of the previous duplicate-line counter.
// It can read input either from standard input or from one or more files passed as command-line arguments.
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) // os.Open returns: (1) *os.File (2)error

			if err != nil {
				fmt.Fprintf(os.Stdin, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

// This design avoids code duplication.
// f *os.File → any file or stdin
// counts map[string]int → shared state
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}
