package types

import "time"

type WasmExec struct {
	OneRowID        bool      `db:"one_row_id"`
	Height          int64     `db:"height"`
	ContractAddress time.Time `db:"contract_address"`
	Params          string    `db:"params"`
	Sender          string    `db:"sender"`
}
