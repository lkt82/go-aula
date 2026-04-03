package aulaapi

import "strings"

// SubjectNames maps Danish school subject acronyms to their full names.
var SubjectNames = map[string]string{
	"BIL": "Billedkunst",
	"DAN": "Dansk",
	"ENG": "Engelsk",
	"FYS": "Fysik/Kemi",
	"GEO": "Geografi",
	"HDS": "Håndværk og Design",
	"HIS": "Historie",
	"IDR": "Idræt",
	"KLA": "Klassens tid",
	"KRI": "Kristendomskundskab",
	"MAD": "Madkundskab",
	"MAT": "Matematik",
	"MUS": "Musik",
	"N/T": "Natur/Teknologi",
	"PÆD": "Pædagogisk tid",
	"SAM": "Samfundsfag",
	"STØ": "Støtte",
	"TYS": "Tysk",
	"FRA": "Fransk",
}

// ExpandSubject expands a subject acronym to its full Danish name.
// Returns the original string if no match is found.
func ExpandSubject(acronym string) string {
	if full, ok := SubjectNames[strings.TrimSpace(acronym)]; ok {
		return full
	}
	return acronym
}

// ExpandTitle expands subject acronyms in an event title.
// Handles plain acronyms ("MAT") and prefixed ones ("P4 - MAT").
func ExpandTitle(title string) string {
	if parts := strings.SplitN(title, " - ", 2); len(parts) == 2 {
		return parts[0] + " - " + ExpandSubject(parts[1])
	}
	return ExpandSubject(title)
}
