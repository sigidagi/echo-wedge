package cmd

import (
	"bytes"
	"echo-wedge/backend/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	//"os"
)

var cfgFile string
var version string

var rootCmd = &cobra.Command{
	Use:   "echo-wedge",
	Short: "Rest service",
	Long: `Rest service with wedge gateway connectivity is providing frontend service for
	> creating Wapp applications. 
	> documentation and support: https://www.seluxit.com
	> source and copyrigth information: https://www.seluxit.com`,
	RunE: run,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "path to configuration file (optional)")
	rootCmd.PersistentFlags().Int("log-level", 4, "debug=5, info=4, error=2, fatal=1, panic=0")
	// bind flag to config vars
	viper.BindPFlag("general.log_level", rootCmd.PersistentFlags().Lookup("log-level"))
	// defaults
	viper.SetDefault("gateway.url", "localhost:8051")
	viper.SetDefault("rest.url", "0.0.0.0:8000")
	viper.SetDefault("rest.bind", "0.0.0.0:8060")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
}

func Execute(v string) {
	version = v
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		b, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			log.WithError(err).WithField("config", cfgFile).Fatal("error loading config file")
		}
		viper.SetConfigType("toml")
		if err := viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
			log.WithError(err).WithField("config", cfgFile).Fatal("error loading config file")
		}
	} else {
		viper.SetConfigName("echo-wedge")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config/echo-wedge")
		viper.AddConfigPath("/etc/echo-wedge")
		if err := viper.ReadInConfig(); err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				log.Warning("No configuration file found, using defaults.")
			default:
				log.WithError(err).Fatal("read configuration file error")
			}
		}
	}

	if err := viper.Unmarshal(&config.C); err != nil {
		log.WithError(err).Fatal("unmarshal config error")
	}
}
