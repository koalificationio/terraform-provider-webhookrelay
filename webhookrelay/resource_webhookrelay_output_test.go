package webhookrelay

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/buckets"
)

func TestAccWebhookrelayOutput_Basic(t *testing.T) {
	outputName := testAccPrefix + acctest.RandString(5)
	outputNewName := testAccPrefix + acctest.RandString(5)

	bucketName := testAccPrefix + acctest.RandString(5)

	resName := "webhookrelay_output.foo"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWebhookrelayOutputDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWebhookrelayOutputConfig(outputName, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayOutputExists(resName),
					resource.TestCheckResourceAttr(
						resName, "name", outputName),
					resource.TestCheckResourceAttr(
						resName, "description", "foo"),
					resource.TestCheckResourceAttr(
						resName, "destination", "http://localhost:8080"),
					resource.TestCheckResourceAttr(
						resName, "internal", "true"),
				),
			},
			{
				Config: testAccCheckWebhookrelayOutputConfigUpdated(outputName, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayOutputExists(resName),
					resource.TestCheckResourceAttr(
						resName, "name", outputName),
					resource.TestCheckResourceAttr(
						resName, "destination", "https://example.com:8080"),
					resource.TestCheckResourceAttr(
						resName, "internal", "false"),
				),
			},
			{
				Config: testAccCheckWebhookrelayOutputConfig(outputNewName, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayOutputExists(resName),
					resource.TestCheckResourceAttr(
						resName, "name", outputNewName),
					resource.TestCheckResourceAttr(
						resName, "description", "foo"),
				),
			},
		},
	})
}

func testAccCheckWebhookrelayOutputDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Openapi)
	for _, r := range s.RootModule().Resources {
		if r.Type != "webhookrelay_output" {
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
		for _, i := range resp.GetPayload().Outputs {
			if r.Primary.ID == i.ID {
				return fmt.Errorf("output %s still exists", r.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckWebhookrelayOutputExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if r.Primary.ID == "" {
			return fmt.Errorf("No Output ID is set")
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

		for _, o := range found.Outputs {
			if o.ID == r.Primary.ID {
				return nil
			}
		}

		return fmt.Errorf("output not found: %v - %v", r.Primary.ID, found.Outputs)
	}
}

func testAccCheckWebhookrelayOutputConfig(name, bucket string) string {
	return fmt.Sprintf(`
resource "webhookrelay_bucket" "foo" {
  name        = "%s"
  description = "foo"
}

resource "webhookrelay_output" "foo" {
  name        = "%s"
  description = "foo"
  destination = "http://localhost:8080"
  internal    = true
  bucket_id   = webhookrelay_bucket.foo.id
}`, bucket, name)
}

func testAccCheckWebhookrelayOutputConfigUpdated(name, bucket string) string {
	return fmt.Sprintf(`
resource "webhookrelay_bucket" "foo" {
  name        = "%s"
  description = "foo"
}

resource "webhookrelay_output" "foo" {
  name        = "%s"
  description = "bar"
  destination = "https://example.com:8080"
  internal    = false
  bucket_id   = webhookrelay_bucket.foo.id
}`, bucket, name)
}