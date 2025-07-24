package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"path/filepath"
	"os"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocv",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
				examples and usage of using your application. For example:

				Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Cobra init
func initConfig() {
	home, _ := os.UserHomeDir()
	viper.AddConfigPath(filepath.Join(home, ".config/gocv"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gocv.yaml)")
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


