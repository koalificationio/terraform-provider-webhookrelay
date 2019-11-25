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
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/buckets"
)

func init() {
	resource.AddTestSweepers("webhookrelay_bucket", &resource.Sweeper{
		Name: "webhookrelay_bucket",
		F: func(region string) error {
			config, err := sharedConfigForRegion(region)
			if err != nil {
				return fmt.Errorf("error getting client: %s", err)
			}

			client := config.(*client.Openapi)
			list, err := client.Buckets.GetV1Buckets(buckets.NewGetV1BucketsParams())
			if err != nil {
				return fmt.Errorf("error listing buckets: %s", err)
			}

			for _, b := range list.GetPayload() {
				if strings.HasPrefix(b.Name, testAccPrefix) {
					for _, i := range b.Inputs {
						err := deleteInput(client, b.ID, i.ID)
						if err != nil {
							log.Printf("error destroying input %s during sweep: %s", i.ID, err)
						}
					}

					params := buckets.NewDeleteV1BucketsBucketIDParams().WithBucketID(b.ID)
					_, err := client.Buckets.DeleteV1BucketsBucketID(params)
					if err != nil {
						log.Printf("error destroying bucket %s during sweep: %s", b.ID, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccWebhookrelayBucket_Basic(t *testing.T) {
	bucket := testAccPrefix + acctest.RandString(5)
	bucketNew := testAccPrefix + acctest.RandString(5)

	resName := "webhookrelay_bucket.foo"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWebhookrelayBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWebhookrelayBucketConfig(bucket),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayBucketExists(resName),
					resource.TestCheckResourceAttr(
						resName, "name", bucket),
					resource.TestCheckResourceAttr(
						resName, "description", "foo"),
					resource.TestCheckResourceAttr(
						resName, "default_input.#", "1"),
					resource.TestCheckResourceAttr(
						resName, "input.#", "0"),
				),
			},
			{
				Config: testAccCheckWebhookrelayBucketConfigUpdated(bucket),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayBucketExists(resName),
					resource.TestCheckResourceAttr(
						resName, "name", bucket),
					resource.TestCheckResourceAttr(
						resName, "description", "bar"),
				),
			},
			{
				Config: testAccCheckWebhookrelayBucketConfigUpdated(bucketNew),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayBucketExists(resName),
					resource.TestCheckResourceAttr(
						resName, "name", bucketNew),
					resource.TestCheckResourceAttr(
						resName, "description", "bar"),
					resource.TestCheckResourceAttr(
						resName, "default_input.#", "1"),
				),
			},
			{
				Config: testAccCheckWebhookrelayBucketConfigDeleteDefault(bucketNew),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayBucketExists(resName),
					resource.TestCheckResourceAttr(
						resName, "default_input.#", "0"),
				),
			},
			{
				Config: testAccCheckWebhookrelayBucketConfigRemoveDesc(bucketNew),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebhookrelayBucketExists(resName),
					resource.TestCheckResourceAttr(
						resName, "default_input.#", "0"),
					resource.TestCheckResourceAttr(
						resName, "description", ""),
				),
			},
		},
	})
}

func testAccCheckWebhookrelayBucketDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Openapi)
	for _, r := range s.RootModule().Resources {
		if r.Type != "webhookrelay_bucket" {
			continue
		}

		params := buckets.NewGetV1BucketsBucketIDParams().WithBucketID(r.Primary.ID)

		_, err := client.Buckets.GetV1BucketsBucketID(params)
		if err != nil {
			if _, ok := err.(*buckets.GetV1BucketsBucketIDNotFound); ok {
				return nil
			}
			return fmt.Errorf("failed getting buckets: %w", err)
		}
		return fmt.Errorf("bucket %s still exists", r.Primary.ID)
	}
	return nil
}

func testAccCheckWebhookrelayBucketExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if r.Primary.ID == "" {
			return fmt.Errorf("No Bucket ID is set")
		}

		client := testAccProvider.Meta().(*client.Openapi)

		params := buckets.NewGetV1BucketsBucketIDParams().WithBucketID(r.Primary.ID)

		resp, err := client.Buckets.GetV1BucketsBucketID(params)
		if err != nil {
			if v, ok := err.(*buckets.GetV1BucketsBucketIDNotFound); ok {
				return fmt.Errorf("bucket not found: %w", v)
			}
			return fmt.Errorf("failed getting buckets: %v", err.Error())
		}

		found := resp.GetPayload()

		if found.ID != r.Primary.ID {
			return fmt.Errorf("bucket not found: %v - %v", r.Primary.ID, found)
		}

		return nil
	}
}

func testAccCheckWebhookrelayBucketConfig(name string) string {
	return fmt.Sprintf(`
resource "webhookrelay_bucket" "foo" {
  name        = "%s"
  description = "foo"
}`, name)
}

func testAccCheckWebhookrelayBucketConfigUpdated(name string) string {
	return fmt.Sprintf(`
resource "webhookrelay_bucket" "foo" {
  name        = "%s"
  description = "bar"
}`, name)
}

func testAccCheckWebhookrelayBucketConfigDeleteDefault(name string) string {
	return fmt.Sprintf(`
resource "webhookrelay_bucket" "foo" {
  name                 = "%s"
  description          = "bar"
  delete_default_input = true
}`, name)
}

func testAccCheckWebhookrelayBucketConfigRemoveDesc(name string) string {
	return fmt.Sprintf(`
resource "webhookrelay_bucket" "foo" {
  name                 = "%s"
  delete_default_input = true
}`, name)
}
