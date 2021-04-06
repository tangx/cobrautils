package cobrax

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func BindFlags(cmd *cobra.Command, opts interface{}) {

	rvPtr := reflect.ValueOf(opts)

	// 不是指针不能进行操作
	// Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。
	if rvPtr.Kind() != reflect.Ptr && rvPtr.Elem().Kind() != reflect.Struct {
		fmt.Printf("want a Struct Ptr, got %#T \n", rvPtr.Type())
		return
	}

	// 获取 opts 结构体实例对象的反射
	// Indirect: 持有的指针指向的值的Value
	rv := reflect.Indirect(rvPtr)

	// fmt.Println(rv.Type()) // (ex) student :  具体的 结构体名字
	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		/*
			var stu = student{
				Name: "zhangsan",
				Age:  20,
			}
		*/
		// typField : 结构体字段本身的属性， 与结构体实例化无关 (ex. stu.Name)
		typField := typ.Field(i)
		// valueField : 结构体实例化后字段对应的值的属性。 (ex. stu.Name -> zhangsan)
		valueField := rv.Field(i)

		// 2. 获取 name, shorthand。
		// 2.1. 如果 name 不存在， 则为字段本身名称
		name := typField.Tag.Get("name")
		if len(name) == 0 {
			name = strings.ToLower(typField.Name)
		}
		// 2.2. 获取
		shorthand := typField.Tag.Get("shorthand")

		// 3. 获取 usage
		usage := typField.Tag.Get("usage")

		// 4. 初始化 flags 变量
		flags := cmd.Flags()

		// 4.1. 是否为 Persistent flags
		if val, ok := typField.Tag.Lookup("persistent"); ok && val == "true" {
			fmt.Println("val=", val)
			flags = cmd.PersistentFlags()
		}

		// 6. get default value
		value := typField.Tag.Get("value")

		// 5. 判断 kind 类型
		switch typ.Field(i).Type.Kind() {
		case reflect.String:
			// 1.1 done : Addr() 获取值的内存地址， Interface() 并以 interface 类型返回， (*string) 并进行 类型指针类型 断言
			valuePtr := valueField.Addr().Interface().(*string)
			// 1.2 done : 将 reflect.Type 值转换为对应的值
			// value := valueField.String()
			// 1.3 done: 设置 flag
			flags.StringVarP(valuePtr, name, shorthand, value, usage)

		case reflect.Int64:
			flags.Int64VarP(valueField.Addr().Interface().(*int64), name, shorthand, mustConvInt64(value), usage)

		case reflect.Bool:
			flags.BoolVarP(valueField.Addr().Interface().(*bool), name, shorthand, mustConvBool(value), usage)

		case reflect.Slice:
			flags.StringSliceVarP(valueField.Addr().Interface().(*[]string), name, shorthand, mustConvStringSlice(value), "")
		}
	}

}

func mustConvInt64(str string) int64 {
	if len(str) == 0 {
		return 0
	}

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func mustConvBool(str string) bool {
	if len(str) == 0 {
		return false
	}
	ok, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return ok
}

func mustConvStringSlice(str string) []string {
	if len(str) == 0 {
		return []string{}
	}
	return strings.Split(str, ",")
}
