
## 基于计数器的迭代 

```go
for 初始化语句；迭代语句；修饰语句{
	
}
``` 


## 基于条件判断的迭代

for结构的第二种形式是没有头部的条件判断迭代，类似于for循环。
```go
for i>=0 {
	i = i - 1
	fmt.Printf("123")
}
```


## 无限循环

`for {}`

无限循环的经典应用是服务器

```go
for t,err=p.Token(); err!=nil; t,err=p.Token(){
	
}
```

## for-range结构

类似于foreach语句，可以获得每次迭代所对应的索引。
`for ix, val := range coll {}`

注意val值是只读的，为集合对应索引的值拷贝，对它所做的任何修改都不会影响到集合原有的值。



## break 

* for结构，退出循环
* 对于select和switch， 一个break的作用是跳过整个代码块执行后面的代码。

## continue

忽略剩余的循环体而直接进入下一次循环的过程，而不是无条件的执行下一次循环，执行前仍需要判断循环条件。








