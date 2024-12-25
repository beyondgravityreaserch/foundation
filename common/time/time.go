package time

import "fmt"

// `Seconds“ formats a floating-point number representing seconds into a string
// with two decimal places.
// For example, an input of 3.14159 will return "3.14".
func Seconds(seconds float64) string {
	return fmt.Sprintf("%.2f", seconds)
}
