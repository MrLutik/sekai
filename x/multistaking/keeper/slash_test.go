package keeper_test

import (
	"github.com/KiraCore/sekai/x/multistaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestSlashStakingPool() {
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	// set pool
	delCoins := sdk.Coins{sdk.NewInt64Coin("ukex", 1000000)}
	pool := types.StakingPool{
		Id:                 1,
		Validator:          sdk.ValAddress(addr).String(),
		Enabled:            true,
		Slashed:            0,
		TotalStakingTokens: delCoins,
		TotalShareTokens:   sdk.Coins{sdk.NewInt64Coin("v1/ukex", 1000000)},
		TotalRewards:       sdk.Coins(nil),
	}
	suite.app.MultiStakingKeeper.SetStakingPool(suite.ctx, pool)

	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, delCoins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, delCoins)
	suite.Require().NoError(err)

	// slash 10%
	suite.app.MultiStakingKeeper.SlashStakingPool(suite.ctx, pool.Validator, 10)

	// check pool after slash
	p, found := suite.app.MultiStakingKeeper.GetStakingPoolByValidator(suite.ctx, pool.Validator)
	suite.Require().True(found)
	suite.Require().Equal(p, types.StakingPool{
		Id:                 1,
		Validator:          sdk.ValAddress(addr).String(),
		Enabled:            false,
		Slashed:            10,
		TotalStakingTokens: sdk.Coins{sdk.NewInt64Coin("ukex", 900000)},
		TotalShareTokens:   sdk.Coins{sdk.NewInt64Coin("v1/ukex", 1000000)},
		TotalRewards:       sdk.Coins(nil),
		Commission:         sdk.ZeroDec(),
	})
}
