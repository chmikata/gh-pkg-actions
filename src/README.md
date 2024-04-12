# gh-pkg-cli - CLI for GitHub Packages

## Introduction

CLI tool to manipulate the REST API of github packages.

> [!WARNING]
> Not all REST APIs are supported.

## Parameter Reference

### Command Usage

```bash
Usage:
  gh-pkg-cli [command] [flags]

Available Commands:
  package     Display package
  tag         Display container image tags

Flags:
  -h, --help   help for package
  ・・・

Global Flags:
  -m, --matcher string   Name of the container image to match (default ".*")
  -o, --org string       Organization name
  -t, --token string     Token for authentication
```

### Global Flags

| Name      | Shortened Name | Type     | Required | Default | Description                          |
| --------- | -------------- | -------- | -------- | ------- | ------------------------------------ |
| `org`     | `o`            | `String` | `true`   | `''`    | Organization name                    |
| `token`   | `t`            | `String` | `true`   | `''`    | PAT on GitHub                        |
| `matcher` | `m`            | `String` | `false`  | `.*`    | Name of the container image to match |

### Command [tag] Flags

| Name      | Shortened Name | Type     | Required | Default | Description                                   |
| --------- | -------------- | -------- | -------- | ------- | --------------------------------------------- |
| `pattern` | `p`            | `String` | `true`   | `''`    | Pattern to <sem> or <sha> match image to tags |
| `depth`   | `d`            | `Int`    | `false`  | `0`     | Depth of tags to display                      |


## Command Output

### Command [package]

```bash
$ gh-pkg-cli package --org org --token ********** --matcher test/
{"id":1234567,"name":"test/package1"}
{"id":2345678,"name":"test/package2"}
```

### Command [tag]

```bash
$ gh-pkg-cli tag --org org --token ********** --matcher test/ --pattern sem --depth 2
{"id":1234567,"name":"test/package1","tags":["1.1.0-rc2","1.1.0-rc1"]}
{"id":2345678,"name":"test/package2","tags":["1.2.0","1.1.0"]}
```
