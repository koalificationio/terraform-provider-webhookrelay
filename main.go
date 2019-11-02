package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/koalificationio/terraform-provider-webhookrelay/webhookrelay"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: webhookrelay.Provider})
}
