package viite_test

import (
	"fmt"
	"testing"

	"kkn.fi/viite"
)

func ExampleFormat() {
	fmt.Println(viite.Format("12303216"))
	// Output: 123 03216
}

func TestFormat(t *testing.T) {
	var testData = []struct {
		input    string
		expected string
	}{
		{"12303216", "123 03216"},
		{"123032161234", "12 30321 61234"},
		{"1234567890123456789", "1234 56789 01234 56789"},
		{"123456789012345678901234", "1234 56789 01234 56789 01234"},
		{"1234", "1234"},
		{"3629466132657495", "3 62946 61326 57495"},
	}
	for _, tc := range testData {
		v := viite.Format(tc.input)
		if v != tc.expected {
			t.Fatalf("expected '%v', got '%v'", tc.expected, v)
		}
	}
}

func ExampleGenerate() {
	fmt.Println(viite.Generate("123123"))
	// Output: 1231234 <nil>
}

func TestGenerate(t *testing.T) {
	var testData = []struct {
		input    string
		expected string
	}{
		{"1230321", "12303216"},

		{"1231", "12315"},
		{"12481", "124818"},
		{"12345678", "123456780"},
		{"123456781", "1234567813"},
		{"36226134", "362261341"},
		{"362946613265749", "3629466132657495"},
		{"1234567890123456781", "12345678901234567810"},
	}
	for _, tc := range testData {
		v, err := viite.Generate(tc.input)
		if err != nil {
			t.Fatalf("expected nil error, got '%v'", err)
		}
		if v != tc.expected {
			t.Fatalf("expected '%v', got '%v'", tc.expected, v)
		}
	}
}

func TestGenerateErrInvalidInput(t *testing.T) {
	var testData = []struct {
		input string
	}{
		{"123e"},
		{"1"},
		{"12"},
		{"123456789012345678901234"},
	}
	for _, tc := range testData {
		_, err := viite.Generate(tc.input)
		if err != viite.ErrInvalidInput {
			t.Fatalf("expected viite.ErrInvalidInput, got %v with '%v'", err, tc.input)
		}
	}
}

func ExampleValidate() {
	fmt.Println(viite.Validate("12303216"))
	// Output: true
}

func TestValidateFails(t *testing.T) {
	var testData = []struct {
		input string
	}{
		{"123"},
		{"123456789012345678901"},
		{"123x"},
	}
	for _, tc := range testData {
		if invalid := viite.Validate(tc.input); invalid {
			t.Fatalf("expected validation to fail with '%v' ", tc.input)
		}
	}
}
func TestValidateSucceeds(t *testing.T) {
	var testData = []struct {
		input string
	}{
		{"12303216"},
		{"12315"},
		{"124818"},
		{"123456780"},
		{"1234567813"},
		{"362261341"},
		{"3629466132657495"},
		{"12345678901234567810"},
	}
	for _, tc := range testData {
		if valid := viite.Validate(tc.input); !valid {
			t.Fatalf("expected validation to succeed with '%v' ", tc.input)
		}
	}
}
