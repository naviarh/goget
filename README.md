
### This utility is designed to automatically install and update all imported packages in specified files.



**Install goget to linux:**


```sh
wget -O - https://raw.githubusercontent.com/naviarh/goget/master/install.sh | bash
```



**How to use**


Update imported packages for all files in the current folder:

```sh
goget .
```

Updating imported packages for specified files:

```sh
goget file1.go ~/file2.go
```
 

