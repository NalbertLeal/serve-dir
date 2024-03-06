# Serve dir

![GitHub](https://img.shields.io/github/license/NalbertLeal/serve-dir)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/NalbertLeal/serve-dir)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/NalbertLeal/serve-dir)

## Description

_Serve dir_ is a GO cli tool developed to serve all files and folders inside a directory using a HTTP server. Don't use this as your main option to serve files, the security level of this tool was imagined only on the level of someone that wish to serve (for any reason) files inside a directory using the HTTP protocol.

## How to Use

### Prerequisites

Ensure you have Go installed on your machine. For instructions on Go installation, visit [https://golang.org/doc/install](https://golang.org/doc/install).

### Installation

```bash
go get -u github.com/NalbertLeal/serve-dir
```

### Basic Usage

To serve a folder content just call _Serve Dir_ inside the foder

```bash
serve-dir
```

The executable don't accept a path as parameter, so the follow is not supported:

```bash
serve-dir ./some-subfolder
```

The program return a file or an list of the files inside the folder. If the
folder has an index.html file than return this file as default behavior. To
avoid this behavior use the flag **-no-index** and all the folders content are
returned. To never return the folders content (user receive an 404 when require
access to folder content) jst use the flag **-no-list**.

## Contribution

If you wish to contribute improvements, bug fixes, or new features, please feel free to open an issue or submit a pull request. Your contributions are highly appreciated!

## License

This project is licensed under the [GNU LESSER GENERAL PUBLIC LICENSE](https://choosealicense.com/licenses/lgpl-2.1/).