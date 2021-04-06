# cobrax 使用反射获取 flag 配置

## Usage

```go
type student struct {
	Name   string `name:"name" value:"zhangsan" usage:"student name" persistent:"true"`
	Age    int64  `name:"age" value:"20" usage:"student age" shorthand:"a"`
	Gender bool   `value:"true"`
	Address []string `value:"addr1,addr2,addr3"`
}

var stu = student{}
var cmd = *cobra.Command{}

func init(){
	cobrax.BindFlags(cmd,&stu)
}
```

> Attention: 数值的类型的必须是 `int64`。 因为不想做兼容。


## QA

### `kind` and `type`

相较于 Type 而言，Kind 所表示的范畴更大。类似于家用电器（Kind）和电视机（Type）之间的对应关系。或者电视机（Kind）和 42 寸彩色电视机（Type）

Type 是类型。Kind 是类别。Type 和 Kind 可能相同，也可能不同。通常基础数据类型的 Type 和 Kind 相同，自定义数据类型则不同。

对于反射中的 kind 我们既可以通过 reflect.Type 来获取，也可以通过 reflect.Value 来获取。他们得到的值和类型均是相同的。



