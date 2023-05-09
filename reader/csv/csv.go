package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"sync"

	"emaildomainstats/reader"
	"emaildomainstats/utils"
)

var (
	_errHeaderNotFound = errors.New("header not found")
	_errColumnNotFound = errors.New("column not found")
)

type csvReader struct {
	file    *csv.Reader
	columns map[int]struct{}
}

// NewReader returns a new csv reader with the specified settings
func NewReader(filepath string, column []int, header []string, delimeter *rune) (reader.Reader, error) {
	csvFile, err := utils.OpenCSVFile(filepath)
	if err != nil {
		return nil, err
	}
	if delimeter != nil {
		csvFile.Comma = *delimeter
	}

	headers, err := csvFile.Read()
	if err != nil {
		return nil, err
	}

	columns := make(map[int]struct{})

	if column == nil && header == nil {
		return &csvReader{
			file:    csvFile,
			columns: columns,
		}, nil
	}

	for _, c := range column {
		if c > len(headers) || c < 0 {
			return nil, _errColumnNotFound
		}
		columns[c] = struct{}{}
	}

	for _, h := range header {
		c, ok := utils.FindHeader(h, headers)
		if !ok {
			return nil, _errHeaderNotFound
		}
		columns[c] = struct{}{}
	}

	return &csvReader{
		file:    csvFile,
		columns: columns,
	}, nil
}

// Read reads the csv file and returns a map of email domains and the number of customers for each domain.
func (c *csvReader) Read() (map[string]int, error) {
	var wg sync.WaitGroup
	result := NewResult()
	rowChannel := make(chan []string)

	searchFunction := func(row []string) ([]string, bool) {
		return utils.SearchAllColumnsForDomain(row)
	}

	if len(c.columns) > 0 {
		searchFunction = func(row []string) ([]string, bool) {
			return utils.SearchColumnsForDomain(row, c.columns)
		}
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for row := range rowChannel {
			searchRow := row
			wg.Add(1)
			go func() {
				defer wg.Done()
				if domains, ok := searchFunction(searchRow); ok {
					for _, domain := range domains {
						result.Add(domain)
					}
				}
			}()
		}
	}()

	for {
		row, err := c.file.Read()
		if err != nil {
			if err == io.EOF {
				close(rowChannel)
				break
			}
			return nil, err
		}
		rowChannel <- row
	}

	wg.Wait()
	return result.values, nil
}
