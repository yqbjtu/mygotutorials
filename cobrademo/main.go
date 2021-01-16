package main

import (
	goflag "flag"
	"fmt"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"k8s.io/klog"
	"os"
	"yq.c.com/cobrademo/cmd"
)

var name string
var age int

/*main.exe subcmd1  -n=hello -a=35
  subcmd1 --name=tom  --log_file=c:\\f\\klogdemo.log  --alsologtostderr=true
  go run main.go  subcmd1 --name=tom --log_file=c:\\f\\klogdemo.log  -n=abc --alsologtostderr=false  --logtostderr=false
*/
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "A cobra demo",
	Long:  `very simple demo case`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("main cmd")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		klog.Fatalf("root cmd execute failed, err=%v", err)
		os.Exit(-1)
	}
}

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.AddCommand(cmd.Runcmd)

	klog.InitFlags(nil)
	goflag.Parse()
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

func main() {
	//创建并执行
	Execute()
	klog.Info("exit main")
	klog.Flush()
}
