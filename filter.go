package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/urfave/cli/v3"
)

var filterCommand = &cli.Command{
	Name:      "filter",
	Usage:     "Filter semantic versions from stdin by constraint",
	ArgsUsage: "CONSTRAINT",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "ignore-invalid",
			Aliases: []string{"i"},
			Usage:   "ignore invalid version lines instead of failing",
		},
	},
	Action: filterAction,
}

func filterAction(ctx context.Context, cmd *cli.Command) error {
	if !cmd.Args().Present() {
		return fmt.Errorf("constraint argument is required")
	}

	constraint := cmd.Args().First()
	ignoreInvalid := cmd.Bool("ignore-invalid")

	versions, err := readVersionsFromStdinWithOptions(os.Stdin, ignoreInvalid)
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

// readVersionsFromStdinWithOptions reads semantic versions from stdin with optional error handling
func readVersionsFromStdinWithOptions(r io.Reader, ignoreInvalid bool) ([]*semver.Version, error) {
	var versions []*semver.Version
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		v, err := semver.NewVersion(line)
		if err != nil {
			if ignoreInvalid {
				continue
			}
			return nil, fmt.Errorf("invalid version %q: %w", line, err)
		}
		versions = append(versions, v)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return versions, nil
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
