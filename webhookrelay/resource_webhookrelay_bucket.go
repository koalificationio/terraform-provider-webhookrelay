package webhookrelay

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/buckets"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

const (
	defaultInputName = "Default public endpoint"
)

func resourceWebhookrelayBucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebhookrelayBucketCreate,
		Read:   resourceWebhookrelayBucketRead,
		Update: resourceWebhookrelayBucketUpdate,
		Delete: resourceWebhookrelayBucketDelete,
		Exists: resourceWebhookrelayBucketExists,
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
			"websocket_streaming": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ephemeral_webhooks": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_input": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem:     schemaInput(),
			},
			"delete_default_input": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"input": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     schemaInput(),
			},
		},
	}
}

func resourceWebhookrelayBucketCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	request := &models.BucketRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	params := buckets.NewPostV1BucketsParams().WithBody(request)

	resp, err := client.Buckets.PostV1Buckets(params)
	if err != nil {
		return fmt.Errorf("error creating bucket: %w", err)
	}
	d.SetId(resp.GetPayload().ID)

	return resourceWebhookrelayBucketRead(d, meta)
}

func resourceWebhookrelayBucketRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	params := buckets.NewGetV1BucketsBucketIDParams().WithBucketID(d.Id())
	resp, err := client.Buckets.GetV1BucketsBucketID(params)
	if err != nil {
		if _, ok := err.(*buckets.GetV1BucketsBucketIDNotFound); ok {
			log.Printf("[WARN] Removing bucket %s from state because it's gone", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error reading bucket details: %w", err)
	}

	bucket := resp.GetPayload()
	d.Set("description", bucket.Description)
	d.Set("websocket_streaming", bucket.Stream)
	d.Set("ephemeral_webhooks", bucket.Ephemeral)

	var defaultInput, inputs []*models.Input
	for _, i := range bucket.Inputs {
		if i.Name == defaultInputName {
			defaultInput = append(defaultInput, i)
		} else {
			inputs = append(inputs, i)
		}
	}

	if err := d.Set("default_input", flattenInputs(defaultInput)); err != nil {
		return fmt.Errorf("error setting default input: %w", err)
	}

	if err := d.Set("input", flattenInputs(inputs)); err != nil {
		return fmt.Errorf("error setting inputs: %w", err)
	}

	return resourceWebhookrelayBucketUpdate(d, meta)
}

func resourceWebhookrelayBucketUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	if d.HasChanges("name", "description", "websocket_streaming", "ephemeral_webhooks") {
		request := &models.Bucket{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Stream:      d.Get("websocket_streaming").(bool),
			Ephemeral:   d.Get("ephemeral_webhooks").(bool),
		}

		params := buckets.NewPutV1BucketsBucketIDParams().WithBucketID(d.Id()).WithBody(request)

		_, err := client.Buckets.PutV1BucketsBucketID(params)
		if err != nil {
			return fmt.Errorf("error updating bucket: %w", err)
		}
	}

	if d.Get("delete_default_input").(bool) {
		for _, i := range d.Get("default_input").([]interface{}) {
			err := deleteInput(client, d.Id(), i.(map[string]interface{})["id"].(string))
			if err != nil {
				return fmt.Errorf("error deleting default input %w", err)
			}
		}
		d.Set("default_input", nil)
	}

	return nil
}

func resourceWebhookrelayBucketDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	getParams := buckets.NewGetV1BucketsBucketIDParams().WithBucketID(d.Id())
	resp, err := client.Buckets.GetV1BucketsBucketID(getParams)
	if err != nil {
		if _, ok := err.(*buckets.GetV1BucketsBucketIDNotFound); ok {
			log.Printf("[WARN] Removing bucket %s from state because it's gone", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error reading bucket details: %w", err)
	}

	for _, i := range resp.GetPayload().Inputs {
		err := deleteInput(client, d.Id(), i.ID)
		if err != nil {
			return fmt.Errorf("error destroying input %s: %w", i.ID, err)
		}
	}

	delParams := buckets.NewDeleteV1BucketsBucketIDParams().WithBucketID(d.Id())
	_, err = client.Buckets.DeleteV1BucketsBucketID(delParams)
	if err != nil {
		return fmt.Errorf("Error destroying bucket %s: %s", d.Id(), err)
	}

	return nil
}

func resourceWebhookrelayBucketExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(*client.Openapi)

	params := buckets.NewGetV1BucketsBucketIDParams().WithBucketID(d.Id())
	_, err := client.Buckets.GetV1BucketsBucketID(params)
	if err != nil {
		if _, ok := err.(*buckets.GetV1BucketsBucketIDNotFound); ok {
			log.Printf("[WARN] Removing bucket %s because it's gone", d.Id())
			d.SetId("")
			return false, nil
		}
		return false, fmt.Errorf("Error getting bucket: %w", err)
	}

	return true, nil
}
