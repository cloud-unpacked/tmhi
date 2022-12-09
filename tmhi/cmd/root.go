package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var cfgFile string
var versionFl bool

var rootCmd = &cobra.Command{
	Use:   "tmhi",
	Short: "A CLI to manage the T-Mobile Home Internet Router",
	Run: func(cmd *cobra.Command, args []string) {
		if versionFl {
			versionCmd.Run(cmd, []string{"--short"})
		} else {
			cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolVar(&versionFl, "version", false, "runs version --short")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	viper.SetConfigFile(os.ExpandEnv("$HOME/.config/tmhi/creds.yaml"))

	viper.ReadInConfig()
	viper.WriteConfig()
}
