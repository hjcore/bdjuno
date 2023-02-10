package types

import "time"

type WasmExec struct {
	Id              int       `db:"id"`
	Height          int64     `db:"height"`
	ContractAddress time.Time `db:"contract_address"`
	Params          string    `db:"params"`
	Sender          string    `db:"sender"`
}
