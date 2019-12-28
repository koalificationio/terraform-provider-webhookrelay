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
		Update: resourceWebhookrelayInputUpdate,
		Delete: resourceWebhookrelayInputDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bucket_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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

	request := &models.Input{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	if v, ok := d.GetOk("status_code"); ok {
		request.StatusCode = v.(int64)
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
	d.Set("status_code", int(input.StatusCode))

	return nil
}

func resourceWebhookrelayInputUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	bucketID := d.Get("bucket_id").(string)

	if d.HasChanges("name", "description") {
		request := &models.Input{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
		}

		if v, ok := d.GetOk("status_code"); ok {
			request.StatusCode = int64(v.(int))
		}

		params := inputs.NewPutV1BucketsBucketIDInputsInputIDParams().
			WithBucketID(bucketID).
			WithInputID(d.Id()).
			WithBody(request)

		_, err := client.Inputs.PutV1BucketsBucketIDInputsInputID(params)
		if err != nil {
			return fmt.Errorf("error updating input: %w", err)
		}
	}

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
