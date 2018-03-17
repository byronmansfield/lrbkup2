// Copyright © 2018 Byron Mansfield <byron@byronmansfield.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/byronmansfield/lrbkup/lrbkup/lte"
	"github.com/manifoldco/promptui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

type option struct {
	Number        int
	Description   string
	performAction func(int, string)
}

var options = []option{
	{
		Number:        1,
		Description:   "Backup my local photos to an external drive (not ready yet)",
		performAction: lte.BackupToExt,
	},
	{
		Number:        2,
		Description:   "Backup my external drive to S3 (not ready yet)",
		performAction: lte.BackupToExt,
	},
	{
		Number:        3,
		Description:   "Clear local photos from /Users/bmansfield/Pictures/Lightroom (not ready yet)",
		performAction: lte.BackupToExt,
	},
	{
		Number:        4,
		Description:   "Sync my local from an external drive (not ready yet)",
		performAction: lte.BackupToExt,
	},
	{
		Number:        5,
		Description:   "Pull down photos from S3 (not ready yet)",
		performAction: lte.BackupToExt,
	},
}

var templates = &promptui.SelectTemplates{
	Label:    "{{ . }}",
	Active:   "\U0001F336 {{ .Number | cyan }}) {{ .Description | red }}",
	Inactive: "  {{ .Number | cyan }}) {{ .Description | red }}",
	Selected: "\U0001F336 {{ .Number | red | cyan }}",
	Details: `
--------- Option ----------
{{ .Number }}) {{ .Description }}`,
}

var mainPrompt = promptui.Select{
	Label:     "What sort of Lightroom backup or sync would you like to perform?",
	Items:     options,
	Templates: templates,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lrbkup",
	Short: "Lightroom Backup utility",
	Long:  "A utility to assist with managing the storage of your Lightroom photos",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("		Lightroom Backup")
		fmt.Println("")

		i, _, err := mainPrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		options[i].performAction(i+1, options[i].Description)

	},
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lrbkup.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("greeting", "g", false, "Greeting")

	// var myHelpText string
	// myHelpText = "Prints help text!"
	// cobra.SetHelpTemplate(s string)
	// rootCmd.SetHelpTemplate(myHelpText string)
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

		// Search config in home directory with name ".lrbkup" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".lrbkup")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
