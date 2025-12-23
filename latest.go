package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Masterminds/semver/v3"
	"github.com/urfave/cli/v3"
)

var latestCommand = &cli.Command{
	Name:   "latest",
	Usage:  "Find the highest semantic version from stdin",
	Action: latestAction,
}

func latestAction(ctx context.Context, cmd *cli.Command) error {
	versions, err := readVersionsFromStdin(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read versions: %w", err)
	}

	if len(versions) == 0 {
		return fmt.Errorf("no versions provided")
	}

	latest, err := getLatestVersion(versions)
	if err != nil {
		return fmt.Errorf("failed to find latest version: %w", err)
	}

	fmt.Println(latest.Original())

	return nil
}

// getLatestVersion returns the highest semantic version from a slice
func getLatestVersion(versions []*semver.Version) (*semver.Version, error) {
	if len(versions) == 0 {
		return nil, fmt.Errorf("no versions provided")
	}

	latest := versions[0]
	for _, v := range versions[1:] {
		if v.GreaterThan(latest) {
			latest = v
		}
	}

	return latest, nil
}
