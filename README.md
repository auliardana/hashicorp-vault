# HashiCorp Vault Setup - Non-Development Mode

Welcome to the HashiCorp Vault setup guide! This guide will walk you through configuring Vault in non-development mode using Docker, enabling you to manage secrets and sensitive data securely.

## Getting Started

### 1. Set Up the Vault Server

To start the Vault server in non-development mode, run the following command:

```bash
docker run -d \
  --hostname vault \
  --name vault-non-dev \
  --cap-add=IPC_LOCK \
  -e VAULT_ADDR=http://0.0.0.0:8200 \
  -e VAULT_API_ADDR=http://0.0.0.0:8200 \
  -p 8200:8200 \
  -v $(pwd)/conf:/vault/conf \
  -v vault-data:/vault/data \
  --entrypoint vault \
  --network my-network \
  hashicorp/vault:latest server -config=/vault/conf/config.hcl
```

### 2. Access the UI
Once the server is running, you can access the Vault UI at http://0.0.0.0:8200.

### 3. Initialize Vault
```bash
vault operator init
```

### 4. Seal/Unseal the Vault
Seal the Vault: This command will make the server inaccessible and stop it from sharing data
```bash
vault operator seal
```

Unseal the Vault: Use the following command and input the unseal key repeatedly until the unseal process is complete.
```bash
vault operator unseal
```
Login to Vault: Use the token received during initialization.
```bash
vault login <token>
```

### 5. Enable the KV Secret Engine
To enable the KV secret engine, run:
```bash
vault secrets enable -version=2 kv
vault secrets enable -path=<path-to-store-secret>/ <plugin-name>
# Example: 
vault secrets enable -path=my-secret/ kv
```

### 6. Add/Store New Secrets
To add new secrets, use:
```bash
vault kv put kv/my-secret postgres_username=adminvault postgres_password=vault
```

### 7. Environment Variables
```bash
VAULT_TOKEN=<Token>
VAULT_ADDR="http://0.0.0.0:8200"
```

### RUNNNNNNNN
```bash
go run ./Api/main.go
```