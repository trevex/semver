package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadVersionsFromStdin(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantLen int
		wantErr bool
	}{
		{
			name:    "valid versions",
			input:   "1.0.0\n2.1.0\n1.5.0\n",
			wantLen: 3,
			wantErr: false,
		},
		{
			name:    "empty lines are skipped",
			input:   "1.0.0\n\n2.1.0\n\n",
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "invalid version",
			input:   "1.0.0\ninvalid\n2.1.0\n",
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			versions, err := readVersionsFromStdin(r)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, versions, tt.wantLen)
			}
		})
	}
}

func TestSortVersions(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "basic sort",
			input:    []string{"2.0.0", "1.0.0", "3.0.0"},
			expected: []string{"1.0.0", "2.0.0", "3.0.0"},
		},
		{
			name:     "mixed patch versions",
			input:    []string{"1.0.2", "1.0.1", "1.0.3"},
			expected: []string{"1.0.1", "1.0.2", "1.0.3"},
		},
		{
			name:     "mixed minor versions",
			input:    []string{"1.2.0", "1.1.0", "1.3.0"},
			expected: []string{"1.1.0", "1.2.0", "1.3.0"},
		},
		{
			name:     "mixed major versions",
			input:    []string{"2.0.0", "1.0.0", "3.0.0"},
			expected: []string{"1.0.0", "2.0.0", "3.0.0"},
		},
		{
			name:     "complex versions",
			input:    []string{"1.9.0", "1.10.0", "1.2.0", "2.0.0"},
			expected: []string{"1.2.0", "1.9.0", "1.10.0", "2.0.0"},
		},
		{
			name:     "single version",
			input:    []string{"1.0.0"},
			expected: []string{"1.0.0"},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(strings.Join(tt.input, "\n"))
			versions, err := readVersionsFromStdin(r)
			assert.NoError(t, err)

			sorted, err := sortVersions(versions)
			assert.NoError(t, err)

			result := make([]string, len(sorted))
			for i, v := range sorted {
				result[i] = v.String()
			}

			assert.Equal(t, tt.expected, result)
		})
	}
}
