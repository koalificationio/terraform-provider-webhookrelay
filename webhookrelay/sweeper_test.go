package webhookrelay

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/koalificationio/go-webhookrelay/pkg/client"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// sharedConfigForRegion returns a common config setup needed for the sweeper
// functions for a given region
func sharedConfigForRegion(region string) (interface{}, error) {
	if v := os.Getenv("RELAY_KEY"); v == "" {
		return nil, fmt.Errorf("$RELAY_KEY must be set")
	}
	if v := os.Getenv("RELAY_SECRET"); v == "" {
		return nil, fmt.Errorf("$RELAY_SECRET must be set")
	}

	config := &client.Config{
		APIKey:    os.Getenv("RELAY_KEY"),
		APISecret: os.Getenv("RELAY_SECRET"),
	}

	client := client.New(config)

	return client, nil
}
