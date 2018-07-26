###变量

> 注释为 php 写法

```go
// $name = 'php';
var name string = 'go'
var name = 'go'
name := 'go'
```

#### 交换两个变量的值

```go
// $temp = $a; $a = $b; $b = $temp;
a,b = b,a
```

### 循环
```go
// foreach($lists as $key => $value) {}
for key,value := rang lists {}

// for($i = 0; $i < count($lists); $i++) {}
for  i := 0; i < len(lists); i++ {}

// while($i < count($lists)) {}
for i < len(lists) {}

// do {} while (true)
for {}

```

### 类，对象
> go 没有面向对象语言中的类或对象，但其具有与集成代码和行为的数据结构的相同定义想匹配的类型，被称为结构体。

```go
type rect struct {
    Width int
    Height int
}

fun (r *rect) area() int { //在结构体上添加方法
    return r.Width * r.Height
}

r := rect{Wight: 10, Height: 9} //初始化

fmt.Print(r.area())
```
> 定义一个复杂的结构体

```go
type Employee struct {
    Name string
    Job Job
}

type Job struct {
    Employee string
    Position string
}

//去结构化
e := Employee{
    Name: "satoshi",
    Job: {
        Employee: "no body",
        Position: "god",
    }
}
```

