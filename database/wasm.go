package database

import (
	"github.com/CosmWasm/wasmd/x/wasm"
)

func (db *Db) SaveWasmExec(params wasm.MsgExecuteContract, height int64) error {
	paramsBz, err := params.GetMsg().MarshalJSON()
	if err != nil {
		return err
	}
	stmt := `
	INSERT INTO wasm_exec (height, contract_address, params, sender) 
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
        height = excluded.height
WHERE wasm_exec.height <= excluded.height
	`

	_, err = db.Sql.Exec(stmt, height, params.Contract, string(paramsBz), params.Sender)
	return err
}
