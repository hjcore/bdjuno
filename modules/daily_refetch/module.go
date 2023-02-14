package daily_refetch

import (
	"github.com/forbole/juno/v4/node"

	gjunodb "github.com/gotabit/gjuno/v3/database"

	"github.com/forbole/juno/v4/modules"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

type Module struct {
	node     node.Node
	database *gjunodb.Db
}

// NewModule builds a new Module instance
func NewModule(
	node node.Node,
	database *gjunodb.Db,
) *Module {
	return &Module{
		node:     node,
		database: database,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "daily refetch"
}
