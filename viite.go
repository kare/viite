package viite

import (
	"errors"
	"fmt"
	"unicode"
)

const (
	// minLen is the minimum length of valid viitenumero.
	minLen = 4
	// maxLen is the maximum length of valid viitenumero.
	maxLen = 20
)

// ErrInvalidInput represents an error where user input is invalid.
var ErrInvalidInput = errors.New("viite: invalid input")

// Format returns given v formatted.
func Format(v string) string {
	// g is the length of the format group
	const g = 5
	if len(v) < g {
		return v
	}
	s := ""
	lo := len(v) - g
	for i := len(v); i >= 0; i -= g {
		s = v[lo:i] + s
		if lo > 0 {
			s = " " + s
		}
		if lo-g < 0 {
			lo = 0
		} else {
			lo -= g
		}
	}
	return s
}

// Generate returns a Finnish viitenumero based on given v.
// Viitenumero checksum is appended to v and returned.
// Returns ErrInvalidInput if given v doesn't only contain digits.
func Generate(v string) (string, error) {
	if len(v) < minLen-1 || len(v) > maxLen-1 {
		return "", ErrInvalidInput
	}
	if !digitsOnly(v) {
		return "", ErrInvalidInput
	}
	return fmt.Sprintf("%v%v", v, checksum(v)), nil
}

// digitsOnly returns true for strings that contain only digits and false otherwise.
func digitsOnly(v string) bool {
	for _, r := range v {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// Validate returns true for valid input and false otherwise.
func Validate(v string) bool {
	if len(v) < minLen || len(v) > maxLen {
		return false
	}
	if !digitsOnly(v) {
		return false
	}

	// given checksum
	gc := int(v[len(v)-1:][0] - '0')
	// correct checksum
	cc := checksum(v[:len(v)-1])
	return gc == cc
}

// checksum calculates the checksum for a Finnish viitenumero base number.
// Function assumes digit only input.
func checksum(b string) int {
	weights := []int{7, 3, 1}
	var checksum int
	for i, r := range reverse([]rune(b)) {
		checksum += int(r-'0') * weights[i%3]
	}
	return (10 - checksum%10) % 10
}

// reverse reverses a slice of runes.
func reverse(v []rune) []rune {
	n := len(v)
	for i := 0; i < n/2; i++ {
		v[i], v[n-1-i] = v[n-1-i], v[i]
	}
	return v
}
