package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tangx/cobrax"
)

type student struct {
	Name    string `name:"name" usage:"student name" persistent:"true"`
	Age     int64  `name:"age" usage:"student age" shorthand:"a"`
	Gender  bool
	Address []string `name:"addr"`
}

var rootCmd = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		fmt.Println(stu)
	},
}

var stu = student{
	Name:    "zhangsanfeng",
	Age:     20100,
	Gender:  false,
	Address: []string{"addr1", "addr2"},
}

func main() {

	cobrax.BindFlags(rootCmd, &stu)

	_ = rootCmd.Execute()

}
