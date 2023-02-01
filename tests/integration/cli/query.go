package cli

import (
	"github.com/canow-co/canow-chain/tests/integration/helpers"
	"github.com/canow-co/canow-chain/tests/integration/network"

	didtypes "github.com/canow-co/cheqd-node/x/did/types"
	resourcetypes "github.com/canow-co/cheqd-node/x/resource/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
)

var CLIQueryParams = []string{
	"--chain-id", network.ChainID,
	"--output", OutputFormat,
}

func Query(module, query string, queryArgs ...string) (string, error) {
	args := []string{"query", module, query}

	// Common params
	args = append(args, CLIQueryParams...)

	// Other args
	args = append(args, queryArgs...)

	return Exec(args...)
}

func QueryBalance(address, denom string) (sdk.Coin, error) {
	res, err := Query("bank", "balances", address, "--denom", denom)
	if err != nil {
		return sdk.Coin{}, err
	}

	var resp sdk.Coin
	err = helpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return sdk.Coin{}, err
	}

	return resp, nil
}

func QueryParams(subspace, key string) (paramproposal.ParamChange, error) {
	res, err := Query("params", "subspace", subspace, key)
	if err != nil {
		return paramproposal.ParamChange{}, err
	}

	var resp paramproposal.ParamChange
	err = helpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return paramproposal.ParamChange{}, err
	}

	return resp, nil
}

func QueryGetBalances(address string) (banktypes.QueryAllBalancesResponse, error) {
	res, err := Query("bank", "balances", address)
	if err != nil {
		return banktypes.QueryAllBalancesResponse{}, err
	}

	var resp banktypes.QueryAllBalancesResponse
	err = helpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return banktypes.QueryAllBalancesResponse{}, err
	}

	return resp, err
}

func QueryDidDoc(did string) (didtypes.QueryDidDocResponse, error) {
	res, err := Query("cheqd", "did-document", did)
	if err != nil {
		return didtypes.QueryDidDocResponse{}, err
	}

	var resp didtypes.QueryDidDocResponse
	err = helpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return didtypes.QueryDidDocResponse{}, err
	}

	return resp, nil
}

func QueryResource(collectionID string, resourceID string) (resourcetypes.QueryResourceResponse, error) {
	res, err := Query("resource", "specific-resource", collectionID, resourceID)
	if err != nil {
		return resourcetypes.QueryResourceResponse{}, err
	}

	var resp resourcetypes.QueryResourceResponse
	err = helpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return resourcetypes.QueryResourceResponse{}, err
	}

	return resp, nil
}

func QueryResourceMetadata(collectionID string, resourceID string) (resourcetypes.QueryResourceMetadataResponse, error) {
	res, err := Query("resource", "metadata", collectionID, resourceID)
	if err != nil {
		return resourcetypes.QueryResourceMetadataResponse{}, err
	}

	var resp resourcetypes.QueryResourceMetadataResponse
	err = helpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return resourcetypes.QueryResourceMetadataResponse{}, err
	}

	return resp, nil
}

func QueryResourceCollection(collectionID string) (resourcetypes.QueryCollectionResourcesResponse, error) {
	res, err := Query("resource", "collection-metadata", collectionID)
	if err != nil {
		return resourcetypes.QueryCollectionResourcesResponse{}, err
	}

	var resp resourcetypes.QueryCollectionResourcesResponse
	err = helpers.Codec.UnmarshalJSON([]byte(res), &resp)
	if err != nil {
		return resourcetypes.QueryCollectionResourcesResponse{}, err
	}

	return resp, nil
}
