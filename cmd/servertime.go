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
	Long: `Will try to get information about current time and timezone directly from remote server.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiservername, _ := cmd.Flags().GetString("servername")
		apiport, _ := cmd.Flags().GetInt("port")
		if apiservername == "" || apiport == 0 {
			fmt.Errorf("Server name and port are required!")
			cmd.Help()
			return
		}
		stime, err := veeamtime.GetServerTime(apiservername, apiport)
		if err != nil{
			fmt.Println("Error:", err)
			return
		}
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
