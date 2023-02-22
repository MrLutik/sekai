package keeper_test

import (
	recoverytypes "github.com/KiraCore/sekai/x/recovery/types"
	stakingtypes "github.com/KiraCore/sekai/x/staking/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestAllocateTokensToRecoveryTokenValidator() {
	suite.SetupTest()
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	addr2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	valAddr := sdk.ValAddress(addr1)
	pubkeys := simapp.CreateTestPubKeys(1)
	pubKey := pubkeys[0]
	consAddr := sdk.ConsAddress(pubKey.Address())
	val, err := stakingtypes.NewValidator(valAddr, pubKey)
	suite.Require().NoError(err)

	coins := sdk.Coins{sdk.NewInt64Coin("ukex", 10000000)}

	suite.app.CustomStakingKeeper.AddValidator(suite.ctx, val)
	suite.app.RecoveryKeeper.SetRecoveryToken(suite.ctx, recoverytypes.RecoveryToken{
		Address:          addr1.String(),
		Token:            "rr/validator1",
		RrSupply:         sdk.NewInt(1000_000),
		UnderlyingTokens: coins,
	})

	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr2, coins)
	suite.Require().NoError(err)

	suite.app.DistrKeeper.AllocateTokens(suite.ctx, 10, 10, consAddr, []abci.VoteInfo{})
	balance := suite.app.BankKeeper.GetAllBalances(suite.ctx, addr1)
	suite.Require().Equal(balance, coins)

	// recoveryToken, err := suite.app.RecoveryKeeper.GetRecoveryToken(suite.ctx, addr1.String())
	// suite.Require().NoError(err)
	// suite.Require().NotEqual(coins.String(), sdk.Coins(recoveryToken.UnderlyingTokens).String())
}

func (suite *KeeperTestSuite) TestAllocateTokens() {
	suite.SetupTest()
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	addr2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	consAddr := sdk.ConsAddress(addr1)
	coins := sdk.Coins{sdk.NewInt64Coin("ukex", 10000000)}
	suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
	suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr2, coins)

	oldTreasury := suite.app.DistrKeeper.GetFeesTreasury(suite.ctx)
	suite.app.DistrKeeper.AllocateTokens(suite.ctx, 10, 10, consAddr, []abci.VoteInfo{})
	newTreasury := suite.app.DistrKeeper.GetFeesTreasury(suite.ctx)
	suite.Require().True(oldTreasury.DenomsSubsetOf(newTreasury))

	// TODO: add case for validator exit case
	// TODO: add case for staking pool exist case
	// TODO: check tokens are distributed correctly
}
