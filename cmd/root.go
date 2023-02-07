/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "beeboop",
	Short: "A brief description of your application",
	Long: `Meet BeeBoop, the CLI with a BZZZ-tastic personality!
This little guy may be small in size, but he's got big aspirations.
BeeBoop dreams of one day transforming from a simple code machine into a real live human boy.
Can you imagine a CLI with a love for honey, flowers, and adventure?
It's like a scene straight out of a Disney movie! With his determination and can-do spirit, BeeBoop is on a mission to make his wildest dreams come true.
So buzz on over and join in on the fun with BeeBoop, the CLI with a heart of honey and a smile as sweet as nectar. Just dont let'm use knives :)`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.beeboop.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
