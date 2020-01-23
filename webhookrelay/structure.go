package webhookrelay

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

func flattenInputs(inputs []*models.Input) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(inputs))

	for _, i := range inputs {
		input := map[string]interface{}{
			"name":        i.Name,
			"description": i.Description,
			"id":          i.ID,
		}
		result = append(result, input)
	}
	return result
}

func flattenScopes(scopes *models.TokenScopes) []map[string]interface{} {
	result := make([]map[string]interface{}, 1)

	s := make(map[string]interface{})

	s["buckets"] = scopes.Buckets
	s["tunnels"] = scopes.Tunnels

	result[0] = s

	return result
}

func expandScopes(scopes []interface{}) *models.TokenScopes {
	var buckets, tunnels []string

	for _, b := range scopes[0].(map[string]interface{})["buckets"].([]interface{}) {
		buckets = append(buckets, b.(string))
	}

	for _, t := range scopes[0].(map[string]interface{})["tunnels"].([]interface{}) {
		tunnels = append(tunnels, t.(string))
	}

	result := &models.TokenScopes{
		Buckets: buckets,
		Tunnels: tunnels,
	}

	return result
}

func flattenBucketAuth(auth *models.BucketAuth) []map[string]interface{} {
	result := make([]map[string]interface{}, 1)

	config := make(map[string]interface{})

	config["type"] = auth.Type

	switch auth.Type {
	case "none":
		return nil
	case "basic":
		config["username"] = auth.Username
		config["password"] = auth.Password
	case "token":
		config["token"] = auth.Token
	}

	result[0] = config

	return result
}

func expandBucketAuth(config []interface{}) *models.BucketAuth {
	result := &models.BucketAuth{
		// set default type to disabled
		Type: "none",
	}

	if len(config) == 0 || config[0] == nil {
		return result
	}

	auth := config[0].(map[string]interface{})

	result.Type = auth["type"].(string)

	switch result.Type {
	case "basic":
		result.Username = auth["username"].(string)
		result.Password = auth["password"].(string)
	case "token":
		result.Token = auth["token"].(string)
	}

	return result
}

func flattenHeaders(headers models.Headers) map[string]interface{} {
	config := make(map[string]interface{})

	m := interface{}(headers).(map[string]interface{})

	for h, v := range m {
		a := v.([]interface{})
		b := make([]string, len(a))
		for i := range v.([]interface{}) {
			b[i] = a[i].(string)
		}
		config[h] = strings.Join(b, ";")
	}

	return config
}

func expandHeaders(config map[string]interface{}) models.Headers {
	headers := make(map[string][]string)

	for h, v := range config {
		headers[h] = []string{v.(string)}
	}

	return interface{}(headers).(models.Headers)
}

func flattenOutputRules(rules *models.Rules) (string, error) {
	b, err := rules.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("Error parsing rules: %v", err)
	}

	config, err := structure.NormalizeJsonString(string(b))
	if err != nil {
		return "", fmt.Errorf("Rules contain an invalid JSON: %v", err)
	}

	return config, nil
}

func expandOutputRules(config string) (*models.Rules, error) {
	rules := models.Rules{}

	if config == "" {
		return nil, nil
	}

	err := rules.UnmarshalBinary([]byte(config))
	if err != nil {
		return nil, fmt.Errorf("Error parsing rules config: %v", err)
	}

	return &rules, nil
}
