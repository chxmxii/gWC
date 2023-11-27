package pkg

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

func WordCount() {
	// Declare and parse flags
	var (
		printBytesCount, printCharCount, printLineCount, printWordCount bool
		filename                                                        string
	)

	flag.BoolVar(&printBytesCount, "b", false, "Print bytes count")
	flag.BoolVar(&printCharCount, "c", false, "Print character count")
	flag.BoolVar(&printLineCount, "l", false, "Print line count")
	flag.BoolVar(&printWordCount, "w", false, "Print word count")
	flag.Parse()

	// Set filename as the first argument
	filename = flag.CommandLine.Arg(0)

	// Set default flags if none are specified
	if !printBytesCount && !printWordCount && !printCharCount && !printLineCount {
		printLineCount, printWordCount, printBytesCount, printCharCount = true, true, true, true
	}

	// Read file content
	var file []byte
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Printf("Error reading file [%s]: %s\n", file, err)
		os.Exit(1)
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		file, err = io.ReadAll(os.Stdin)
	} else {
		file, err = os.ReadFile(filename)
	}

	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filename, err)
		os.Exit(2)
	}

	// Count and print statistics
	if printBytesCount {
		fmt.Printf("Bytes => %d\n", len(file))
	}

	if printCharCount {
		chars := utf8.RuneCount(file)
		fmt.Printf("Chars => %d\n", chars)
	}

	if printLineCount {
		lines := bytes.Count(file, []byte{'\n'})
		fmt.Printf("Lines => %d\n", lines)
	}

	if printWordCount {
		words := len(bytes.Fields(file))
		fmt.Printf("Words => %d\n\n", words)
	}
	fmt.Print("Thank you for using gWC!\n\n")
}
