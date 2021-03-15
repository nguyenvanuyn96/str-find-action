# Find String Action

[![GitHub Marketplace](https://img.shields.io/badge/Marketplace-Find%20String-blue.svg?colorA=24292e&colorB=0366d6&style=flat&longCache=true&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAYAAAAfSC3RAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAM6wAADOsB5dZE0gAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAAERSURBVCiRhZG/SsMxFEZPfsVJ61jbxaF0cRQRcRJ9hlYn30IHN/+9iquDCOIsblIrOjqKgy5aKoJQj4O3EEtbPwhJbr6Te28CmdSKeqzeqr0YbfVIrTBKakvtOl5dtTkK+v4HfA9PEyBFCY9AGVgCBLaBp1jPAyfAJ/AAdIEG0dNAiyP7+K1qIfMdonZic6+WJoBJvQlvuwDqcXadUuqPA1NKAlexbRTAIMvMOCjTbMwl1LtI/6KWJ5Q6rT6Ht1MA58AX8Apcqqt5r2qhrgAXQC3CZ6i1+KMd9TRu3MvA3aH/fFPnBodb6oe6HM8+lYHrGdRXW8M9bMZtPXUji69lmf5Cmamq7quNLFZXD9Rq7v0Bpc1o/tp0fisAAAAASUVORK5CYII=)](https://github.com/nguyenvanuyn96/str-find-action)
[![Actions Status](https://github.com/nguyenvanuyn96/str-find-action/workflows/Build/badge.svg)](https://github.com/nguyenvanuyn96/str-find-action/actions)
[![Actions Status](https://github.com/nguyenvanuyn96/str-find-action/workflows/Integration%20Test/badge.svg)](https://github.com/nguyenvanuyn96/str-find-action/actions)

This action will find strings in your project files.

## Usage

### Example workflow

This example find `hello` string in all of your project files.

```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Find String
        uses: nguyenvanuyn96/str-find-action@master
        with:
          find: "hello"
```

### Inputs

| Input                  | Description                                                                              |
| ---------------------- | ---------------------------------------------------------------------------------------- |
| `find`                 | A string to find in your project files. _(This can be a [regular expression](https://github.com/google/re2/wiki/Syntax))_ |
| `include` _(optional)_ | A regular expression of files to include. _Defaults to `.*`._                            |
| `exclude` _(optional)_ | A regular expression of files to exclude. _Defaults to `.git/`._                         |

### Outputs

| Output          | Description                                 |
| --------------- | ------------------------------------------- |
| `fileFoundCount` | The number of files that have been found |
| `resultArray`    | The list of string which have been found |

## Examples

### Including a subdirectory

You can limit your find to a directory.

```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Find String
        uses: nguyenvanuyn96/str-find-action@master
        with:
          find: "hello"
          include: "justthisdirectory/"
```

### Filter by file name

You can limit your find string to just files with a specific name.

```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Find String
        uses: nguyenvanuyn96/str-find-action@master
        with:
          find: "hello"
          include: "README.md" # Will match all README.md files in any nested directory
```

### Exclude by file type

You can set your find string to ignore certain file types.

```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Find String
        uses: nguyenvanuyn96/str-find-action@master
        with:
          find: "hello"
          exclude: "*.py" # Do not modify Python files
```

## Publishing

To publish a new version of this Action we need to update the Docker image tag in `action.yml` and also create a new release on GitHub.

- Work out the next tag version number.
- Update the Docker image in `action.yml`.
- Create a new release on GitHub with the same tag.
