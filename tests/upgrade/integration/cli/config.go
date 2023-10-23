package cli

import (
	integrationcli "github.com/canow-co/canow-chain/tests/integration/cli"
	integrationnetwork "github.com/canow-co/canow-chain/tests/integration/network"
)

const (
	CliBinaryName = integrationcli.CliBinaryName
	Green         = integrationcli.Green
	Purple        = integrationcli.Purple
)

const (
	KeyringBackend = integrationcli.KeyringBackend
	OutputFormat   = integrationcli.OutputFormat
	Gas            = integrationcli.Gas
	GasAdjustment  = integrationcli.GasAdjustment
	GasPrices      = integrationcli.GasPrices

	BootstrapPeriod            = 20
	BootstrapHeight            = 1
	VotingPeriod         int64 = 10
	ExpectedBlockSeconds int64 = 1
	ExtraBlocks          int64 = 5
	UpgradeName                = "v0.3.0"
	DepositAmount              = "10000000zarx"
	NetworkConfigDir           = "network-config"
	KeyringDir                 = "keyring-test"
)

var (
	TXParams = []string{
		"--chain-id", integrationnetwork.ChainID,
		"--keyring-backend", KeyringBackend,
		"--output", OutputFormat,
		"--yes",
	}
	GasParams = []string{
		"--gas", Gas,
		"--gas-adjustment", GasAdjustment,
		"--gas-prices", GasPrices,
	}
	QueryParamsConst = []string{
		"--chain-id", integrationnetwork.ChainID,
		"--output", OutputFormat,
	}
)
