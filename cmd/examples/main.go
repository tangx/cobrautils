package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tangx/cobrautils"
)

type student struct {
	Name    string `flag:"name" usage:"student name" persistent:"true"`
	Age     int64  `flag:"age" usage:"student age" shorthand:"a"`
	Gender  bool
	Address address `flag:"addr"`
}

type address struct {
	Home   string `flag:"home"`
	School string `flag:"-"`
}

var rootCmd = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		fmt.Println(stu)
	},
}

var stu = student{
	Name:   "zhangsanfeng",
	Age:    20100,
	Gender: false,
	Address: address{
		Home:   "chengdu",
		School: "shuangliu",
	},
}

func main() {

	cobrautils.BindFlags(rootCmd, &stu)

	_ = rootCmd.Execute()

	/*
	   go run . --addr.home sichuan
	   Usage:
	     root [flags]

	   Flags:
	         --addr.home string    (default "chengdu")
	     -a, --age int            student age (default 20100)
	     -h, --help               help for root
	         --name string        student name (default "zhangsanfeng")
	   {zhangsanfeng 20100 false {sichuan shuangliu}}
	*/
}
