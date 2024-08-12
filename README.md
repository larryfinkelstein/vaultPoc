# Vault POC

The goal of this project is to integrate Hashicorp Vault with Viper so that we can load viper configs from 
from enterprise vault.

## Run the vault server (after installing) in dev mode

```
vault server -dev
```

Pay attention to the startup script and make a note of the `Root Token`.  You will need to set this
environment variable:

```
set VAULT_TOKEN=hvs.abcdefg
```

## Set up vault

```
vault kv put -mount=secret vaultpoc/db user="user" password="password"
vault kv put -mount=secret vaultpoc/apikey="api-key"
```

Running this command should return the results from above:

```
vault kv get secret/vaultpoc/db

===== Secret Path =====
secret/data/vaultpoc/db

======= Metadata =======
Key                Value
---                -----
created_time       2024-08-11T15:42:53.6505757Z
custom_metadata    <nil>
deletion_time      n/a
destroyed          false
version            3

====== Data ======
Key         Value
---         -----
password    password
user        user
```

# Running this application

```
>go run .\main.go
This is an example of how we could leverage consul vault and the viper config to
 retrieve secrets from the vault

Usage:
  vaultpoc [command]

Available Commands:
  help        Help about any command
  run         Run the vault POC demo
  setup       Set vaules in interprise vault

Flags:
  -h, --help   help for vaultpoc

Use "vaultpoc [command] --help" for more information about a command.
```

## Run command

If we `run` this command, we should see that nothing has been set.

Don't forget to set the `VAULT_TOKEN` environment variable or else you will see `403 Forbidden` errors.

```
12:46:14 Update viper config api.key from vault:secret/data/vaultpoc/
api#key
12:46:14 Unable to read secret: 403 Forbidden: 2 errors occurred:
        * permission denied
        * invalid token

> set VAULT_TOKEN=vault token
> go run .\main.go run --show

12:50:49 Using config file: C:\Users\larryf\GolandProjects\vaultpoc\config\config.yaml)
12:50:49 Update viper config api.key from vault:secret/data/vaultpoc/api#key
12:50:49 Unable to read secret: 404 Not Found: {"errors":[]}
12:50:49 Update viper config database.user from vault:secret/data/vaultpoc/db#user
12:50:49 Unable to read secret: 404 Not Found: {"errors":[]}
12:50:49 Update viper config database.password from vault:secret/data
/vaultpoc/db#password
12:50:49 Unable to read secret: 404 Not Found: {"errors":[]}
12:50:49

Viper settings after Vault update:
12:50:49 env: dev
12:50:49 api.key: NOTFOUND
12:50:49 database.user: NOTFOUND
12:50:49 database.password: NOTFOUND
```

## Setup command

```
>go run .\main.go setup

12:53:00 Using config file: C:\Users\larryf\GolandPro
jects\vaultpoc\config\config.yaml)
12:53:00 key: vaultpoc/db, Value: user=user, password=password
12:53:00 secret written successfully to vaultpoc/db
12:53:00 key: vaultpoc/api, Value: key=apikey
12:53:00 secret written successfully to vaultpoc/api
```

## Run command again

```
>go run .\main.go run --show
12:53:58 Using config file: C:\Users\larryf\GolandProjects\vaultpoc\config\config.yaml)
12:53:58 Update viper config database.password from vault:secret/data/vaultpoc/db#password
12:53:58 Update viper config api.key from vault:secret/data/vaultpoc/api#key
12:53:58 Update viper config database.user from vault:secret/data/vaultpoc/db#user
12:53:58

Viper settings after Vault update:
12:53:58 env: dev
12:53:58 api.key: apikey
12:53:58 database.password: password
12:53:58 database.user: user
```
# Docker support for running vault

Start docker container running

```
docker pull hashicorp/vault:latest
```
```
docker compose up -d
```
Make a note of the docker machine IP addres.

```
set VAULT_TOKEN=root
set VAULT_ADDR=http://192.168.99.100:8200

go run .\main.go setup
21:09:45 Using config file: vaultpoc\config\config.yaml
21:09:45 key: vaultpoc/api, Value: key=apikey
21:09:45 secret written successfully to vaultpoc/api
21:09:45 key: vaultpoc/db, Value: user=user, password=password
21:09:45 secret written successfully to vaultpoc/db

go run .\main.go run -s
21:09:59 Using config file: vaultpoc\config\config.yaml
21:09:59 Update viper config database.user from vault:secret/data/vaultpoc/db#user
21:09:59 Update viper config database.password from vault:secret/data/vaultpoc/db#password
21:09:59 Update viper config api.key from vault:secret/data/vaultpoc/api#key
21:09:59

Viper settings after Vault update:
21:09:59 database.user: user
21:09:59 database.password: password
21:09:59 api.key: apikey
21:09:59 env: dev
```

And finally, you can stop docker.

```
docker compose down
```
