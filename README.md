# Email Domain Stats

## How it works:
Currently, only provides functionality for inputting csv files. The csv file can be in any format.
By default, the app will search through every column of the provided csv file and find any that contain an @ symbol.
It will then count the number of times each domain appears and output the result.

## How to use it:
Provide the file name to the ProcessCSVFile function.
If you wish to specify either a specific column or a specific delimiter, you can do so by providing the column number and delimiter to the ProcessCSVFile function.

The following functions are available as options:
- WithDelimiter(delimiter string) - allows you to specify a delimiter for the csv file
- WithColumns(column []int) - allows you to specify a set of columns to search for email addresses in
- WithColumnHeaders(columnHeader []string) - allows you to specify a set of column headers to search for email addresses in

### Output
The output is a map of email domains to the number of times they appear in the csv file.

## Example:
```go
package main

import (
	"encoding/json"
	"log"
	"os"

	"emaildomainstats"
)

func main() {
	results, err := emaildomainstats.ProcessCSVFile("customer_data.csv", emaildomainstats.WithColumns([]int{1,2,3}))
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("./example/results.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
```
