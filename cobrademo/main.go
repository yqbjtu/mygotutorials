package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var name string
var age int

//main.exe  testCobra -n=hello -a=35
var RootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "A cobra demo",
	Long:  `very simple demo case`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(name) == 0 {
			fmt.Println("no name")
			return
		}
		Show(name, age)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	//分配标志flag
	RootCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	RootCmd.Flags().IntVarP(&age, "age", "a", 18, "person's age")
}
func main() {
	//创建并执行
	Execute()
}
func Show(name string, age int) {
	//命令执行过程
	fmt.Printf("My Name is %s, My age is %d\n", name, age)
}
