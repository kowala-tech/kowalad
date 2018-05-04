package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/kowala-tech/kowalad"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// flags
var (
	cfgFile string
)

var (
	config kowalad.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kowalad",
	Short: "A light node that simplifies communication with Kowala's blockchains",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listenAndServe(cmd, args); err != nil {
			os.Exit(1)
		}
	},
}

func listenAndServe(cmd *cobra.Command, args []string) error {
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	backend := kowalad.NewBackend()
	if err := backend.StartNode(&config); err != nil {
		return err
	}

	// @NOTE (rgeraldes( doneCh is not being used for now - used for auxiliary services
	_ = handleInterruption(backend)

	kowalaNode := backend.Node().KowalaNode()
	kowalaNode.Wait()

	return nil
}

func handleInterruption(backend kowalad.Backend) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt)
		defer signal.Stop(signalCh)
		<-signalCh
		close(doneCh)
		if err := backend.Node().Stop(); err != nil {
			// @TODO (rgeraldes) - log message
			os.Exit(1)
		}
	}()

	return doneCh
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kowalad.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kowalad" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kowalad")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
