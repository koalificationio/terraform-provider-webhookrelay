package webhookrelay

import (
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/client/inputs"
)

func deleteInput(client *client.Openapi, bucketID, inputID string) error {
	inputParams := inputs.NewDeleteV1BucketsBucketIDInputsInputIDParams().
		WithInputID(inputID).WithBucketID(bucketID)

	_, err := client.Inputs.DeleteV1BucketsBucketIDInputsInputID(inputParams)

	return err
}
