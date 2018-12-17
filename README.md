
### This utility is designed to automatically install and update all imported packages in specified files.



**Installing the utility in Linux**


```yaml
wget -O - https://github.com/naviarh/goget/install.sh | bash
```



**How to use**


Update imported packages for all files in the current folder:

```yaml
goget .
```

Updating imported packages for specified files:

```yaml
goget file1.go ~/file2.go
```
 

