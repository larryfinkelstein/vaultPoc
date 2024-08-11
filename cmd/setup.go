/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"vaultpoc/internal/pkg/vault"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set vaules in interprise vault",
	Long: `This is insecure and is only for development/testing purposes

	This command can be used to initialize vault with default secrets`,
	Run: func(cmd *cobra.Command, args []string) {
		var client, err = vault.InitClient()
		if err != nil {
			log.Fatalf("Unable to load vault client: %s", err)
		}
		vault.UpdateVaultWithDefaults(client)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
