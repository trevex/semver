package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/urfave/cli/v3"
)

var sortCommand = &cli.Command{
	Name:   "sort",
	Usage:  "Sort semantic versions from stdin",
	Action: sortAction,
}

func sortAction(ctx context.Context, cmd *cli.Command) error {
	versions, err := readVersionsFromStdin(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read versions: %w", err)
	}

	sortedVersions, err := sortVersions(versions)
	if err != nil {
		return fmt.Errorf("failed to sort versions: %w", err)
	}

	for _, v := range sortedVersions {
		fmt.Println(v.Original())
	}

	return nil
}

// readVersionsFromStdin reads semantic versions from stdin, one per line
func readVersionsFromStdin(r io.Reader) ([]*semver.Version, error) {
	var versions []*semver.Version
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		v, err := semver.NewVersion(line)
		if err != nil {
			return nil, fmt.Errorf("invalid version %q: %w", line, err)
		}
		versions = append(versions, v)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return versions, nil
}

// sortVersions sorts semantic versions in ascending order
func sortVersions(versions []*semver.Version) ([]*semver.Version, error) {
	sorted := make([]*semver.Version, len(versions))
	copy(sorted, versions)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].LessThan(sorted[j])
	})

	return sorted, nil
}
