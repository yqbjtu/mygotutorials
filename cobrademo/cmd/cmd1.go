package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"yq.c.com/cobrademo/pkg/math"
)

var name string
var age int

var Runcmd = &cobra.Command{
	Use:   "subcmd1",
	Short: "run command",
	Long:  "Run command.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			klog.Info("no name")
			return
		}
		klog.Infof("cmd:%+v, cmd:%+v", cmd, args)
		Show(name, age)
	},
}

func init() {
	Runcmd.Flags().SortFlags = false
	Runcmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	Runcmd.Flags().IntVarP(&age, "age", "a", 18, "person's age")
}

func Show(name string, age int) {
	//命令执行过程
	after20y := math.Add(age, 20)
	fmt.Printf("My Name is %s, My age is %d. After 20y, age is %d\n", name, age, after20y)

}
