package cmd

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client"

	cheqdcmd "github.com/canow-co/cheqd-node/cmd/cheqd-noded/cmd"
	cosmcfg "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/spf13/cobra"
	tmcfg "github.com/tendermint/tendermint/config"
)

func ExtendInit(initCmd *cobra.Command) *cobra.Command {
	baseRunE := initCmd.RunE

	initCmd.RunE = func(cmd *cobra.Command, args []string) error {
		err := baseRunE(cmd, args)
		if err != nil {
			return err
		}

		err = applyConfigDefaults(cmd)
		if err != nil {
			return err
		}

		return nil
	}

	return initCmd
}

func applyConfigDefaults(cmd *cobra.Command) error {
	clientCtx := client.GetClientContextFromCmd(cmd)

	err := cheqdcmd.UpdateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
		config.Consensus.CreateEmptyBlocks = false
		config.Consensus.CreateEmptyBlocksInterval = time.Duration(600*time.Second)
		config.FastSync.Version = "v0"
		config.LogFormat = "json"
		config.LogLevel = "error"
		config.P2P.SendRate = 20000000
		config.P2P.RecvRate = 20000000
		config.P2P.MaxPacketMsgPayloadSize = 10240
		config.Instrumentation.Prometheus = true
		config.Instrumentation.MaxOpenConnections = 3
		config.Instrumentation.Namespace = "tendermint"

		

		// Workaround for Tendermint's bug
		config.Storage = tmcfg.DefaultStorageConfig()
	})
	if err != nil {
		return err
	}

	err = cheqdcmd.UpdateCosmConfig(clientCtx.HomeDir, func(config *cosmcfg.Config) {
		config.BaseConfig.MinGasPrices = "50zarx"
		config.BaseConfig.Pruning = "nothing"
		config.StateSync.SnapshotInterval = 100
		config.StateSync.SnapshotKeepRecent = 2
	})
	if err != nil {
		return err
	}

	return nil
}
