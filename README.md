# gh-pkg-actions - Manipulate GitHub Packages for GitHub Actions

## Introduction

This action refers to GitHub Packages.
You can get a list of the tags they have been assigned.

> [!WARNING]
> Not all REST APIs are supported.

## Parameter Reference

### inputs

| Name      | Type     | Required | Default | Description                                                       |
| --------- | -------- | -------- | ------- | ----------------------------------------------------------------- |
| `command` | `String` | `true`   |         | Search the list of tags by `tag`                                  |
| `org`     | `String` | `true`   |         | Name of the organization to be searched                           |
| `token`   | `String` | `true`   |         | Specify `Personal Access Token`                                   |
| `matcher` | `String` | `true`   | `''`    | Specify a search package names                                    |
| `pattern` | `String` | `false`  | `sem`   | Search semantic version in `sem`, Git commit hash in sha in `sha` |
| `depth`   | `String` | `false`  | `0`     | Depth to search for tags, if 0, search all                        |
| `range`   | `String` | `false`  | `all`   | Specify the range of tags to compare by `major`, `minor`, `all`   |

### outputs

| Name     | Type   | Description      |
| -------- | ------ | ---------------- |
| `result` | `Json` | List of Subjects |

## Usage

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
          token: ${{ secrets.GITHUB_TOKEN }}
          matcher: test-image
          pattern: sem
          depth: 2
          check-range: minor
      - name: Output List
        run: |
          echo ${{ steps.list.outputs.result }}
```
The following output results
```bash
{"name":"test/package1","tags":["1.1.0-rc2","1.0.0"]}
```
