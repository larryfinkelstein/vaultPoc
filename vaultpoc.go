package main

import (
	"github.com/hashicorp/vault/api"
)

func main() {
	config := &api.Config{Address: "https://vault.example.com"}

	var conf, _ = LoadConfig("config/config.yaml")

	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	// Authenticate and read secrets
	secret, err := client.Logical().Read("secret/data/myapp/" + conf.Env)
	if err != nil {
		panic(err)
	}
	// Use secret.Data to get your secrets
	print(secret)
}
