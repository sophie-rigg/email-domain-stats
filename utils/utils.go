package utils

import (
	"encoding/csv"
	"os"
	"strings"
)

// FindHeader finds a header in a slice of headers
func FindHeader(header string, headers []string) (int, bool) {
	header = strings.ToLower(header)
	for i, h := range headers {
		if strings.ToLower(h) == header {
			return i, true
		}
	}
	return 0, false
}

// OpenCSVFile opens a csv file and returns a csv reader
func OpenCSVFile(filepath string) (*csv.Reader, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return csv.NewReader(file), nil
}

// SearchColumnsForDomain searches a slice of columns for a domain
func SearchColumnsForDomain(headers []string, columns map[int]struct{}) ([]string, bool) {
	var domains []string
	var found bool
	for i, h := range headers {
		if _, ok := columns[i]; ok {
			header, ok := findDomain(h)
			if ok {
				domains = append(domains, header)
				found = true
			}
		}
	}
	return domains, found
}

// SearchAllColumnsForDomain searches all columns for a domain
func SearchAllColumnsForDomain(headers []string) ([]string, bool) {
	var domains []string
	var found bool
	for _, h := range headers {
		header, ok := findDomain(h)
		if ok {
			domains = append(domains, header)
			found = true
		}
		continue
	}
	return domains, found
}

// findDomain finds an @ symbol in a string
func findDomain(entry string) (string, bool) {
	values := strings.Split(entry, "@")
	if len(values) == 1 {
		return "", false
	}
	return values[1], true
}
