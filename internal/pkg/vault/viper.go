package vault

import (
	"fmt"
	"github.com/hashicorp/vault-client-go"
	"github.com/spf13/viper"
	"log"
	"strings"
	_ "strings"
)

func UpdateViperConfigFromVault(v *viper.Viper, client *vault.Client) {
	// Enumerate all configuration keys and values
	for _, key := range v.AllKeys() {
		value := fmt.Sprintf("%s", v.Get(key))
		if strings.HasPrefix(value, "vault:") {
			log.Printf("Update viper config %s from %v", key, value)
			secretPath, secretField := ParseVaultPath(value)
			secret, err := GetSecretFromVault(client, secretPath, secretField)
			if err != nil {
				log.Fatalf("Unable to read secret: %v", err)
			}
			viper.Set(key, secret)
		}
	}
}
