package emaildomainstats

type Option func(c *client)

// WithColumnHeaders is an option to specify the column headers examined
func WithColumnHeaders(column []string) Option {
	return func(c *client) {
		c.columnHeaders = column
	}
}

// WithColumns is an option to specify the column numbers examined
func WithColumns(column []int) Option {
	return func(c *client) {
		c.columnNumbers = column
	}
}

// WithDelimeter is an option to specify the delimeter used in the csv file
func WithDelimeter(delimeter rune) Option {
	return func(c *client) {
		c.delimeter = &delimeter
	}
}
