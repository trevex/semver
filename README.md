# semver

A simple semantic versioning CLI tool with helpers for common semver operations.

## Features

- **sort** - Sort semantic versions in ascending order
- **filter** - Filter versions matching a semantic version constraint
- **latest** - Find the highest version from a list

## Installation

```bash
go build -o semver
```

## Usage

### sort

Sort semantic versions from stdin:

```bash
echo -e "3.0.0\n1.5.0\n2.1.0" | ./semver sort
```

Output:
```
1.5.0
2.1.0
3.0.0
```

### filter

Filter versions matching a constraint. Supports standard semver constraint syntax:
- Exact version: `1.0.0`
- Caret: `^1.0.0` (allows patch and minor updates)
- Tilde: `~1.2.0` (allows patch updates only)
- Range: `>=1.0.0, <2.0.0`

```bash
echo -e "0.9.0\n1.0.0\n1.1.0\n1.5.0\n2.0.0" | ./semver filter "^1.0.0"
```

Output:
```
1.0.0
1.1.0
1.5.0
```

### latest

Find the highest version from stdin:

```bash
echo -e "1.0.0\n3.5.2\n2.1.0\n10.0.0" | ./semver latest
```

Output:
```
10.0.0
```

## Piping Commands

Commands can be piped together for powerful version operations:

```bash
# Find the latest version matching a constraint
echo -e "1.0.0\n1.1.0\n1.5.0\n2.0.0" | ./semver filter "^1.0.0" | ./semver latest
```

Output:
```
1.5.0
```

## Testing

Run the test suite:

```bash
go test -v
```

## Dependencies

- [github.com/Masterminds/semver/v3](https://github.com/Masterminds/semver) - Semantic versioning library
- [github.com/urfave/cli/v3](https://github.com/urfave/cli) - CLI framework
- [github.com/stretchr/testify](https://github.com/stretchr/testify) - Testing utilities

## License

See LICENSE file for details.
