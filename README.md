# seimei

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
[![test](https://github.com/glassmonkey/seimei/workflows/test/badge.svg)](https://github.com/glassmonkey/seimei/actions?query=workflow%3Atest)
[![reviewdog](https://github.com/glassmonkey/seimei/workflows/reviewdog/badge.svg)](https://github.com/glassmonkey/seimei/actions?query=workflow%3Areviewdog)

**seimei** is a tool for dividing the Japanese full name into a last name and a fist name in Python to Go.
SEIMEI is a Go port of a tool ([namedivider-python](https://github.com/rskmoi/namedivider-python)) created in python to split Japanese first and last names.  

For implementation details, please check ([namedivider-python](https://github.com/rskmoi/namedivider-python)) from which you are porting.


# Installation

```bash
go install github.com/glassmonkey/seimei/cmd/seimei@latest
```

# Usage

## Options

```bash
$ seimei -h
  -name string
        separate full name(ex. 田中太郎)
  -parse string
        separate characters (default " ")
```

## Example

```bash
$ seimei -name 竈門炭治郎
竈門 炭治郎

$ seimei -name 竈門禰豆子 -parse @
竈門@禰豆子
```

# Licence
[Mit](LICENSE)

# Author
glassmonkey([@glassmonekey](https://twitter.com/glassmonekey))

