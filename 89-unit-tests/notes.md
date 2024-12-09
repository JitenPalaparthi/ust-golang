The files should be suffixed with _test.go

Every Test functionality should start with TestXXX(t *testing.T)

### To test the project

```
go test ./...
```
### To test only a file

```
go test -timeout 30s -run ^TestNew$ test-demo/shape/rect
```

### To test a package

```
go test ./shape/rect
```

### To test a file

```
go test ./shape/rect/rect_test.go 
```

### test cover

```
go test ./... -cover
go test ./shape/rect/ -cover  
```

### test cover with output file

```
go test ./... -coverprofile=coverage.out
go tool cover -html coverage.out
```