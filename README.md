# GoBat2Exe

![Continuous Integration](https://github.com/jamesmoriarty/gobat2exe/workflows/Continuous%20Integration/badge.svg)

Convert Windows batch files into executable files.

## Usage

You can drag and drop `.bat` files onto `bat2exe.exe` or invoke via command line:

```
.\bat2exe.exe .\example.bat
```

## Build

```
go build .\cmd\bat2exe.go
```

## Test

```
go test
```