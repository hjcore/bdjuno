package wasm

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"

	"github.com/CosmWasm/wasmd/x/wasm"
)

var (
	GID_CONTRACT_ADDRESS = "GID_CONTRACT_ADDRESS"
)

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch msg := msg.(type) {
	case *wasm.MsgExecuteContract:
		contractAddress := msg.Contract
		switch contractAddress {
		case GID_CONTRACT_ADDRESS:
			return m.HandleMsgGidWasmExec(tx, msg, tx.Height)
		default:
			return fmt.Errorf("unsupport contract!")
		}
	default:
		return nil
	}
}

func (m *Module) HandleMsgGidWasmExec(tx *juno.Tx, params *wasm.MsgExecuteContract, height int64) error {
	return m.db.SaveWasmExec(*params, height)
}
