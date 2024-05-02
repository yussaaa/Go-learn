//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	// var hour, minute, second int
	components := strings.Split(input, ":")
	// _, err := fmt.Sscanf(input, "%d:%d:%d", &hour, &minute, &second)
	// if err != nil{
	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid time string", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{"Invalid hour", input}
		}
		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{"Invalid minute", input}
		}
		second, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{"Invalid second", input}
		}
		if hour < 0 || hour > 23 {
			return Time{}, &TimeParseError{"Invalid hour", input}
		}
		if minute < 0 || minute > 59 {
			return Time{}, &TimeParseError{"Invalid minute", input}
		}
		if second < 0 || second > 59 {
			return Time{}, &TimeParseError{"Invalid second", input}
		}

		return Time{hour, minute, second}, nil
	}

}
