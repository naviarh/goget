
### This utility is designed to automatically install and update all imported packages in specified files.



**Installing the utility in Linux:**


```yaml
go get github.com/naviarh/goget
cd $GOPATH/src/github.com/naviarh/goget
go build
strip -sxv goget
cp ./goget $GOPATH/bin
```



**How to use:**


Update imported packages for all files in the current folder:

```yaml
goget .
```

Updating imported packages for specified files:

```yaml
goget file1.go ~/file2.go
```
 

