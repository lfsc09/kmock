package fake

import (
	"math/rand/v2"
	"strings"
	"time"

	"github.com/lfsc09/kmock/internal/randkit"
)

const (
	dateDefaultFormat     = "2006-01-02"
	timeDefaultFormat     = "15:04:05"
	dateTimeDefaultFormat = "2006-01-02 15:04:05"
)

type Date struct {
	Rng *rand.Rand
}

// Date generates a random date between the provided 'from' and 'to' dates using the provided random number generator.
// The 'from' and 'to' parameters are expected to be in the format "YYYY-MM-DD".
// The 'format' parameter specifies the output format of the date.
func (d Date) Date(from string, to string, format string) string {
	if from == "" {
		from = "1970-01-01"
	}
	if to == "" {
		to = time.Now().Format(dateDefaultFormat)
	}
	if format == "" {
		format = dateDefaultFormat
	}
	fromTime, err := time.Parse(dateDefaultFormat, from)
	if err != nil {
		return ""
	}
	toTime, err := time.Parse(dateDefaultFormat, to)
	if err != nil {
		return ""
	}
	randomDate := randkit.RandomDateTime(d.Rng, fromTime, toTime)
	return randomDate.Format(parseFormat(format))
}

// Time generates a random time between the provided 'from' and 'to' times using the provided random number generator.
// The 'from' and 'to' parameters are expected to be in the format "hh:mm:ss" (24-hour format).
// The 'format' parameter specifies the output format of the time.
func (d Date) Time(from string, to string, format string) string {
	if from == "" {
		from = "00:00:00"
	}
	if to == "" {
		to = "23:59:59"
	}
	if format == "" {
		format = timeDefaultFormat
	}
	fromTime, err := time.Parse(timeDefaultFormat, from)
	if err != nil {
		return ""
	}
	toTime, err := time.Parse(timeDefaultFormat, to)
	if err != nil {
		return ""
	}
	randomTime := randkit.RandomDateTime(d.Rng, fromTime, toTime)
	return randomTime.Format(parseFormat(format))
}

// DateTime generates a random date and time between the provided 'from' and 'to' date-times using the provided random number generator.
// The 'from' and 'to' parameters are expected to be in the format "YYYY-MM-DD hh:mm:ss".
// The 'format' parameter specifies the output format of the date and time.
func (d Date) DateTime(from string, to string, format string) string {
	if from == "" {
		from = "1970-01-01 00:00:00"
	}
	if to == "" {
		to = time.Now().Format(dateTimeDefaultFormat)
	}
	if format == "" {
		format = dateTimeDefaultFormat
	}
	fromTime, err := time.Parse(dateTimeDefaultFormat, from)
	if err != nil {
		return ""
	}
	toTime, err := time.Parse(dateTimeDefaultFormat, to)
	if err != nil {
		return ""
	}
	randomDateTime := randkit.RandomDateTime(d.Rng, fromTime, toTime)
	return randomDateTime.Format(parseFormat(format))
}

// Now generates the current date and time in the specified format. If no format is provided, it defaults to "YYYY-MM-DD hh:mm:ss".
func (d Date) Now(format string) string {
	if format == "" {
		format = dateTimeDefaultFormat
	}
	return time.Now().Format(parseFormat(format))
}

// RuntimeDocs provides runtime documentation for the Date struct and its methods
func (d Date) RuntimeDocs() []*RunTimeDocs {
	return []*RunTimeDocs{
		{
			Domain:      "Date",
			Method:      "Date",
			Description: "Generates a random date, from/to: YYYY-MM-DD. Default format: YYYY-MM-DD",
			Params:      []string{"from", "to", "format"},
		},
		{
			Domain:      "Date",
			Method:      "Time",
			Description: "Generates a random time, from/to: hh:mm:ss. Default format: hh:mm:ss",
			Params:      []string{"from", "to", "format"},
		},
		{
			Domain:      "Date",
			Method:      "DateTime",
			Description: "Generates a random datetime, from/to: YYYY-MM-DD hh:mm:ss. Default format: YYYY-MM-DD hh:mm:ss",
			Params:      []string{"from", "to", "format"},
		},
		{
			Domain:      "Date",
			Method:      "Now",
			Description: "Generates the current datetime. Default format: YYYY-MM-DD hh:mm:ss",
			Params:      []string{"format"},
		},
	}
}

// parseFormat is a helper function to convert custom date format strings into Go's time layout format.
// For example, it can convert "YYYY-MM-DD" to "2006-01-02", "MM/DD/YYYY" to "01/02/2006", etc.
func parseFormat(format string) string {
	result := format
	replacements := map[string]string{
		"YYYY": "2006",
		"MM":   "01",
		"DD":   "02",
		"hh":   "15",
		"mm":   "04",
		"ss":   "05",
		"sss":  "000",
	}
	for key, value := range replacements {
		result = strings.ReplaceAll(result, key, value)
	}
	return result
}
