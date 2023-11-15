/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// timezoneCmd represents the timezone command
var timezoneCmd = &cobra.Command{
	Use:   "timezone",
	Short: "Get the current tie in a given timezone",
	Long: `Get the current time in a given timezone.
               This command takes one argument, the timezone you want to get the current time in.
               It returns the current time in RFC1123 format.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a timezone as an argument.")
			fmt.Println("For example : \"Europe/London\"")
			return
		}

		timezone := args[0]
		location, _ := time.LoadLocation(timezone)
		dateFlag, _ := cmd.Flags().GetString("date")
		var timeStr string

		if dateFlag != "" {
			timeStr = time.Now().In(location).Format(dateFlag)
		} else {
			timeStr = time.Now().In(location).Format(time.Kitchen)
		}
		fmt.Printf("Current time in %v: %v\n", timezone, timeStr)
	},
}

func init() {
	rootCmd.AddCommand(timezoneCmd)

	timezoneCmd.PersistentFlags().String("date", "", "returns the date in a time zone in a specified format")
	timezoneCmd.Flags().String("date", "", "Date for which to get the time (format: yyyy-mm-dd)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timezoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timezoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
