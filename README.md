# GoBat2Exe

![GitHub release (latest by date)](https://img.shields.io/github/v/release/jamesmoriarty/gobat2exe) ![Continuous Integration](https://github.com/jamesmoriarty/gobat2exe/workflows/Continuous%20Integration/badge.svg)

Convert Windows batch files into executable files.

## Usage

You can drag and drop `.bat` files onto `bat2exe.exe` or invoke via command line:

```
.\bat2exe.exe .\example.bat
```

## Download

You can download releases [here](https://github.com/jamesmoriarty/gobat2exe/releases).

## Install

You can install with Golang with:

```
go get github.com/jamesmoriarty/gobat2exe
go install github.com/jamesmoriarty/gobat2exe
```

## Build

```
go build .\cmd\bat2exe.go
```

## Test

```
go test
```