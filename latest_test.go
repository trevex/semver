package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLatestVersion(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
		wantErr  bool
	}{
		{
			name:     "basic latest",
			input:    []string{"1.0.0", "2.0.0", "3.0.0"},
			expected: "3.0.0",
			wantErr:  false,
		},
		{
			name:     "latest with unsorted input",
			input:    []string{"3.0.0", "1.5.0", "2.1.0", "1.0.0"},
			expected: "3.0.0",
			wantErr:  false,
		},
		{
			name:     "single version",
			input:    []string{"1.5.3"},
			expected: "1.5.3",
			wantErr:  false,
		},
		{
			name:     "latest with patch versions",
			input:    []string{"1.0.0", "1.0.1", "1.0.2", "1.0.10"},
			expected: "1.0.10",
			wantErr:  false,
		},
		{
			name:     "latest with minor versions",
			input:    []string{"1.5.0", "1.2.0", "1.10.0", "1.9.0"},
			expected: "1.10.0",
			wantErr:  false,
		},
		{
			name:     "latest with major versions",
			input:    []string{"2.0.0", "10.0.0", "5.0.0"},
			expected: "10.0.0",
			wantErr:  false,
		},
		{
			name:     "empty list",
			input:    []string{},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(strings.Join(tt.input, "\n"))
			versions, err := readVersionsFromStdin(r)
			assert.NoError(t, err)

			latest, err := getLatestVersion(versions)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, latest.String())
			}
		})
	}
}
