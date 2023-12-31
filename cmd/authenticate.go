/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"

	veeamauthenticate "github.com/patrikcze/go-veeam-cli/packages/authenticate"
	"github.com/spf13/cobra"
)

var (
	servername string
	port       int
)

// authenticateCmd represents the authenticate command
var authenticateCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "Authenticate with the Veeam B&R RestAPI",
	Long:  `This command allows you to authenticate with the Veeam B&R RestAPI using your username and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		if username == "" || password == "" {
			fmt.Println("Username and password are required!")
			cmd.Help()
			return
		}

		// generate a random key for encryption
		key := make([]byte, 32) // generates a random 32-byte key for AES-256
		if _, err := rand.Read(key); err != nil {
			log.Fatal(err)
		}

		// Store random key into Environment variable to be able to decrypt.
		os.Setenv("TOKEN_ENCRYPTION_KEY", string(key))

		// pass the key to the Authenticate function
		token, err := veeamauthenticate.Authenticate(servername, username, password, port, key)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Authentication successful!")
		fmt.Println("Access Token:", token.AccessToken)
	},
}

func init() {
	rootCmd.AddCommand(authenticateCmd)

	rootCmd.PersistentFlags().StringVar(&servername, "servername", "", "Veeam B&R server name")
	rootCmd.PersistentFlags().IntVar(&port, "port", 9419, "Veeam B&R server port")

	authenticateCmd.Flags().String("username", "", "Your Veeam B&R username.")
	authenticateCmd.Flags().String("password", "", "Your Veeam B&R password.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authenticateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authenticateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
