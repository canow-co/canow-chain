package cli

import (
	integrationhelpers "github.com/canow-co/canow-chain/tests/integration/helpers"
	resourcetypesv2 "github.com/canow-co/cheqd-node/x/resource/types"
)

func QueryResource(collectionID string, resourceID string, container string) (resourcetypesv2.QueryResourceResponse, error) {
	res, err := Query(container, CliBinaryName, "resource", "specific-resource", collectionID, resourceID)
	if err != nil {
		return resourcetypesv2.QueryResourceResponse{}, err
	}

	var resp resourcetypesv2.QueryResourceResponse
	err = integrationhelpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return resourcetypesv2.QueryResourceResponse{}, err
	}

	return resp, nil
}
