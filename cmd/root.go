package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/uhhc/amf/pkg/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "amf",
	Short: "Amf",
	Long:  `Amf is a CLI tool for Another micro framework.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .config.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
