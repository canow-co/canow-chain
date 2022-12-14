package cli

import (
	"canow-chain/tests/integration/network"
	"encoding/json"
	"strings"

	"github.com/canow-co/cheqd-node/tests/integration/helpers"
	"github.com/canow-co/cheqd-node/x/did/client/cli"
	"github.com/canow-co/cheqd-node/x/did/types"
	resourcecli "github.com/canow-co/cheqd-node/x/resource/client/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var CLI_TX_PARAMS = []string{
	"--chain-id", network.CHAIN_ID,
	"--keyring-backend", KEYRING_BACKEND,
	"--output", OUTPUT_FORMAT,
	"--gas", GAS,
	"--gas-adjustment", GAS_ADJUSTMENT,
	"--gas-prices", GAS_PRICES,
	"--yes",
}

func Tx(module, tx, from string, txArgs ...string) (sdk.TxResponse, error) {
	args := []string{"tx", module, tx}

	// Common params
	args = append(args, CLI_TX_PARAMS...)

	// Cosmos account
	args = append(args, "--from", from)

	// Other args
	args = append(args, txArgs...)

	output, err := Exec(args...)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// Skip 'gas estimate: xxx' string
	output = strings.Split(output, "\n")[1]

	var resp sdk.TxResponse

	err = helpers.Codec.UnmarshalJSON([]byte(output), &resp)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	return resp, nil
}

func TransferToken(from string, to string, amount string) (sdk.TxResponse, error) {
	return Tx("bank", "send", from, from, to, amount)
}

func CreateDidDoc(tmpDit string, payload types.MsgCreateDidDocPayload, signInputs []cli.SignInput, from string) (sdk.TxResponse, error) {
	// Payload
	payloadJson, err := helpers.Codec.MarshalJSON(&payload)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadWithSignInputs := cli.PayloadWithSignInputs{
		Payload:    payloadJson,
		SignInputs: signInputs,
	}

	payloadWithSignInputsJson, err := json.Marshal(&payloadWithSignInputs)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadFile := helpers.MustWriteTmpFile(tmpDit, []byte(payloadWithSignInputsJson))

	return Tx("cheqd", "create-did", from, payloadFile)
}

func UpdateDidDoc(tmpDir string, payload types.MsgUpdateDidDocPayload, signInputs []cli.SignInput, from string) (sdk.TxResponse, error) {
	// Payload
	payloadJson, err := helpers.Codec.MarshalJSON(&payload)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadWithSignInputs := cli.PayloadWithSignInputs{
		Payload:    payloadJson,
		SignInputs: signInputs,
	}

	payloadWithSignInputsJson, err := json.Marshal(&payloadWithSignInputs)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadFile := helpers.MustWriteTmpFile(tmpDir, []byte(payloadWithSignInputsJson))

	return Tx("cheqd", "update-did", from, payloadFile)
}

func DeactivateDidDoc(tmpDir string, payload types.MsgUpdateDidDocPayload, signInputs []cli.SignInput, from string) (sdk.TxResponse, error) {
	// Payload
	payloadJson, err := helpers.Codec.MarshalJSON(&payload)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadWithSignInputs := cli.PayloadWithSignInputs{
		Payload:    payloadJson,
		SignInputs: signInputs,
	}

	payloadWithSignInputsJson, err := json.Marshal(&payloadWithSignInputs)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadFile := helpers.MustWriteTmpFile(tmpDir, []byte(payloadWithSignInputsJson))

	return Tx("cheqd", "deactivate-did", from, payloadFile)
}

func CreateResource(tmpDir string, options resourcecli.CreateResourceOptions, signInputs []cli.SignInput, from string) (sdk.TxResponse, error) {
	// Payload
	payloadJson, err := json.Marshal(&options)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadWithSignInputs := cli.PayloadWithSignInputs{
		Payload:    payloadJson,
		SignInputs: signInputs,
	}

	payloadWithSignInputsJson, err := json.Marshal(&payloadWithSignInputs)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadFile := helpers.MustWriteTmpFile("", []byte(payloadWithSignInputsJson))

	return Tx("resource", "create-resource", from, payloadFile)
}
