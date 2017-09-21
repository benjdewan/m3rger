# m3rger
[![Build Status](https://travis-ci.org/benjdewan/m3rger.svg?branch=master)](https://travis-ci.org/benjdewan/m3rger)
[![Go Report Card](https://goreportcard.com/badge/github.com/benjdewan/m3rger)](https://goreportcard.com/report/github.com/benjdewan/m3rger)
[![codebeat badge](https://codebeat.co/badges/38252ff8-f3e5-4c64-826d-7f4489d3b6b8)](https://codebeat.co/projects/github-com-benjdewan-m3rger-master)
[![Docker Automated Build](https://img.shields.io/docker/automated/benjdewan/m3rger.svg)](https://hub.docker.com/r/benjdewan/m3rger/)

m3rger is a tool for merging up to 3 layers of YAML configuration files.

## Usage

```console
$ ./m3rger-darwin --help
usage: m3rger-darwin [<flags>] <default> [<low>] [<high>]

Flags:
      --help     Show context-sensitive help (also try --help-long and --help-man).
  -o, --out=OUT  Write the output to a file (by default m3rger prints to stdout)
      --version  Show application version.

Args:
  <default>  A YAML file with default values that may be overridden
  [<low>]    A YAML file that will override any values in 'default', but not any in the 'high' file
  [<high>]   A YAML file whose keys will override anything from 'low' or 'default'
```

`m3rger` merges up to three files.
*   `default`: A YAML file with default/generic values. If no other files are
     provided `default` is returned as is.
*   `low`: Scalar values and arrays in this file will overwrite those of
    `default` whenever there is a collision. Maps are merged.
*   `high`: Scalar values and arrays in this file will overwrite those of
    `low` and/or `default` whenever there is a collision. Maps are merged.

## Docker usage
The only trick with running `m3rger` in a container is remembering that the
files to be merged have to exist *inside* the container for `m3rger` to be
able to read them. The easiest way to do that is using `--volume` to mount
the directory the files are in:
```console
$ ls
default.yml high.yml    low.yml
$ docker run --rm --read-only --volume=$(pwd):/input:ro \
    benjdewan/m3rger /input/default.yml /input/low.yml /input/high.yml
foo: mumble
key: val
map:
  baz: quux
  key: val
```

## Examples

### 1. Just a default file

default.yml
```yaml
foo: bar
map:
  key: val
  baz: qux
```

```console
$ ./m3rger ./default.yml
foo: bar
map:
  key: val
  baz: qux
```

### 2. Two config files with no collision

default.yml
```yaml
foo: bar
map:
  key: val
  baz: qux
```

override.yml
```yaml
mumble: foobar
```

```console
$ ./m3rger ./default.yml ./override.yml
foo: bar
map:
  key: val
  baz: qux
mumble: foobar
```

### 3. Two config files with a scalar collision

default.yml
```yaml
foo: bar
map:
  key: val
  baz: qux
```

override.yml
```yaml
foo: mumble
```

```console
$ ./m3rger ./default.yml ./override.yml
foo: mumble
map:
  key: val
  baz: qux
```

### 4. Two conig files with a map merge
default.yml
```yaml
foo: bar
map:
  key: val
  baz: qux
```

override.yml
```yaml
map:
  key: mumble
  foo: bar
```

```console
$ ./m3rger ./default.yml ./override.yml
foo: bar
map:
  key: mumble
  baz: qux
  foo: bar
```

### 5. Three way merge
default.yml
```yaml
foo: bar
map:
  key: val
  baz: qux
```

low.yml
```yaml
foo: baz
key: val
map:
  baz: zab
```

high.yml
```yaml
foo: mumble
map:
  baz: quux
```

```console
$ ./m3rger ./default.yml ./low.yml ./high.yml
foo: mumble
key: val
map:
  baz: quux
  key: val
```
