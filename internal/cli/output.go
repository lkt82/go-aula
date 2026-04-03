package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"golang.org/x/term"
)

// ---------------------------------------------------------------------------
// Color support
// ---------------------------------------------------------------------------

var (
	colorsOnce    sync.Once
	colorsEnabled bool
)

// ColorsEnabled returns true when the terminal supports color output.
// Color is disabled when NO_COLOR is set or stdout is not a terminal.
func ColorsEnabled() bool {
	colorsOnce.Do(func() {
		if _, ok := os.LookupEnv("NO_COLOR"); ok {
			colorsEnabled = false
			return
		}
		colorsEnabled = term.IsTerminal(int(os.Stdout.Fd()))
	})
	return colorsEnabled
}

// ANSI escape helpers -- return plain text when color is off.

func Bold(s string) string {
	if ColorsEnabled() {
		return "\033[1m" + s + "\033[0m"
	}
	return s
}

func Red(s string) string {
	if ColorsEnabled() {
		return "\033[31m" + s + "\033[0m"
	}
	return s
}

func Green(s string) string {
	if ColorsEnabled() {
		return "\033[32m" + s + "\033[0m"
	}
	return s
}

func Yellow(s string) string {
	if ColorsEnabled() {
		return "\033[33m" + s + "\033[0m"
	}
	return s
}

func Dim(s string) string {
	if ColorsEnabled() {
		return "\033[2m" + s + "\033[0m"
	}
	return s
}

// ---------------------------------------------------------------------------
// JSON output
// ---------------------------------------------------------------------------

// PrintJSON prints any value as pretty-printed JSON to stdout.
func PrintJSON(value any) {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Printf("{\"error\": \"serialization failed: %s\"}\n", err)
		return
	}
	fmt.Println(string(data))
}

// ---------------------------------------------------------------------------
// Text helpers
// ---------------------------------------------------------------------------

// Truncate truncates a string to at most max characters, appending "..." if needed.
func Truncate(s string, max int) string {
	runes := []rune(s)
	if len(runes) <= max {
		return s
	}
	if max <= 3 {
		return string(runes[:max])
	}
	return string(runes[:max-3]) + "..."
}

// FormatDatetime extracts date+time from an ISO datetime string.
// "2024-01-15T08:30:00" -> "2024-01-15 08:30"
func FormatDatetime(s string) string {
	if len(s) >= 16 {
		return strings.Replace(s[:16], "T", " ", 1)
	}
	return s
}

// SplitDatetime extracts date and time portions from an ISO datetime string.
func SplitDatetime(datetime *string) (string, string) {
	if datetime == nil {
		return "", ""
	}
	dt := *datetime
	if len(dt) >= 16 {
		return dt[:10], dt[11:16]
	}
	if len(dt) >= 10 {
		return dt[:10], ""
	}
	return dt, ""
}

// ExtractTime extracts HH:MM from datetime strings like "2024-01-15T08:30:00".
func ExtractTime(s string) string {
	idx := strings.Index(s, "T")
	if idx >= 0 {
		timePart := s[idx+1:]
		if len(timePart) >= 5 {
			return timePart[:5]
		}
	}
	return s
}

// StripHTMLTags strips HTML tags and decodes common HTML entities.
func StripHTMLTags(html string) string {
	var result strings.Builder
	result.Grow(len(html))
	inTag := false
	for _, ch := range html {
		switch {
		case ch == '<':
			inTag = true
		case ch == '>':
			inTag = false
		case !inTag:
			result.WriteRune(ch)
		}
	}
	s := result.String()
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "&gt;", ">")
	s = strings.ReplaceAll(s, "&quot;", "\"")
	s = strings.ReplaceAll(s, "&#39;", "'")
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	return s
}

// ---------------------------------------------------------------------------
// Status coloring
// ---------------------------------------------------------------------------

// ColorPresenceStatus applies color to a presence status string.
func ColorPresenceStatus(status string) string {
	lower := strings.ToLower(status)
	switch {
	case strings.Contains(lower, "sick"):
		return Red(status)
	case strings.Contains(lower, "present") && !strings.Contains(lower, "not"):
		return Green(status)
	case strings.Contains(lower, "absence") || strings.Contains(lower, "notpresent") || strings.Contains(lower, "not"):
		return Yellow(status)
	default:
		return status
	}
}

// UnreadMarker returns a bold asterisk for unread items.
func UnreadMarker(isRead bool) string {
	if isRead {
		return " "
	}
	return Bold("*")
}

// ---------------------------------------------------------------------------
// Table builder
// ---------------------------------------------------------------------------

// Column defines a table column.
type Column struct {
	Header string
	Width  int
}

// Table is a simple column-aligned table printer.
type Table struct {
	Columns []Column
}

// NewTable creates a new table with the given columns.
func NewTable(columns []Column) *Table {
	return &Table{Columns: columns}
}

// PrintHeader prints the header row and separator line.
func (t *Table) PrintHeader() {
	parts := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		parts[i] = fmt.Sprintf("%-*s", c.Width, c.Header)
	}
	fmt.Println(Bold(strings.Join(parts, " ")))

	totalWidth := 0
	for _, c := range t.Columns {
		totalWidth += c.Width
	}
	totalWidth += len(t.Columns) - 1
	fmt.Println(Dim(strings.Repeat("-", totalWidth)))
}

// PrintRow prints a data row. Values are truncated to column width.
func (t *Table) PrintRow(values []string) {
	parts := make([]string, len(t.Columns))
	for i, col := range t.Columns {
		val := ""
		if i < len(values) {
			val = values[i]
		}
		display := Truncate(val, col.Width)
		parts[i] = fmt.Sprintf("%-*s", col.Width, display)
	}
	fmt.Println(strings.Join(parts, " "))
}

// PrintColoredRow prints a row where display_values may contain ANSI codes.
// raw_values are used for width calculation, display_values for output.
func (t *Table) PrintColoredRow(rawValues, displayValues []string) {
	var parts []string
	for i, col := range t.Columns {
		raw := ""
		display := ""
		if i < len(rawValues) {
			raw = rawValues[i]
		}
		if i < len(displayValues) {
			display = displayValues[i]
		}
		rawTruncated := Truncate(raw, col.Width)
		rawLen := len([]rune(rawTruncated))
		padding := col.Width - rawLen
		if padding < 0 {
			padding = 0
		}
		displayTruncated := display
		if len([]rune(raw)) > col.Width {
			displayTruncated = Truncate(raw, col.Width)
		}
		parts = append(parts, displayTruncated+strings.Repeat(" ", padding))
	}
	fmt.Println(strings.Join(parts, " "))
}

// ---------------------------------------------------------------------------
// Pagination hint
// ---------------------------------------------------------------------------

// PrintPaginationHint prints a pagination hint to stderr.
func PrintPaginationHint(currentPage *int, hasMore bool, flag string) {
	if hasMore {
		page := 0
		if currentPage != nil {
			page = *currentPage
		}
		fmt.Fprintf(os.Stderr, "\n(more available -- use %s %d)\n", flag, page+1)
	}
}
