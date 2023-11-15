/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/patrikcze/go-veeam-cli/packages/time"
	"github.com/spf13/cobra"
)

// servertimeCmd represents the servertime command
var servertimeCmd = &cobra.Command{
	Use:   "servertime",
	Short: "This command allows you to get current date and time on the backup server.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiservername, _ := cmd.Flags().GetString("servername")
		apiport, _ := cmd.Flags().GetInt("port")
		stime := veeamtime.GetServerTime(apiservername, apiport)
		fmt.Println(stime)
	},
}

func init() {
	rootCmd.AddCommand(servertimeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	servertimeCmd.PersistentFlags().String("servername", "", "Veeam B&R RestAPI Server name.")
	servertimeCmd.PersistentFlags().Int("port",9419, "Veeam V&R RestAPI Server Port. (Default: 9419)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// servertimeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
