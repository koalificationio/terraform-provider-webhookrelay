package webhookrelay

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/buckets"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

func TestAccWebhookrelayInput_Basic(t *testing.T) {
	var input, newInput models.Input
	var bucket models.Bucket
	inputName := testAccPrefix + acctest.RandString(5)
	inputNewName := testAccPrefix + acctest.RandString(5)

	bucketName := testAccPrefix + acctest.RandString(5)

	resName := "webhookrelay_input.foo"

	guidCheck := regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWebhookrelayInputDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWebhookrelayInputConfig(inputName, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayInputExists(resName, &input, &bucket),
					resource.TestCheckResourceAttr(
						resName, "name", inputName),
					resource.TestCheckResourceAttr(
						resName, "description", "foo"),
					resource.TestCheckResourceAttr(
						resName, "response_body", "{\"message\":\"test\"}"),
					resource.TestCheckResourceAttr(
						resName, "headers.%", "2"),
					resource.TestCheckResourceAttr(
						resName, "headers.Content-Type", "text/html; charset=utf-8"),
					resource.TestMatchResourceAttr(
						resName, "function_id", guidCheck),
				),
			},
			{
				Config: testAccCheckWebhookrelayInputConfigUpdated(inputName, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayInputExists(resName, &input, &bucket),
					resource.TestCheckResourceAttr(
						resName, "name", inputName),
					resource.TestCheckResourceAttr(
						resName, "description", "bar"),
					resource.TestCheckResourceAttr(
						resName, "status_code", "402"),
					resource.TestCheckResourceAttr(
						resName, "response_body", ""),
					resource.TestCheckResourceAttr(
						resName, "headers.%", "0"),
					resource.TestCheckResourceAttr(
						resName, "function_id", ""),
				),
			},
			{
				Config: testAccCheckWebhookrelayInputConfig(inputNewName, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayInputExists(resName, &newInput, &bucket),
					resource.TestCheckResourceAttr(
						resName, "name", inputNewName),
					resource.TestCheckResourceAttr(
						resName, "description", "foo"),
					resource.TestMatchResourceAttr(
						resName, "function_id", guidCheck),
				),
			},
		},
	})
}

func testAccCheckWebhookrelayInputDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Openapi)
	for _, r := range s.RootModule().Resources {
		if r.Type != "webhookrelay_input" {
			continue
		}

		params := buckets.NewGetV1BucketsBucketIDParams().WithBucketID(r.Primary.Attributes["bucket_id"])

		resp, err := client.Buckets.GetV1BucketsBucketID(params)
		if err != nil {
			if _, ok := err.(*buckets.GetV1BucketsBucketIDNotFound); ok {
				return nil
			}
			return fmt.Errorf("failed getting buckets: %w", err)
		}
		for _, i := range resp.GetPayload().Inputs {
			if r.Primary.ID == i.ID {
				return fmt.Errorf("input %s still exists", r.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckWebhookrelayInputExists(n string, input *models.Input, bucket *models.Bucket) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if r.Primary.ID == "" {
			return fmt.Errorf("No Input ID is set")
		}

		client := testAccProvider.Meta().(*client.Openapi)

		params := buckets.NewGetV1BucketsBucketIDParams().WithBucketID(r.Primary.Attributes["bucket_id"])

		resp, err := client.Buckets.GetV1BucketsBucketID(params)
		if err != nil {
			if v, ok := err.(*buckets.GetV1BucketsBucketIDNotFound); ok {
				return fmt.Errorf("bucket not found: %w", v)
			}
			return fmt.Errorf("failed getting buckets: %v", err.Error())
		}

		found := resp.GetPayload()
		*bucket = *found

		for _, i := range found.Inputs {
			if i.ID == r.Primary.ID {
				*input = *i
				return nil
			}
		}

		return fmt.Errorf("input not found: %v - %v", r.Primary.ID, found.Inputs)
	}
}

func testAccCheckWebhookrelayInputConfig(name, bucket string) string {
	return fmt.Sprintf(`
resource "webhookrelay_bucket" "foo" {
  name                 = "%s"
  description          = "foo"
  delete_default_input = true
}

resource "webhookrelay_input" "foo" {
  name          = "%s"
  description   = "foo"
  bucket_id     = webhookrelay_bucket.foo.id
  response_body = jsonencode(
    map("message", "test")
  )
  headers = {
    Content-Type  = "text/html; charset=utf-8"
    Foo           = "Bar"
  }
  function_id = webhookrelay_function.foo.id
}

resource "webhookrelay_function" "foo" {
  name    = "%s"
  payload = base64encode("r:SetRequestMethod('PUT')")
  driver  = "lua"
}
`, bucket, name, name)
}

func testAccCheckWebhookrelayInputConfigUpdated(name, bucket string) string {
	return fmt.Sprintf(`
resource "webhookrelay_bucket" "foo" {
  name                 = "%s"
  description          = "foo"
  delete_default_input = true
}

resource "webhookrelay_input" "foo" {
  name        = "%s"
  description = "bar"
  bucket_id   = webhookrelay_bucket.foo.id
  status_code = 402
}

resource "webhookrelay_function" "foo" {
  name    = "%s"
  payload = base64encode("r:SetRequestMethod('PUT')")
  driver  = "lua"
}
`, bucket, name, name)
}
