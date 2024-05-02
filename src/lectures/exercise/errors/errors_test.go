package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"12:34:56", true},
		{"12:34", false},
		{"12:34:56:78", false},
		{"12:34:78", false},
		{"12:78:56", false},
		{"78:34:56", false},
		{"", false},
	}
	for _, test := range tests {
		_, err := ParseTime(test.input)
		if test.want && err != nil {
			t.Errorf("ParseTime(%q) returned error %v, want %v", test.input, err, test.want)
		}
	}
}
