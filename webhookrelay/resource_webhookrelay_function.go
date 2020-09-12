package webhookrelay

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/functions"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

func resourceWebhookrelayFunction() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebhookrelayFunctionCreate,
		Read:   resourceWebhookrelayFunctionRead,
		Update: resourceWebhookrelayFunctionUpdate,
		Delete: resourceWebhookrelayFunctionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"payload": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsBase64,
			},
			"driver": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "lua",
				ValidateFunc: validation.StringInSlice([]string{"lua", "wasi"}, false),
			},
			"config": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceWebhookrelayFunctionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	request := &models.FunctionRequest{
		Driver:  d.Get("driver").(string),
		Name:    d.Get("name").(string),
		Payload: d.Get("payload").(string),
	}

	params := functions.NewPostV1FunctionsParams().WithBody(request)

	resp, err := client.Functions.PostV1Functions(params)
	if err != nil {
		return fmt.Errorf("error creating function: %w", err)
	}
	d.SetId(resp.GetPayload().ID)

	return resourceWebhookrelayFunctionUpdate(d, meta)
}

func resourceWebhookrelayFunctionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	params := functions.NewGetV1FunctionsFunctionIDParams().WithFunctionID(d.Id())
	resp, err := client.Functions.GetV1FunctionsFunctionID(params)
	if err != nil {
		if _, ok := err.(*functions.GetV1FunctionsFunctionIDNotFound); ok {
			log.Printf("[WARN] Removing function %s from state because it's gone", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error reading function details: %w", err)
	}

	function := resp.GetPayload()
	d.Set("name", function.Name)
	d.Set("payload", function.Payload)

	configParams := functions.NewGetV1FunctionsFunctionIDConfigParams().WithFunctionID(d.Id())
	configResp, err := client.Functions.GetV1FunctionsFunctionIDConfig(configParams)
	if err != nil {
		return fmt.Errorf("error reading function details: %w", err)
	}

	if err := d.Set("config", flattenFunctionConfig(configResp.Payload)); err != nil {
		return fmt.Errorf("error parsing function config: %w", err)
	}

	return nil
}

func resourceWebhookrelayFunctionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	if d.HasChanges("name", "payload", "driver") {
		request := &models.FunctionRequest{
			Driver:  d.Get("driver").(string),
			Name:    d.Get("name").(string),
			Payload: d.Get("payload").(string),
		}

		params := functions.NewPutV1FunctionsFunctionIDParams().WithFunctionID(d.Id()).WithBody(request)

		_, err := client.Functions.PutV1FunctionsFunctionID(params)
		if err != nil {
			return fmt.Errorf("error updating function: %w", err)
		}
	}

	if d.HasChange("config") {
		o, n := d.GetChange("config")
		if o == nil {
			o = make(map[string]interface{})
		}
		if n == nil {
			n = make(map[string]interface{})
		}

		oc := o.(map[string]interface{})
		nc := n.(map[string]interface{})

		for k := range oc {
			if _, ok := nc[k]; !ok {
				params := functions.DeleteV1FunctionsFunctionIDConfigConfigKeyParams{
					ConfigKey:  k,
					FunctionID: d.Id(),
				}

				_, err := client.Functions.DeleteV1FunctionsFunctionIDConfigConfigKey(&params)
				if err != nil {
					return fmt.Errorf("error deleting function config key: %w", err)
				}
			}
		}

		for k, v := range nc {
			config := models.FunctionConfig{
				Key:   k,
				Value: v.(string),
			}
			params := functions.NewPutV1FunctionsFunctionIDConfigParams().WithFunctionID(d.Id()).WithBody(&config)

			_, err := client.Functions.PutV1FunctionsFunctionIDConfig(params)
			if err != nil {
				return fmt.Errorf("error updating function config: %w", err)
			}
		}
	}

	return resourceWebhookrelayFunctionRead(d, meta)
}

func resourceWebhookrelayFunctionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	getParams := functions.NewGetV1FunctionsFunctionIDParams().WithFunctionID(d.Id())
	_, err := client.Functions.GetV1FunctionsFunctionID(getParams)
	if err != nil {
		if _, ok := err.(*functions.GetV1FunctionsFunctionIDNotFound); ok {
			log.Printf("[WARN] Removing function %s from state because it's gone", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error reading function details: %w", err)
	}

	delParams := functions.NewDeleteV1FunctionsFunctionIDParams().WithFunctionID(d.Id())
	_, err = client.Functions.DeleteV1FunctionsFunctionID(delParams)
	if err != nil {
		return fmt.Errorf("Error destroying function %s: %s", d.Id(), err)
	}

	return nil
}
