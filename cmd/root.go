package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configPath string

var rootCMD = &cobra.Command{
	Use:   "url shortener",
	Short: "url shortener service!",
}

func init() {
	rootCMD.PersistentFlags().StringVarP(&configPath, "config-path", "c", "env.yaml", "path to config directory")

}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
