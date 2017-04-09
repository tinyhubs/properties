# properties for golang

`*.properties`文件是java里面很常见的配置文件。这里是一个go语言版的*.properties文件读取库。

## go-properties文件格式定义

为了使得properties文件的识别更加简单快速，go的properties的文件格式和java的properties文件并不是等价的。它将java里面一些很少用到的格式特性都去掉了。

golang版本的properties文件的格式定义如下：

- 一行如果第一个非空白字符是`#`或者`!`，那么整行将被识别为注释行，注释行将被计息器忽略。
- 每个配置项都是单行的key-value对，不支持跨行，key和value以`=`分隔。
- key和value都是区分大小写；

    比如，下面其实是三个不同的配置项：
    ```
    SizeRange=1-20
    sizerange=1-20
    SIZERANGE=1-20
    ```
    
- key中不允许出现`=`，但value部分可以出现`=`；

    比如，下面这个配置项key为`expr`,而value是`A-B=C`:
    ```
    expr=A-B=C
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

####全局函数：

`properties.Load` 生成一个`properties.Properties`对象。

####`properties.Properties`的成员

`Pairs` 这是一个`map[string][string]`类型的成员。当我们需要注入一些外部的配置项时，可以直接使用该来。

#### `properties.Properties`的成员函数：

- `String` 按指定key返回对应的value，如果key不存在，返回`""`(空字符串)。
- `Int` 同与`String`类似，只是返回值是`int64`类型的且缺省值是`0`。`Int`只支持10进制的数据读取。
- `Float` 同与`String`类似，只是返回值是`float64`类型的且缺省值是`0`。
- `Bool` 同与`String`类似，只是返回值是`bool`类型的且缺省值是`false`。`Bool`函数会将`1`, `t`, `T`, `true`, `TRUE`, `True`识别为`true`，将`0`, `f`, `F`, `false`, `FALSE`, `False`识别为`false`。
- `Object` 这个函数提供了一个数据映射能力，可以将找到的value映射为任何类型。
- `StringDefault`，`IntDefault`、`FloatDefault`、`BoolDefault`、`ObjectDefault` 这几个函数的返回值和前面不带`Default`后缀的函数的行为类似，只是当配置项不存在时或者数据格式错误时，会直接返回参数中的`def`(缺省值)。
- `Get` 第一个参数返回的是key所对应的value；当配置项不存在时，第二个参数为false。我们可以借助这个函数来判断指定key的配置项是否存在（直接访问`Pairs`也可以判断是否存在，只是不推荐大家这样做）。




