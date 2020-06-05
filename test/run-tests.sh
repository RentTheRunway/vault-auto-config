#!/usr/bin/env bash

# Generate a gpg key and encrypt our secrets
cat >key-spec <<EOF
     Key-Type: default
     Subkey-Type: default
     Name-Real: Infrastructure
     Name-Comment: For unit testing with sops
     Name-Email: infrastructure@renttherunway.com
     Expire-Date: 0
     %no-protection
     %commit
EOF

gpg --batch --generate-key key-spec
FINGERPRINT=$(gpg --list-secret-keys | egrep [A-F0-9]{40} | awk '{$1=$1};1')

pushd ./vault-auto-config > /dev/null

# Create sops configuration file
cat >.sops.yaml <<EOF
---
creation_rules:
  - pgp: ${FINGERPRINT}
EOF

# Encrypt our secrets
sops --encrypt --input-type yaml --output-type yaml secrets.yaml.dec > secrets.yaml

popd > /dev/null

# Run the unit tests
go test ./vault-auto-config
