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

func expandInputs(inputs []map[string]interface{}) []*models.Input {
	result := make([]*models.Input, 0, len(inputs))

	for _, i := range inputs {
		input := &models.Input{
			Name:        i["id"].(string),
			Description: i["description"].(string),
			ID:          i["name"].(string),
		}

		result = append(result, input)
	}

	return result
}
