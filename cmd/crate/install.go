/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package crate

import (
	"fmt"
	"log"

	"github.com/justjordant/cargo/internals"
	"github.com/spf13/cobra"
)

var (
	appName string = "tool.yml"
)

// TODO - in this cmd this will need to
// 1. parse the user/org repositories for crates `my-application-crate`
// 2. download yaml temp for getting the download URL of the application
// 3. download the application.
// 4. install the application.

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")

		//  fmt.Println(internals.FormatFileName(appName))

		filePath, err := internals.DownloadCargoYaml(args[0])
		if err != nil {
			log.Fatal("Error: ", err)
		}
		fmt.Println(filePath)
	},
}

func init() {
	CrateCmd.AddCommand(installCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
