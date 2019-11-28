package webhookrelay

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/buckets"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/outputs"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

func resourceWebhookrelayOutput() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebhookrelayOutputCreate,
		Read:   resourceWebhookrelayOutputRead,
		Update: resourceWebhookrelayOutputUpdate,
		Delete: resourceWebhookrelayOutputDelete,
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
			},
			"destination": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"internal": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"bucket_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWebhookrelayOutputCreate(d *schema.ResourceData, meta interface{}) error {
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

	request := &models.Output{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Destination: d.Get("destination").(string),
		Internal:    d.Get("internal").(bool),
	}

	params := outputs.NewPostV1BucketsBucketIDOutputsParams().
		WithBucketID(bucketID).WithBody(request)

	resp, err := client.Outputs.PostV1BucketsBucketIDOutputs(params)
	if err != nil {
		return fmt.Errorf("failed creating output: %w", err)
	}
	d.SetId(resp.GetPayload().ID)

	return nil
}

func resourceWebhookrelayOutputRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	bucketID := d.Get("bucket_id").(string)

	bucketParams := buckets.NewGetV1BucketsBucketIDParams().WithBucketID(bucketID)
	resp, err := client.Buckets.GetV1BucketsBucketID(bucketParams)
	if err != nil {
		if _, ok := err.(*buckets.GetV1BucketsBucketIDNotFound); ok {
			log.Printf("[WARN] Removing output %s from state because its bucket is gone", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error reading bucket details: %w", err)
	}

	var output *models.Output
	for _, o := range resp.GetPayload().Outputs {
		if d.Id() == o.ID {
			output = o
		}
	}

	if output == nil {
		log.Printf("[WARN] Removing output %s from state because it's gone", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("description", output.Description)
	d.Set("destination", output.Destination)
	d.Set("internal", output.Internal)

	return nil
}

func resourceWebhookrelayOutputUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	if d.HasChange("description") || d.HasChange("destination") || d.HasChange("internal") {
		request := &models.Output{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Destination: d.Get("destination").(string),
			Internal:    d.Get("internal").(bool),
		}

		params := outputs.NewPutV1BucketsBucketIDOutputsOutputIDParams()
		params.SetBucketID(d.Get("bucket_id").(string))
		params.SetOutputID(d.Id())
		params.SetBody(request)

		_, err := client.Outputs.PutV1BucketsBucketIDOutputsOutputID(params)
		if err != nil {
			return fmt.Errorf("failed updating output: %w", err)
		}
	}

	return nil
}

func resourceWebhookrelayOutputDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	params := outputs.NewDeleteV1BucketsBucketIDOutputsOutputIDParams().
		WithBucketID(d.Get("bucket_id").(string)).WithOutputID(d.Id())

	_, err := client.Outputs.DeleteV1BucketsBucketIDOutputsOutputID(params)
	if err != nil {
		return fmt.Errorf("failed deleting output: %w", err)
	}

	return nil
}
