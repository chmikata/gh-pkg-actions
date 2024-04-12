# gh-pkg-actions - Manipulate GitHub Packages for GitHub Actions

## Introduction

This action refers to GitHub Packages.
You can get a list of packages and the tags they have been assigned.

> [!WARNING]
> Not all REST APIs are supported.

## Parameter Reference

### inputs

| Name      | Type     | Required | Default | Description                                                            |
| --------- | -------- | -------- | ------- | ---------------------------------------------------------------------- |
| `command` | `String` | `true`   |         | Search the list of packages by `package` and the list of tags by `tag` |
| `org`     | `String` | `true`   |         | Name of the organization to be searched                                |
| `token`   | `String` | `true`   |         | Specify `Personal Access Token`                                        |
| `matcher` | `String` | `false`  | `'.*'`  | Specify a regular expression to search package names                   |
| `pattern` | `String` | `false`  | `sem`   | Search semantic version in `sem`, Git commit hash in sha in `sha`      |
| `depth`   | `String` | `false`  | `0`     | Depth to search for tags, if 0, search all                             |

### outputs

| Name     | Type   | Description      |
| -------- | ------ | ---------------- |
| `result` | `Json` | List of Subjects |

## Usage

### package command

```yaml
name: ci

on:
  push:
    branches:
      - 'main'

jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - name: Package List
        id: list
        uses: chmikata/gh-pkg-actions@v1
        with:
          command: package
          org: organization
          token: ${{ secrets.PAT }}
          matcher: test-image
      - name: Output List
        run: |
          echo ${{ steps.list.outputs.result }}
```
The following output results
```bash
[{"id":1234567,"name":"test/package1"},{"id":2345678,"name":"test/package2"}]
```

### tag command

```yaml
name: ci

on:
  push:
    branches:
      - 'main'

jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - name: Package List
        id: list
        uses: chmikata/gh-pkg-actions@v1
        with:
          command: tag
          org: organization
          token: ${{ secrets.PAT }}
          matcher: test-image
          pattern: sem
          depth: 2
      - name: Output List
        run: |
          echo ${{ steps.list.outputs.result }}
```
The following output results
```bash
[{"id":1234567,"name":"test/package1","tags":["1.1.0-rc2","1.1.0-rc1"]},{"id":2345678,"name":"test/package2","tags":["1.2.0","1.1.0"]}]
```
