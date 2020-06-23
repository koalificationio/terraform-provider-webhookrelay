package webhookrelay

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/functions"
)

func init() {
	resource.AddTestSweepers("webhookrelay_function", &resource.Sweeper{
		Name: "webhookrelay_function",
		F: func(region string) error {
			config, err := sharedConfigForRegion(region)
			if err != nil {
				return fmt.Errorf("error getting client: %s", err)
			}

			client := config.(*client.Openapi)
			list, err := client.Functions.GetV1Functions(functions.NewGetV1FunctionsParams())
			if err != nil {
				return fmt.Errorf("error listing functions: %s", err)
			}

			for _, f := range list.GetPayload() {
				if strings.HasPrefix(f.Name, testAccPrefix) {
					params := functions.NewDeleteV1FunctionsFunctionIDParams().WithFunctionID(f.ID)
					_, err := client.Functions.DeleteV1FunctionsFunctionID(params)
					if err != nil {
						log.Printf("error destroying function %s during sweep: %s", f.ID, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccWebhookrelayFunction_Basic(t *testing.T) {
	function := testAccPrefix + acctest.RandString(5)
	functionNew := testAccPrefix + acctest.RandString(5)

	payload := "r:SetRequestMethod('PUT')"
	payloadNew := "r:SetRequestMethod('DELETE')"

	payloadBase64 := base64.StdEncoding.EncodeToString([]byte(payload))
	payloadNewBase64 := base64.StdEncoding.EncodeToString([]byte(payloadNew))

	resName := "webhookrelay_function.foo"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWebhookrelayFunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWebhookrelayFunctionConfig(function, payload),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayFunctionExists(resName),
					resource.TestCheckResourceAttr(
						resName, "name", function),
					resource.TestCheckResourceAttr(
						resName, "driver", "lua"),
					resource.TestCheckResourceAttr(
						resName, "payload", payloadBase64),
				),
			},
			{
				Config: testAccCheckWebhookrelayFunctionConfig(functionNew, payload),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayFunctionExists(resName),
					resource.TestCheckResourceAttr(
						resName, "name", functionNew),
					resource.TestCheckResourceAttr(
						resName, "driver", "lua"),
					resource.TestCheckResourceAttr(
						resName, "payload", payloadBase64),
				),
			},
			{
				Config: testAccCheckWebhookrelayFunctionConfig(functionNew, payloadNew),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayFunctionExists(resName),
					resource.TestCheckResourceAttr(
						resName, "name", functionNew),
					resource.TestCheckResourceAttr(
						resName, "driver", "lua"),
					resource.TestCheckResourceAttr(
						resName, "payload", payloadNewBase64),
				),
			},
		},
	})
}

func testAccCheckWebhookrelayFunctionDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Openapi)
	for _, r := range s.RootModule().Resources {
		if r.Type != "webhookrelay_function" {
			continue
		}

		params := functions.NewGetV1FunctionsFunctionIDParams().WithFunctionID(r.Primary.ID)

		_, err := client.Functions.GetV1FunctionsFunctionID(params)
		if err != nil {
			if _, ok := err.(*functions.GetV1FunctionsFunctionIDNotFound); ok {
				return nil
			}
			return fmt.Errorf("failed getting functions: %w", err)
		}
		return fmt.Errorf("functions %s still exists", r.Primary.ID)
	}
	return nil
}

func testAccCheckWebhookrelayFunctionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if r.Primary.ID == "" {
			return fmt.Errorf("No Function ID is set")
		}

		client := testAccProvider.Meta().(*client.Openapi)

		params := functions.NewGetV1FunctionsFunctionIDParams().WithFunctionID(r.Primary.ID)

		resp, err := client.Functions.GetV1FunctionsFunctionID(params)
		if err != nil {
			if v, ok := err.(*functions.DeleteV1FunctionsFunctionIDNotFound); ok {
				return fmt.Errorf("function not found: %w", v)
			}
			return fmt.Errorf("failed getting function: %v", err.Error())
		}

		found := resp.GetPayload()

		if found.ID != r.Primary.ID {
			return fmt.Errorf("function not found: %v - %v", r.Primary.ID, found)
		}

		return nil
	}
}

func testAccCheckWebhookrelayFunctionConfig(name, payload string) string {
	return fmt.Sprintf(`
resource "webhookrelay_function" "foo" {
  name    = "%s"
  payload = base64encode("%s")
  driver  = "lua"
}`, name, payload)
}
