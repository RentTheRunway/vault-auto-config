package vault_auto_config_test

import (
	"os"
	"path"
	"testing"

	pkg "github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config"
	"github.com/hashicorp/vault/api"
)

var autoConfig *pkg.VaultAutoConfig
var client *api.Client

const vaultAddress = "http://vault:8200"
const vaultToken = "dev-vault-token"

// sadly, verifying these is not as simple as just walking the directory structure and doing a GET
// request for that path, because the vault api is fairly inconsistent from one endpoint to another
var resources = []string{
	"/v1/auth/approle/role/java-app",
	"/v1/auth/approle/role/java-app/role-id",
	"/v1/auth/kubernetes/config",
	"/v1/auth/kubernetes/role/tech",
	"/v1/auth/okta/config",
	"/v1/auth/okta/groups/tech",
	"/v1/auth/okta/users/bob",
	"/v1/auth/token/roles/tech",
	"/v1/sys/auth/kubernetes/tune",
	"/v1/sys/auth/okta/tune",
	"/v1/sys/policy/tech",
}

var files = []string{
	"v1/auth/approle/role/java-app.yaml",
	"v1/auth/approle/role/java-app/role-id.yaml",
	"v1/auth/approle/role/java-app/secret-id.yaml",
	"v1/auth/kubernetes/role/tech.yaml",
	"v1/auth/kubernetes/config.yaml",
	"v1/auth/okta/groups/tech.yaml",
	"v1/auth/okta/users/bob.yaml",
	"v1/auth/okta/config.yaml",
	"v1/auth/token/roles/tech.yaml",
	"v1/sys/auth/approle.yaml",
	"v1/sys/auth/kubernetes.yaml",
	"v1/sys/auth/okta.yaml",
	"v1/sys/policy/tech.yaml",
}

func init() {
	autoConfig = pkg.NewVaultAutoConfig()

	var err error
	if client, err = api.NewClient(&api.Config{Address: vaultAddress}); err != nil {
		panic(err)
	}

	client.SetToken(vaultToken)
}

// Tests vault-auto-config by applying a maximal configuration that tests all configurable resources, checking that the
// resources exist in vault, then applying an empty config and verifying the resources are all removed
func TestApply(t *testing.T) {
	if err := autoConfig.Apply(vaultAddress, vaultToken, "full-config", "secrets.yaml"); err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	for _, resource := range resources {
		verifyResourceExists(t, resource)
	}

	if t.Failed() {
		return
	}

	// applying an empty config should remove these resources
	if err := autoConfig.Apply(vaultAddress, vaultToken, "empty-config", ""); err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	for _, resource := range resources {
		verifyResourceNotExists(t, resource)
	}
}

// Tests vault-auto-config by applying a maximal configuration that tests all configurable resources, then doing a dump
// of that configuration and testing that the expected files exist
func TestDump(t *testing.T) {
	if err := autoConfig.Apply(vaultAddress, vaultToken, "full-config", "secrets.yaml"); err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if err := autoConfig.Dump(vaultAddress, vaultToken, "dump", false); err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	for _, file := range files {
		verifyFileExists(t, "dump", file)
	}
}

func verifyResourceExists(t *testing.T, resource string) {
	verifyResource(t, resource, true)
}

func verifyResourceNotExists(t *testing.T, resource string) {
	verifyResource(t, resource, false)
}

func verifyResource(t *testing.T, resource string, exists bool) {
	_, err := client.RawRequest(client.NewRequest("GET", resource))

	if err != nil && exists {
		t.Errorf("Could not find resource %s", resource)
		t.Error(err)
		t.Fail()
		return
	}

	if err == nil && !exists {
		t.Errorf("Found unexpected resource %s", resource)
		t.Fail()
		return
	}
}

func verifyFileExists(t *testing.T, dir string, file string) {
	filePath := path.Join(dir, file)
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		t.Errorf("File not found %s", filePath)
		t.Fail()
		return
	}
}
