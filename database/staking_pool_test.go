package database_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	dbtypes "github.com/gotabit/gjuno/v3/database/types"
	"github.com/gotabit/gjuno/v3/types"
)

func (suite *DbTestSuite) TestGatabitDb_SaveStakingPool() {
	// Save the data
	original := types.NewPool(sdk.NewInt(50), sdk.NewInt(100), sdk.NewInt(5), sdk.NewInt(1), 10)
	err := suite.database.SaveStakingPool(original)
	suite.Require().NoError(err)

	// Verify the data
	expected := dbtypes.NewStakingPoolRow(50, 100, 5, 1, 10)

	var rows []dbtypes.StakingPoolRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM staking_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected))

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating using a lower height
	pool := types.NewPool(sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1), 8)
	err = suite.database.SaveStakingPool(pool)
	suite.Require().NoError(err)

	// Verify the data
	rows = []dbtypes.StakingPoolRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM staking_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with a lower height should not modify the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with the same height
	pool = types.NewPool(sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1), 10)
	err = suite.database.SaveStakingPool(pool)
	suite.Require().NoError(err)

	// Verify the data
	expected = dbtypes.NewStakingPoolRow(1, 1, 1, 1, 10)

	rows = []dbtypes.StakingPoolRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM staking_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with a lower height should not modify the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a higher height
	pool = types.NewPool(sdk.NewInt(1000000), sdk.NewInt(1000000), sdk.NewInt(20), sdk.NewInt(15), 20)
	err = suite.database.SaveStakingPool(pool)
	suite.Require().NoError(err)

	// Verify the data
	expected = dbtypes.NewStakingPoolRow(1000000, 1000000, 20, 15, 20)

	rows = []dbtypes.StakingPoolRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM staking_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with a lower height should not modify the data")
}
