package vault

import (
	"context"
	"fmt"
	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	"log"
	"os"
	"strings"
	"time"
)

func InitClient() (client *vault.Client, err error) {
	var vaultAddr = os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		vaultAddr = "http://127.0.0.1:8200"
	}
	var vaultToken = os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		log.Printf("WARN: Could not load VAULT_TOKEN from environment")
		vaultToken = "root"
	}
	client, err = vault.New(
		vault.WithAddress(vaultAddr),
		vault.WithRequestTimeout(30*time.Second),
	)
	err = client.SetToken(vaultToken)
	if err != nil {
		log.Fatalf("Unable to set vault token: %s", err)
	}
	return
}

func ParseVaultPath(path string) (string, string) {
	parts := strings.Split(path, "#")
	return strings.TrimPrefix(parts[0], "vault:"), parts[1]
}

func GetSecretFromVault(client *vault.Client, path string, field string) (string, error) {
	// Define the context
	ctx := context.Background()
	// Retrieve the secret from Vault using the KV v2 client
	secret, err := client.Read(ctx, path)
	if err != nil {
		log.Printf("Unable to read secret: %v", err)
	} else {
		// Access the nested data for KV v2
		if rawData, ok := secret.Data["data"].(map[string]interface{}); ok {
			if secretValue, exists := rawData[field]; exists {
				//fmt.Printf("Secret value: %v\n", secretValue)
				return fmt.Sprint(secretValue), nil
			} else {
				log.Fatalf("Secret key not found: %s", field)
			}
		} else {
			log.Fatalf("Failed to retrieve data from secret at path: %s", path)
		}
	}
	return "NOTFOUND", nil
}

func UpdateVaultWithDefaults(client *vault.Client) {
	data := map[string]interface{}{
		"vaultpoc/db":  "user=user, password=password",
		"vaultpoc/api": "key=apikey",
	}
	for k, v := range data {
		log.Printf("key: %s, Value: %s", k, v)
		err := SetSecretInVault(client, k, fmt.Sprint(v))
		if err != nil {
			log.Fatalf("Unable to set vault token: %s", err)
		}

	}
}

func SetSecretInVault(client *vault.Client, secretPath string, secretValues string) error {
	// Define the context
	ctx := context.Background()

	// write a secret (overriding any existing vaules)
	_, err := client.Secrets.KvV2Write(ctx, secretPath, schema.KvV2WriteRequest{Data: stringToMap(secretValues)},
		vault.WithMountPath("secret"),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("secret written successfully to %s", secretPath)

	// Update the field with the new secret
	return nil
}

func stringToMap(data string) map[string]any {
	// Initialize an empty map
	result := make(map[string]any)

	// Split the input string by commas to separate each key-value pair
	pairs := strings.Split(data, ",")
	for _, pair := range pairs {
		// Split each pair by '=' to separate the key and the value
		kv := strings.Split(strings.TrimSpace(pair), "=")
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			// Add the key-value pair to the map
			result[key] = value
		}
	}
	return result
}
