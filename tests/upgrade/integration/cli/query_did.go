package cli

import (
	integrationhelpers "github.com/canow-co/canow-chain/tests/integration/helpers"
	didtypesv2 "github.com/canow-co/cheqd-node/x/did/types"
)

func QueryDid(did string, container string) (didtypesv2.QueryDidDocResponse, error) {
	res, err := Query(container, CliBinaryName, "cheqd", "did-document", did)
	if err != nil {
		return didtypesv2.QueryDidDocResponse{}, err
	}

	var resp didtypesv2.QueryDidDocResponse
	err = integrationhelpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return didtypesv2.QueryDidDocResponse{}, err
	}

	return resp, nil
}
