# b2-folder
- A folder uploader tool for [Backblaze](https://www.backblaze.com/)
- Utilizes the [b2 CLI](https://www.backblaze.com/docs/cloud-storage-upload-files-with-the-cli)

## Install

    go install github.com/peyton-spencer/b2-folder@latest

## Usage

```
b2-folder -h
Batch upload folders to Backblaze

Usage:
  b2-folder [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  upload      Upload a folder to a Backblaze bucket

Flags:
  -h, --help   help for b2-folder

Use "b2-folder [command] --help" for more information about a command.
```

```
b2-folder upload -h
Upload a folder to a Backblaze bucket

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
