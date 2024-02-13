# windows-buildr

Simply build Windows Resources file for GO applications.

This will build a _*.syso_ file for your GO application to build a windows resource file that will be utilized during the native __go build__ process. 

# Build

``` bash

make build

```

## Command File

Once your build completes it will create a __bin__ folder that will contain a _VERSION_/macos/_ARCH_ folder and under that will be the _windows-buildr_ file that can be executed.

# Usage

## Arguments

- n
  - Name of the Software you're packaging
- v
  - Version of the Software "_v0.0.0_"
- i
  - File path to the Icon PNG to be displayed when viewing the application
- d
  - Human readable description of the software
- w
  - Working directory
- o
  - File name for the ouput resource file with the -w prepended
  - __DEFAULT__ rsrc_windows_amd64.syso
- c
  - Company name of software
- r
  - Copyright of Software

## Example of Usage

Below is an example of how to use this command line app for an automated build process with __make__.

``` bash
VERSION := v0.0.0
COMPANY := Your Company name
COPYRIGHT := Your Copyright Info Here

build: clean create-sys build-windows build-macos

create-sys:
	./cmd/windowsinfo -n "Software Name" -v "$(VERSION)" -c "$(COMPANY)" -r "$(COPYRIGHT)" -w "/Users/username/Developer/project-name" -d "description of software being packaged" -i "theme/icon.png" 

build-windows:
	mkdir -p dist/windows
	GOOS=windows GOARCH=amd64 go build -o dist/windows/XXXX_x64.exe

build-macos:
	mkdir -p dist/mac
	GOOS=darwin GOARCH=amd64 go build -ldflags "-w -X main.Version=$(VERSION) -X main.Company=$(COMPANY)" -o dist/mac/XXXX_x64
	GOOS=darwin GOARCH=arm64 go build -ldflags "-w -X main.Version=$(VERSION) -X main.Company=$(COMPANY)" -o dist/mac/XXXX_arm64

clean:
	rm -rf dist
	rm -f rsrc_windows_amd64.syso

```

Simply save your build file after changing the appropriate values and in the terminal just type _make build_ and now your building process is much more automated.

# Questions

Post any questions or comments to [me](mailto:braddschick@duck.com)
