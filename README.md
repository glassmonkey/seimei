# seimei

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
[![test](https://github.com/glassmonkey/seimei/workflows/test/badge.svg)](https://github.com/glassmonkey/seimei/actions?query=workflow%3Atest)
[![reviewdog](https://github.com/glassmonkey/seimei/workflows/reviewdog/badge.svg)](https://github.com/glassmonkey/seimei/actions?query=workflow%3Areviewdog)

**seimei** is a Go port of a tool ([namedivider-python](https://github.com/rskmoi/namedivider-python)) created in python to split Japanese first and last names.  

For implementation details, please check ([namedivider-python](https://github.com/rskmoi/namedivider-python)) from which porting.


# Installation

```bash
go install github.com/glassmonkey/seimei/v2/cmd/seimei@latest
```

# Usage

## Options

```bash
$ seimei -h
Usage:
  seimei [flags]
  seimei [command]

Available Commands:
  name        It parse single full name.
  file        It bulk parse full name lit in the file.
  help        Help about any command

Flags:
  -h, --help      help for seimei
  -v, --version   version for seimei

Use "seimei [command] --help" for more information about a command

```

## Example

```bash
$ seimei name --name 竈門炭治郎
竈門 炭治郎

$ seimei name --name 竈門禰豆子 --parse @
竈門@禰豆子
```

```
$ cat /tmp/kimetsu.txt
竈門炭治郎
竈門禰豆子
我妻善逸
嘴平伊之助

$ seimei file --file /tmp/kimetsu.txt
竈門 炭治郎
竈門 禰豆子
我妻 善逸
嘴平 伊之助

$ seimei file --file /tmp/kimetsu.txt --parse @
竈門@炭治郎
竈門@禰豆子
我妻@善逸
嘴平@伊之助
```

# License
[Mit](LICENSE)

# Author
glassmonkey([@glassmonekey](https://twitter.com/glassmonekey))

