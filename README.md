#golang

## go run

```
go run main.go
```

## to build 

```
go build main.go
```

```
go build -o hello main.go
```

- to stripe down the size of the binary

```
go build -ldflags="-s -w" -o demo1 main.go
```

- list all os and architectures go can build 

```
go tool dist list
```
- go build , cross compile to different target

```
GOOS=windows GOARCH=amd64 go build -o build/demo_windows_amd
64.exe main.go

or 

GOOS=linux GOARCH=amd64 go build -o build/demo_linux_amd
64 main.go
```

- escape analysis

```
go run -gcflags="-m" main.go
```

## keywords 

- break,case,const,continue,default,else,fallthrough,for,func, if, import,map, package,range,switch,struct,type,var (18 out of 25)

## builtin

- append,cap,clear,complex,copy,delete,len,imag,make,max,min,new,print,println,real (15  out of 18)

## Why Golang?

- 1. Fast compilation
- 2. No dependencies
- 3. No runtime has to installed on the target machine
- 4. static binaries
- 5. statically linked
- 6. fast startup times
- 7. statically typed, general purpose and systems programming language
- 8. Low latency GC (Comparatively)
- 9. Concurrency (CSP)
- 10. Performance

## Compilation

## GO Environment variables

- GOROOT : Where go is installed
- GOPATH : where all custom or third party packages are downloaded and compiled
- GOBIN  : where all 'go install' are installed into this directory.If empty, $GOPATH/bin directory is used. 
- GOARCH : processor architecture amd64 or arm64 or wasm or 386 etc.
- GOOS   : operating system linux or windows or darwin or android etc.


Compilation ->

