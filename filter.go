package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Masterminds/semver/v3"
	"github.com/urfave/cli/v3"
)

var filterCommand = &cli.Command{
	Name:      "filter",
	Usage:     "Filter semantic versions from stdin by constraint",
	ArgsUsage: "CONSTRAINT",
	Action:    filterAction,
}

func filterAction(ctx context.Context, cmd *cli.Command) error {
	if !cmd.Args().Present() {
		return fmt.Errorf("constraint argument is required")
	}

	constraint := cmd.Args().First()

	versions, err := readVersionsFromStdin(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read versions: %w", err)
	}

	filtered, err := filterVersions(versions, constraint)
	if err != nil {
		return fmt.Errorf("failed to filter versions: %w", err)
	}

	for _, v := range filtered {
		fmt.Println(v.Original())
	}

	return nil
}

// filterVersions filters semantic versions by a constraint
func filterVersions(versions []*semver.Version, constraint string) ([]*semver.Version, error) {
	constraints, err := semver.NewConstraint(constraint)
	if err != nil {
		return nil, fmt.Errorf("invalid constraint %q: %w", constraint, err)
	}

	var filtered []*semver.Version
	for _, v := range versions {
		if constraints.Check(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}
