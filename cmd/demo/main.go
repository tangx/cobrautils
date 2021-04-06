package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tangx/cobrax"
)

type student struct {
	Name   string `name:"name" value:"zhangsan" usage:"student name" persistent:"true"`
	Age    int64  `name:"age" value:"20" usage:"student age" shorthand:"a"`
	Gender bool   `value:"true"`
}

var rootCmd = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		fmt.Println(stu)
	},
}

var stu = student{
	Name: "zhangsanfeng",
	Age:  20100,
}

func main() {

	cobrax.BindFlags(rootCmd, &stu)

	_ = rootCmd.Execute()

}
