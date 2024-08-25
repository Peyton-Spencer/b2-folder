# b2-folder

## Install

    go install github.com/peyton-spencer/b2-folder@latest

## Usage

```
b2-folder upload -h
Using b2 version 4.1.0

b2 uploader

Usage:
  b2-folder upload [flags]

Flags:
  -b, --bucket string      bucket name
      --dry                dry run
  -f, --folder string      folder name
  -h, --help               help for upload
  -r, --replace string     replace string before/after
  -s, --skip stringArray   skip files that contain these strings
      --snake              convert upload filepath to snake_case
```
