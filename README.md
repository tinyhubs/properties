# Overview

[![Build Status](https://travis-ci.org/tinyhubs/properties.svg?branch=master)](https://travis-ci.org/tinyhubs/properties)
[![GoDoc](https://godoc.org/github.com/tinyhubs/properties?status.svg)](https://godoc.org/github.com/tinyhubs/properties)
[![Language](https://img.shields.io/badge/language-go-lightgrey.svg)](https://github.com/tinyhubs/properties)
[![License](https://img.shields.io/badge/license-New%20BSD-yellow.svg?style=flat)](LICENSE)
[![codecov](https://codecov.io/gh/tinyhubs/properties/branch/master/graph/badge.svg)](https://codecov.io/gh/tinyhubs/properties)
[![goreport](https://www.goreportcard.com/badge/github.com/tinyhubs/properties)](https://www.goreportcard.com/report/github.com/tinyhubs/properties)

`*.properties`文件是java里面很常见的配置文件。这里是一个go语言版的*.properties文件读处理库。本库支持properties文件的读取、修改、回写操作。也支持向properties文件中的属性追加、删除注释操作。

## go-properties文件格式定义

为了使得properties文件的识别更加简单快速，go的properties的文件格式和java的properties文件并不是等价的。它将java里面一些很少用到的格式特性都去掉了。

golang版本的properties文件的格式定义如下：

- 一行如果第一个非空白字符是`#`或者`!`，那么整行将被识别为注释行，注释行将被解析器忽略。

    比如，下面三行都会被解析器忽略(第三行是空白行)：
    ```
    #  这是注释行
    !  这也是注释行
        
    ```
    
- 每个配置项都是单行的key-value对，**不支持跨行**，key和value以`=`或者`:`分隔。

    比如，下面其实是三个配置项----`UserName=Fabirc \`、`Boch=`、`Contry=US`：
    ```
    UserName = Fabirc \
    Boch
    Contry=US
    ```
    
- key和value都是区分大小写；

    比如，下面其实是三个不同的配置项：
    ```
    SizeRange=1-20
    sizerange=1-20
    SIZERANGE=1-20
    ```
    
- 一行的第一个`=`即使key与value的分隔字符，所以key中不会出现`=`，但value部分可以出现`=`；

    比如，下面这个第一个行的key为`expr`,value是`A-B=C`；而第二行的key是一个`""`(空字符串)，value是`Hello`。当然第二种情况并没有实际意义:
    ```
    expr=A-B=C
    =Hello
    ```
    
    
- key和value的前后的空白都将被忽略，但key和value中间的空白会被原样保留；

    比如，下面这三个配置项的值都是`1-20`：
    ```
      SizeRange-1=1-20
    SizeRange-2 =  1-20   
       SizeRange-3  =  1-20  
    ```
- properties文件只支持**`UTF-8`**字符集，所以value中可以直接输入中文，遇到中文字符不必像java那样使用`\uxxxx`转义，直接用中文字面文字即可；

    比如，下面这个配置项的key为`地址`，value为`深圳`。
    ```
      地址=深圳
    ```
    
- 当value为空时，等号可忽略；

    所以，下面三个配置项的值都是空字符串(第一个等号后面有个)：
    ```
    Address-1=  
    Address-2=
    Address-3
    ```

## 接口定义

#### 属性文档

一个properties文档由一个`properties.PropertiesDocument`对象来表示。一个properties文档由多个key-value形式的属性组成。每个属性还可以追加一行或者多行注释。

#### 加载属性文档

`properties.Load` 从io流生成一个`properties.PropertiesDocument`对象。

```go
file, err := os.Open("test1.properties")
if nil != err {
    return
}

doc, err := properties.Load(file)
if nil != err {
    t.Error("加载失败")
    return
}

fmt.Println(doc.String("key"))
```


#### 创建一个新的属性文档对象

`properties.New` 直接创建一个新的属性文档对象,常用于属性创建文档文件的场景下。我们随后可以通过`properties.Save`函数将属性文档写入到文件或者输出流。

```go
doc := properties.New()
doc.Set("a", "aaa")
doc.Comment("a", "This is a comment for a")

buf := bytes.NewBufferString("")
properties.Save(doc, buf)
```


#### 将属性回写到文件或者流

`properties.Save` 可用于将一个文档回写到指定的writer中去。 

```go
buf := bytes.NewBufferString("")
properties.Save(doc, buf)
```

#### 属性值的读取

- **通用读取能力**

PropertiesDocument对象的`Get`方法提供了一个基本的元素读取能力：

```go
func (p PropertiesDocument) Get(key string) (value string, exist bool)
```

`Get`函数会返回两个参数，当对应的key在PropertiesDocument文档中存在时，会返回该key对应的value，且exist的值将为true；如果不存在，exist的值将是false。

我们经常利用`Get`来探测，某个指定key的属性是否在属性文件中定义了。

- **读取并转换**
读取属性然后转成对应的数据类型是个很常见的任务，所以PropertiesDocument为最常用的几种类型提供了方便的读取并转换的函数。  
  * `String()` 读取一个字符串型的属性，如果不存在默认返回`""`
  * `Int()` 读取一个属性并转换为`int64`类型，如果key对应的属性不存在，或者转换失败，返回值为0
  * `Uint()` 和`Int()`函数类似，只是返回的数据类型为uint64
  * `Float()` 也是和`Int()`函数类似，但返回值为float64
  * `Bool()` 同与`String`类似，只是返回值是`bool`类型的且缺省值是`false`。`Bool`函数会将`1`, `t`, `T`, `true`, `TRUE`, `True`识别为`true`，将`0`, `f`, `F`, `false`, `FALSE`, `False`识别为`false`。
  * `Object` 这个函数提供了一个数据映射能力，可以将找到的value映射为任何类型。

- **指定读取的缺省值**
前面的`String()`、`Int()`等函数在key不存在或者抓换失败的场景下，默认会返回零值。但零值往往不能满足我们的诉求，我们经常需要自己指定这些场景下的返回值。`StringDefault`，`IntDefault`、`FloatDefault`、`BoolDefault`、`ObjectDefault` 这几个函数的返回值和前面不带`Default`后缀的函数的行为类似，只是当配置项不存在时或者数据格式错误时，会直接返回参数中的`def`(缺省值)。


#### 属性的增删改

- **增加或者修改属性**

`Set()`函数用于修改指定的key的属性的值。如果指定的key的属性不存在，那么自动创建一个。

本库并没有提供按类型设置属性值的功能，前面描述的`Set()`函数只接受字符串类型的属性值作为输入。主要原因是数据的转换方式非常多，没有一种普适 的数据转换方法。所以，对于非字符串类型的值的设置需要自行转换成string类型。

```go
doc.Set("key", "value")
```

- **处理属性不存在的场景**

当属性不存在时，`Set()`函数会新建一个属性值，这种工作方式通常是很有用的。但是，有时候我们不希望`Set()`的这种自作聪明的行为。此时，我们可以通过`Get()`方法判断一些以确定是否需调用`Set()`。

```go
_, exist := doc.Get("key")
if !exist {
    return errors.New("Key is not exist")
}

doc.Set("key", "New-Value")
```

- **删除属性**

`Del()`函数用于删除指定key的属性。它会返回一个bool值，用于表示当前的key的属性是否存在。

```go
exist := doc.Del("key")
```


#### 操作注释

在本库中，注释是绑定到属性的。位于属性的key-value定义前面，且与属性之间没有空白行的多行注释，我们会判定这些注释是属于该属性的，比如：
```properties
 # Comment1
 # Comment2
 
 # Comment3
 # Comment4
 mykey=myvalue
```

上面的Comment3和Comment4是mykey属性的注释，但是Comment1和Comment2却不是。

PropertiesDocument的`Comment()`函数用于为属性指定一些注释。而`Uncomment()`函数用于删除指定的key的注释。

PropertiesDocument的`Comment()`函数允许一次性指定多行注释，而`Uncomment()`用于一次性删除一个指定的key的所有的注释。


#### 文档对象枚举

PropertiesDocument的`Accept()`和`Foreach()`函数都是用来对文档对象进行枚举的，但是`Foreach()`专用于对属性进行遍历。而`Accept()`可以通过对属性和注释进行遍历。

实际上，`Save()`函数就是利用`Accept()`函数来实现的：

```go
func Save(doc *PropertiesDocument, writer io.Writer) {
    doc.Accept(func(typo byte, value string, key string) bool {
        switch typo {
        case '#', '!', ' ':
            {
                fmt.Fprintln(writer, value)
            }
        case '=', ':':
            {
                fmt.Fprintf(writer, "%s%c%s\n", key, typo, value)
            }
        }
        return true
    })
}
```

`Accept()`函数的回调函数里面有个`typo`参数，这个参数决定了当前这条记录是注释还是一个有效的属性。`typo`的取值可能有下面一些：
- `'#'` 表示当前的value是`#`开头的注释
- `'!'` 表示当前的value是`!`开头的注释
- `' '` 表示当前的value是个空行或者空白行
- `'='` 表示当前的value是个以`=`分隔的属性
- `':'` 表示当前的value是个以`:`分隔的属性



