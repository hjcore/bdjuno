package types

import "time"

type GidWasmExec struct {
	OneRowID        bool      `db:"one_row_id"`
	Height          int64     `db:"height"`
	ContractAddress time.Time `db:"contract_address"`
	RegName         string    `db:"reg_name"`
	Duration        string    `db:"duration"`
	Sender          string    `db:"sender"`
	Coins           *DbCoins  `db:"coins"`
}
