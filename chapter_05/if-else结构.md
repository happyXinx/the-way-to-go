if-else 结构

```go
if val:=10; val>max {
	// do something
}
```
但要注意的是，使用简短方式:= 声明的变量的作用域只存在于if结构中，
如果变量在if结构之前已存在，那么在if结构中，该变量原来的值会被隐藏。最简单的方案就是不要在初始化语句中声明变量

```go
if err := file.Chmod(0664); err != nil {
    fmt.Println(err)
    return err
}
```

```go
if value, ok := readData(); ok {
…
}
```


