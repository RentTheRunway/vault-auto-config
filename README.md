# Vault Auto Configuration Tool

## Overview
This tool allows you to store your Vault configuration in a directory structure with yaml files that reflects the Vault
api and then apply that configuration to a Vault instance.  This tool is inspired by HashiCorp's
[article](https://www.hashicorp.com/blog/codifying-vault-policies-and-configuration/) suggesting this as a way of
doing "configuration as code" for Vault, but overcoming some shortcomings of their method.  The main shortcoming of the
article is that the script does not perform deletions when a resource configuration is removed, this tool will handle
that gracefully for the supported resource types.  For example, if you have an okta auth `auth/okta/groups` resource
name `infra.yaml`, and you delete it, performing an `apply` will remove that group from the Vault configuration as well.

## Installation
To build `vault-auto-config` locally, build it with::

```shell script
./build-cli.sh
```

The executable `vault-auto-config` will now be available at the root of the repo.

Run `./vault-auto-config help` to see all available commands.

Alternatively, you can download the appropriate binary from releases and rename and chmod accordingly.

For example:
```shell script
wget -O /usr/local/bin/vault-auto-config https://github.com/RentTheRunway/vault-auto-config/releases/download/v1.0.0/vault-auto-config-linux-amd64
chmod +x /usr/local/bin/vault-auto-config
```

## Running tests
To run tests locally, execute:

```shell script
./run-tests.sh
```

The tests run using docker-compose to create a Vault instance and compare commands against it with sample input.

## Creating a new release
Creating a release is as simple as creating a tag, like `v1.0.0`, and the github actions workflow will automatically cut
a release, build binaries for various operating systems and attach them to the release.


## Tool commands

##### To print out Vault's current configuration, run:
```shell script
vault-auto-config vault-state --url <vault url> --token <vault token>
```

##### To print out your local configuration, run:
```shell script
vault-auto-config file-state --input-dir <config dir>
```

##### To dump Vault's configuration to a directory structure:
```shell script
vault-auto-config dump --url <vault url> --token <vault token> --output-dir <config dir>
```
This requires an empty or non-existent directory unless the `--force` option is passed.

##### To apply your local configuration to vault:
```shell script
vault-auto-config apply --url <vault url> --token <vault token> --input-dir <config dir>
```

##### Secrets
For the `file-state` and `apply` commands, you can optionally pass a sops encrypted secrets yaml file, which will then
be used as values in your configuration files.

For example:

Decrypted `secrets.yaml.dec`:
```yaml
secret: I'm a little teapot short and stout
```

Okta config file at `v1/auth/okta/config.yaml`
```shell script
org_name: renttherunway
api_token: {{ .Secrets.secret }}
base_url: okta.com
```

Then you could run `apply` or `file-state`, passing in sops encrypted secret file:
```shell script
vault-auto-config file-state --input-dir <config dir> --secrets secrets.yaml
```


## Current supported configurable API paths
This tool does not yet implement support for every single auth backend and configuration option, it's being developed
on a need-to-implement basis.  These are the API paths that are configurable via this tool:
```text
/v1/auth/kubernetes/config
/v1/auth/kubernetes/role/*

/v1/auth/approle/role/*
/v1/auth/approle/role/*/role-id
/v1/auth/approle/role/*/secret-id

/v1/auth/okta/config
/v1/auth/okta/groups/*
/v1/auth/okta/users/*

/v1/auth/token/roles/*

/v1/sys/auth/kubernetes
/v1/sys/auth/okta
/v1/sys/auth/approle

/v1/sys/policy
```

## Workflow
The ideal way to use this tool is to always make the configuration changes in your repo, rather than making the
changes in the Vault UI and then dumping them to config files.  The reason for this is that the dump will not
contain secrets, will contain unneeded default values and will also be formatted poorly. (e.g. HCL policy files
will contain "\n")  The `dump` command is more just meant as a starting point if you're starting from scratch.
