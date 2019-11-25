package webhookrelay

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/tokens"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

func resourceWebhookrelayToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebhookrelayTokenCreate,
		Read:   resourceWebhookrelayTokenRead,
		Update: resourceWebhookrelayTokenUpdate,
		Delete: resourceWebhookrelayTokenDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"api_access": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled",
				}, false),
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"scopes": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buckets": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tunnels": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceWebhookrelayTokenCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	request := &models.Token{
		Description: d.Get("description").(string),
		APIAccess:   d.Get("api_access").(string),
	}

	if v, ok := d.Get("scopes").([]interface{}); ok && len(v) > 0 {
		request.Scopes = expandScopes(v)
	}

	params := tokens.NewPostV1TokensParams().WithBody(request)

	resp, err := client.Tokens.PostV1Tokens(params)
	if err != nil {
		return fmt.Errorf("failed creating token: %w", err)
	}

	token := resp.GetPayload()

	d.SetId(token.Key)
	d.Set("key", token.Key)
	d.Set("secret", token.Secret)

	return nil
}

func resourceWebhookrelayTokenRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	params := tokens.NewGetV1TokensParams()
	resp, err := client.Tokens.GetV1Tokens(params)
	if err != nil {
		return fmt.Errorf("failed getting tokens: %w", err)
	}

	var token *models.Token
	for _, t := range resp.GetPayload() {
		if d.Id() == t.ID {
			token = t
		}
	}

	if token == nil {
		log.Printf("[WARN] Removing token %s from state because it's gone", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("description", token.Description)
	d.Set("api_access", token.APIAccess)
	if len(token.Scopes.Buckets) > 0 || len(token.Scopes.Tunnels) > 0 {
		if err := d.Set("scopes", flattenScopes(token.Scopes)); err != nil {
			return fmt.Errorf("error setting scopes: %w", err)
		}
	}

	return nil
}

func resourceWebhookrelayTokenUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	if d.HasChange("description") || d.HasChange("api_access") || d.HasChange("scopes") {
		request := &models.Token{
			Description: d.Get("description").(string),
			APIAccess:   d.Get("api_access").(string),
		}

		if v, ok := d.Get("scopes").([]interface{}); ok && len(v) > 0 {
			request.Scopes = expandScopes(v)
		}

		params := tokens.NewPutV1TokensTokenIDParams()
		params.SetTokenID(d.Id())
		params.SetBody(request)

		_, err := client.Tokens.PutV1TokensTokenID(params)
		if err != nil {
			return fmt.Errorf("failed updating token: %w", err)
		}
	}

	return nil
}

func resourceWebhookrelayTokenDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Openapi)

	params := tokens.NewDeleteV1TokensTokenIDParams().WithTokenID(d.Id())

	_, err := client.Tokens.DeleteV1TokensTokenID(params)
	if err != nil {
		return fmt.Errorf("failed deleting token: %w", err)
	}

	return nil
}
