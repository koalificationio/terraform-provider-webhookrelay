package webhookrelay

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/buckets"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/inputs"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

func resourceWebhookrelayInput() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebhookrelayInputCreate,
		Read:   resourceWebhookrelayInputRead,
		// Update: resourceWebhookrelayInputUpdate, // not yet implemented in sdk
		Delete: resourceWebhookrelayInputDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"bucket_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWebhookrelayInputCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	bucketID := d.Get("bucket_id").(string)

	bucketParams := buckets.NewGetV1BucketsBucketIDParams().WithBucketID(bucketID)
	_, err := client.Buckets.GetV1BucketsBucketID(bucketParams)
	if err != nil {
		if _, ok := err.(*buckets.GetV1BucketsBucketIDNotFound); ok {
			return fmt.Errorf("bucket %s does not exist", bucketID)
		}
		return fmt.Errorf("error reading bucket details: %w", err)
	}

	request := &models.InputRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	params := inputs.NewPostV1BucketsBucketIDInputsParams().
		WithBucketID(bucketID).WithBody(request)

	resp, err := client.Inputs.PostV1BucketsBucketIDInputs(params)
	if err != nil {
		return fmt.Errorf("failed creating input: %w", err)
	}
	d.SetId(resp.GetPayload().ID)

	return nil
}

func resourceWebhookrelayInputRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	bucketID := d.Get("bucket_id").(string)

	bucketParams := buckets.NewGetV1BucketsBucketIDParams().WithBucketID(bucketID)
	resp, err := client.Buckets.GetV1BucketsBucketID(bucketParams)
	if err != nil {
		if _, ok := err.(*buckets.GetV1BucketsBucketIDNotFound); ok {
			log.Printf("[WARN] Removing input %s from state because its bucket is gone", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error reading bucket details: %w", err)
	}

	var input *models.Input
	for _, i := range resp.GetPayload().Inputs {
		if d.Id() == i.ID {
			input = i
		}
	}

	if input == nil {
		log.Printf("[WARN] Removing input %s from state because it's gone", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("description", input.Description)

	return nil
}

func resourceWebhookrelayInputUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceWebhookrelayInputDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	params := inputs.NewDeleteV1BucketsBucketIDInputsInputIDParams().
		WithBucketID(d.Get("bucket_id").(string)).WithInputID(d.Id())

	_, err := client.Inputs.DeleteV1BucketsBucketIDInputsInputID(params)
	if err != nil {
		return fmt.Errorf("failed deleting input: %w", err)
	}

	return nil
}
