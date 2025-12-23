package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterVersions(t *testing.T) {
	tests := []struct {
		name       string
		input      []string
		constraint string
		expected   []string
		wantErr    bool
	}{
		{
			name:       "filter by exact version",
			input:      []string{"1.0.0", "2.0.0", "3.0.0"},
			constraint: "1.0.0",
			expected:   []string{"1.0.0"},
			wantErr:    false,
		},
		{
			name:       "filter by minor version constraint",
			input:      []string{"1.0.0", "1.1.0", "1.2.0", "2.0.0"},
			constraint: "^1.0.0",
			expected:   []string{"1.0.0", "1.1.0", "1.2.0"},
			wantErr:    false,
		},
		{
			name:       "filter by range constraint",
			input:      []string{"1.0.0", "1.5.0", "2.0.0", "3.0.0"},
			constraint: ">=1.0.0, <2.0.0",
			expected:   []string{"1.0.0", "1.5.0"},
			wantErr:    false,
		},
		{
			name:       "filter with no matches",
			input:      []string{"1.0.0", "1.1.0"},
			constraint: "2.0.0",
			expected:   []string{},
			wantErr:    false,
		},
		{
			name:       "tilde constraint",
			input:      []string{"1.2.0", "1.2.1", "1.2.5", "1.3.0", "2.0.0"},
			constraint: "~1.2.0",
			expected:   []string{"1.2.0", "1.2.1", "1.2.5"},
			wantErr:    false,
		},
		{
			name:       "invalid constraint",
			input:      []string{"1.0.0"},
			constraint: "invalid!!!",
			expected:   nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(strings.Join(tt.input, "\n"))
			versions, err := readVersionsFromStdin(r)
			assert.NoError(t, err)

			filtered, err := filterVersions(versions, tt.constraint)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				result := make([]string, len(filtered))
				for i, v := range filtered {
					result[i] = v.String()
				}
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
