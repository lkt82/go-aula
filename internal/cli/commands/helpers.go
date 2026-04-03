package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// mustParseInt64 parses a string as int64 or exits with an error.
func mustParseInt64(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: invalid integer '%s': %v\n", s, err)
		os.Exit(1)
	}
	return v
}

// joinStrings joins strings with a separator (simple helper).
func joinStrings(ss []string, sep string) string {
	return strings.Join(ss, sep)
}

// repeatChar repeats a character n times.
func repeatChar(ch rune, n int) string {
	return strings.Repeat(string(ch), n)
}

// trimString trims leading and trailing whitespace.
func trimString(s string) string {
	return strings.TrimSpace(s)
}
