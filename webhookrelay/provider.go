package webhookrelay

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/koalificationio/go-webhookrelay/pkg/client"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RELAY_KEY", nil),
				Description: "The Webhookrelay API key.",
			},
			"api_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RELAY_SECRET", nil),
				Description: "The Webhookrelay API secret.",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Webhook API host. Default is my.webhookrelay.com.",
			},
			"base_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "API basepath. Default is /.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"webhookrelay_bucket":   resourceWebhookrelayBucket(),
			"webhookrelay_input":    resourceWebhookrelayInput(),
			"webhookrelay_function": resourceWebhookrelayFunction(),
			"webhookrelay_output":   resourceWebhookrelayOutput(),
			"webhookrelay_token":    resourceWebhookrelayToken(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &client.Config{
		APIKey:    d.Get("api_key").(string),
		APISecret: d.Get("api_secret").(string),
		Host:      d.Get("host").(string),
		BasePath:  d.Get("base_path").(string),
	}

	client := client.New(config)

	return client, nil
}
