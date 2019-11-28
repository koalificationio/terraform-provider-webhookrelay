package webhookrelay

import (
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

// TODO: enable updating inputs
// func expandInputs(inputs []map[string]interface{}) []*models.Input {
// 	result := make([]*models.Input, 0, len(inputs))

// 	for _, i := range inputs {
// 		input := &models.Input{
// 			Name:        i["id"].(string),
// 			Description: i["description"].(string),
// 			ID:          i["name"].(string),
// 		}

// 		result = append(result, input)
// 	}

// 	return result
// }

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
