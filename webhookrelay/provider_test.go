package webhookrelay

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	testAccPrefix = "tf-acc-"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"webhookrelay": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) { //nolint:deadcode,unused
	if v := os.Getenv("RELAY_KEY"); v == "" {
		t.Fatal("RELAY_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("RELAY_SECRET"); v == "" {
		t.Fatal("RELAY_SECRET must be set for acceptance tests")
	}
}
