package webhookrelay

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/tokens"
)

func init() {
	resource.AddTestSweepers("webhookrelay_token", &resource.Sweeper{
		Name: "webhookrelay_token",
		F: func(region string) error {
			config, err := sharedConfigForRegion(region)
			if err != nil {
				return fmt.Errorf("error getting client: %s", err)
			}

			client := config.(*client.Openapi)
			list, err := client.Tokens.GetV1Tokens(tokens.NewGetV1TokensParams())
			if err != nil {
				return fmt.Errorf("error listing tokens: %s", err)
			}

			for _, t := range list.GetPayload() {
				if strings.HasPrefix(t.Description, testAccPrefix) {
					params := tokens.NewDeleteV1TokensTokenIDParams().WithTokenID(t.ID)

					_, err := client.Tokens.DeleteV1TokensTokenID(params)
					if err != nil {
						log.Printf("error destroying token: %v", err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccWebhookrelayToken_Basic(t *testing.T) {
	tokenDescription := testAccPrefix + acctest.RandString(5)
	tokenNewDescription := testAccPrefix + acctest.RandString(5)

	resName := "webhookrelay_token.foo"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWebhookrelayTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWebhookrelayTokenConfig(tokenNewDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayTokenExists(resName),
					resource.TestCheckResourceAttr(
						resName, "description", tokenNewDescription),
					resource.TestCheckResourceAttr(
						resName, "api_access", "enabled"),
					resource.TestCheckResourceAttr(
						resName, "scopes.#", "0"),
				),
			},
			{
				Config: testAccCheckWebhookrelayTokenConfigUpdated(tokenDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayTokenExists(resName),
					resource.TestCheckResourceAttr(
						resName, "description", tokenDescription),
					resource.TestCheckResourceAttr(
						resName, "api_access", "disabled"),
					resource.TestCheckResourceAttr(
						resName, "scopes.#", "1"),
					resource.TestCheckResourceAttr(
						resName, "scopes.0.tunnels.#", "2"),
					resource.TestCheckResourceAttr(
						resName, "scopes.0.tunnels.0", "test1"),
				),
			},
			{
				Config: testAccCheckWebhookrelayTokenConfigBuckets(tokenDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayTokenExists(resName),
					resource.TestCheckResourceAttr(
						resName, "description", tokenDescription),
					resource.TestCheckResourceAttr(
						resName, "api_access", "disabled"),
					resource.TestCheckResourceAttr(
						resName, "scopes.#", "1"),
					resource.TestCheckResourceAttr(
						resName, "scopes.0.tunnels.#", "0"),
					resource.TestCheckResourceAttr(
						resName, "scopes.0.buckets.0", "test1"),
				),
			},
			{
				Config: testAccCheckWebhookrelayTokenConfig(tokenNewDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayTokenExists(resName),
					resource.TestCheckResourceAttr(
						resName, "description", tokenNewDescription),
					resource.TestCheckResourceAttr(
						resName, "api_access", "enabled"),
					resource.TestCheckResourceAttr(
						resName, "scopes.#", "0"),
				),
			},
		},
	})
}

func testAccCheckWebhookrelayTokenDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Openapi)
	for _, r := range s.RootModule().Resources {
		if r.Type != "webhookrelay_token" {
			continue
		}

		params := tokens.NewGetV1TokensParams()

		resp, err := client.Tokens.GetV1Tokens(params)
		if err != nil {
			return fmt.Errorf("failed getting tokens: %w", err)
		}
		for _, t := range resp.GetPayload() {
			if r.Primary.ID == t.ID {
				return fmt.Errorf("token %s still exists", r.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckWebhookrelayTokenExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if r.Primary.ID == "" {
			return fmt.Errorf("No Token ID is set")
		}

		client := testAccProvider.Meta().(*client.Openapi)

		params := tokens.NewGetV1TokensParams()

		resp, err := client.Tokens.GetV1Tokens(params)
		if err != nil {
			return fmt.Errorf("failed getting tokens: %w", err)
		}
		for _, t := range resp.GetPayload() {
			if r.Primary.ID == t.ID {
				return nil
			}
		}

		return fmt.Errorf("token not found: %v - %v", r.Primary.ID, resp.GetPayload())
	}
}

func testAccCheckWebhookrelayTokenConfig(name string) string {
	return fmt.Sprintf(`
resource "webhookrelay_token" "foo" {
  description = "%s"
  api_access  = "enabled"
}`, name)
}

func testAccCheckWebhookrelayTokenConfigUpdated(name string) string {
	return fmt.Sprintf(`
resource "webhookrelay_token" "foo" {
  description = "%s"
  api_access  = "disabled"
  scopes {
    tunnels = ["test1", "test2"]
  }
}
`, name)
}

func testAccCheckWebhookrelayTokenConfigBuckets(name string) string {
	return fmt.Sprintf(`
resource "webhookrelay_token" "foo" {
  description = "%s"
  api_access  = "disabled"
  scopes {
    buckets = ["test1", "test2"]
  }
}
`, name)
}
