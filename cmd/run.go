/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"vaultpoc/internal/pkg/vault"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the vault POC demo",
	Long:  `This is an example of how we could leverage consul vault and the viper config to retrieve secrets from the vault`,
	Run: func(cmd *cobra.Command, args []string) {
		var client, err = vault.InitClient()
		if err != nil {
			log.Fatalf("Unable to load vault client: %s", err)
		}
		vault.UpdateViperConfigFromVault(viper.GetViper(), client)

		// Display all the Viper settings after the update
		show, _ := cmd.Flags().GetBool("show")
		if show {
			log.Printf("\n\nViper settings after Vault update:")
			for _, key := range viper.AllKeys() {
				log.Printf("%s: %v\n", key, viper.Get(key))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	runCmd.Flags().BoolP("show", "s", false, "Show updated viper config")
}
