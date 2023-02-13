package database

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	dbtypes "github.com/forbole/bdjuno/v3/database/types"
	"github.com/lib/pq"
)

type GidRegister struct {
	Register struct {
		Duration  int    `json:"duration"`
		Name      string `json:"name"`
		TokenInfo struct {
			TokenID   string `json:"token_id"`
			Owner     string `json:"owner"`
			TokenURI  string `json:"token_uri"`
			Extension string `json:"extension"`
		} `json:"token_info"`
	} `json:"register"`
}

func (db *Db) SaveWasmExec(params wasm.MsgExecuteContract, height int64) error {
	reg := new(GidRegister)
	if err := json.Unmarshal(params.Msg.Bytes(), reg); err != nil {
		return err
	}

	if reg.Register.Name == "" {
		return nil
	}

	stmt := `
	INSERT INTO gid_wasm_exec (height, contract_address, reg_name, duration, sender, coins) 
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (one_row_id) DO UPDATE 
    SET height = excluded.height
	WHERE wasm_exec.height <= excluded.height
	`

	_, err := db.Sql.Exec(stmt, height, params.Contract, reg.Register.Name, reg.Register.Duration, params.Sender, pq.Array(dbtypes.NewDbCoins(params.Funds)))
	return err
}
