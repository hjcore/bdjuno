package types

import (
	"fmt"
	"os"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	cmdxapp "github.com/gotabit/comdex/v8/app"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/forbole/juno/v4/node/remote"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/forbole/juno/v4/node/local"

	nodeconfig "github.com/forbole/juno/v4/node/config"

	banksource "github.com/gotabit/gjuno/v3/modules/bank/source"
	localbanksource "github.com/gotabit/gjuno/v3/modules/bank/source/local"
	remotebanksource "github.com/gotabit/gjuno/v3/modules/bank/source/remote"
	distrsource "github.com/gotabit/gjuno/v3/modules/distribution/source"
	localdistrsource "github.com/gotabit/gjuno/v3/modules/distribution/source/local"
	remotedistrsource "github.com/gotabit/gjuno/v3/modules/distribution/source/remote"
	govsource "github.com/gotabit/gjuno/v3/modules/gov/source"
	localgovsource "github.com/gotabit/gjuno/v3/modules/gov/source/local"
	remotegovsource "github.com/gotabit/gjuno/v3/modules/gov/source/remote"
	mintsource "github.com/gotabit/gjuno/v3/modules/mint/source"
	localmintsource "github.com/gotabit/gjuno/v3/modules/mint/source/local"
	remotemintsource "github.com/gotabit/gjuno/v3/modules/mint/source/remote"
	slashingsource "github.com/gotabit/gjuno/v3/modules/slashing/source"
	localslashingsource "github.com/gotabit/gjuno/v3/modules/slashing/source/local"
	remoteslashingsource "github.com/gotabit/gjuno/v3/modules/slashing/source/remote"
	stakingsource "github.com/gotabit/gjuno/v3/modules/staking/source"
	localstakingsource "github.com/gotabit/gjuno/v3/modules/staking/source/local"
	remotestakingsource "github.com/gotabit/gjuno/v3/modules/staking/source/remote"
	wasmsource "github.com/gotabit/gjuno/v3/modules/wasm/source"
	localwasmsource "github.com/gotabit/gjuno/v3/modules/wasm/source/local"
	remotewasmsource "github.com/gotabit/gjuno/v3/modules/wasm/source/remote"
)

type Sources struct {
	BankSource     banksource.Source
	DistrSource    distrsource.Source
	GovSource      govsource.Source
	MintSource     mintsource.Source
	SlashingSource slashingsource.Source
	StakingSource  stakingsource.Source
	WasmSource     wasmsource.Source
}

func BuildSources(nodeCfg nodeconfig.Config, encodingConfig *params.EncodingConfig) (*Sources, error) {
	switch cfg := nodeCfg.Details.(type) {
	case *remote.Details:
		return buildRemoteSources(cfg)
	case *local.Details:
		return buildLocalSources(cfg, encodingConfig)

	default:
		return nil, fmt.Errorf("invalid configuration type: %T", cfg)
	}
}

func buildLocalSources(cfg *local.Details, encodingConfig *params.EncodingConfig) (*Sources, error) {
	source, err := local.NewSource(cfg.Home, encodingConfig)
	if err != nil {
		return nil, err
	}

	cmdxApp := cmdxapp.New(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), source.StoreDB, nil, true, map[int64]bool{},
		cfg.Home, 0, cmdxapp.MakeEncodingConfig(), simapp.EmptyAppOptions{}, nil, nil,
	)

	sources := &Sources{
		BankSource:     localbanksource.NewSource(source, banktypes.QueryServer(cmdxApp.BankKeeper)),
		DistrSource:    localdistrsource.NewSource(source, distrtypes.QueryServer(cmdxApp.DistrKeeper)),
		GovSource:      localgovsource.NewSource(source, govtypes.QueryServer(cmdxApp.GovKeeper)),
		MintSource:     localmintsource.NewSource(source, minttypes.QueryServer(cmdxApp.MintKeeper)),
		SlashingSource: localslashingsource.NewSource(source, slashingtypes.QueryServer(cmdxApp.SlashingKeeper)),
		StakingSource:  localstakingsource.NewSource(source, stakingkeeper.Querier{Keeper: cmdxApp.StakingKeeper}),
		WasmSource:     localwasmsource.NewSource(source, wasmkeeper.Querier(&cmdxApp.WasmKeeper)),
	}

	// Mount and initialize the stores
	err = source.MountKVStores(cmdxApp, "keys")
	if err != nil {
		return nil, err
	}

	err = source.MountTransientStores(cmdxApp, "tkeys")
	if err != nil {
		return nil, err
	}

	err = source.MountMemoryStores(cmdxApp, "memKeys")
	if err != nil {
		return nil, err
	}

	err = source.InitStores()
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func buildRemoteSources(cfg *remote.Details) (*Sources, error) {
	source, err := remote.NewSource(cfg.GRPC)
	if err != nil {
		return nil, fmt.Errorf("error while creating remote source: %s", err)
	}

	return &Sources{
		BankSource:     remotebanksource.NewSource(source, banktypes.NewQueryClient(source.GrpcConn)),
		DistrSource:    remotedistrsource.NewSource(source, distrtypes.NewQueryClient(source.GrpcConn)),
		GovSource:      remotegovsource.NewSource(source, govtypes.NewQueryClient(source.GrpcConn)),
		MintSource:     remotemintsource.NewSource(source, minttypes.NewQueryClient(source.GrpcConn)),
		SlashingSource: remoteslashingsource.NewSource(source, slashingtypes.NewQueryClient(source.GrpcConn)),
		StakingSource:  remotestakingsource.NewSource(source, stakingtypes.NewQueryClient(source.GrpcConn)),
		WasmSource:     remotewasmsource.NewSource(source, wasmtypes.NewQueryClient(source.GrpcConn)),
	}, nil
}
