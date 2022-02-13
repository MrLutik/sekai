package spending

import (
	kiratypes "github.com/KiraCore/sekai/types"
	govtypes "github.com/KiraCore/sekai/x/gov/types"
	"github.com/KiraCore/sekai/x/spending/keeper"
	"github.com/KiraCore/sekai/x/spending/types"
	spendingtypes "github.com/KiraCore/sekai/x/spending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplyUpdateSpendingPoolProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUpdateSpendingPoolProposalHandler(keeper keeper.Keeper) *ApplyUpdateSpendingPoolProposalHandler {
	return &ApplyUpdateSpendingPoolProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyUpdateSpendingPoolProposalHandler) ProposalType() string {
	return kiratypes.ProposalTypeUpdateSpendingPool
}

func (a ApplyUpdateSpendingPoolProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content) error {
	p := proposal.(*spendingtypes.UpdateSpendingPoolProposal)

	pool := a.keeper.GetSpendingPool(ctx, p.Name)
	if pool == nil {
		return types.ErrPoolDoesNotExist
	}

	a.keeper.SetSpendingPool(ctx, types.SpendingPool{
		Name:          p.Name,
		ClaimStart:    p.ClaimStart,
		ClaimEnd:      p.ClaimEnd,
		Expire:        p.Expire,
		Token:         p.Token,
		Rate:          p.Rate,
		VoteQuorum:    p.VoteQuorum,
		VotePeriod:    p.VotePeriod,
		VoteEnactment: p.VoteEnactment,
		Owners:        &p.Owners,
		Beneficiaries: &p.Beneficiaries,
		Balance:       pool.Balance,
	})

	return nil
}

type ApplySpendingPoolDistributionProposalHandler struct {
	keeper keeper.Keeper
	gk     types.CustomGovKeeper
}

func NewApplySpendingPoolDistributionProposalHandler(keeper keeper.Keeper, gk types.CustomGovKeeper) *ApplySpendingPoolDistributionProposalHandler {
	return &ApplySpendingPoolDistributionProposalHandler{
		keeper: keeper,
		gk:     gk,
	}
}

func (a ApplySpendingPoolDistributionProposalHandler) ProposalType() string {
	return kiratypes.ProposalTypeSpendingPoolDistribution
}

func (a ApplySpendingPoolDistributionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content) error {
	p := proposal.(*spendingtypes.SpendingPoolDistributionProposal)
	_ = p

	pool := a.keeper.GetSpendingPool(ctx, p.PoolName)
	duplicateMap := map[string]bool{}
	var beneficiaries []string

	for _, acc := range pool.Beneficiaries.OwnerAccounts {
		if _, ok := duplicateMap[acc]; !ok {
			duplicateMap[acc] = true
			beneficiaries = append(beneficiaries, acc)
		}
	}
	for _, role := range pool.Beneficiaries.OwnerRoles {
		actorIter := a.gk.GetNetworkActorsByRole(ctx, role)

		for ; actorIter.Valid(); actorIter.Next() {
			if _, ok := duplicateMap[sdk.AccAddress(actorIter.Value()).String()]; !ok {
				duplicateMap[sdk.AccAddress(actorIter.Value()).String()] = true
				beneficiaries = append(beneficiaries, sdk.AccAddress(actorIter.Value()).String())
			}
		}
	}

	for _, beneficiary := range beneficiaries {
		beneficiaryAcc, err := sdk.AccAddressFromBech32(beneficiary)
		if err != nil {
			return err
		}

		err = a.keeper.ClaimSpendingPool(ctx, p.PoolName, beneficiaryAcc)
		if err != nil {
			return err
		}
	}

	return nil
}

type ApplySpendingPoolWithdrawProposalHandler struct {
	keeper keeper.Keeper
	bk     types.BankKeeper
}

func NewApplySpendingPoolWithdrawProposalHandler(keeper keeper.Keeper, bk types.BankKeeper) *ApplySpendingPoolWithdrawProposalHandler {
	return &ApplySpendingPoolWithdrawProposalHandler{
		keeper: keeper,
		bk:     bk,
	}
}

func (a ApplySpendingPoolWithdrawProposalHandler) ProposalType() string {
	return kiratypes.ProposalTypeSpendingPoolWithdraw
}

func (a ApplySpendingPoolWithdrawProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content) error {
	p := proposal.(*spendingtypes.SpendingPoolWithdrawProposal)
	_ = p

	pool := a.keeper.GetSpendingPool(ctx, p.PoolName)
	if pool == nil {
		return types.ErrPoolDoesNotExist
	}

	for _, beneficiary := range p.Beneficiaries {
		beneficiaryAcc, err := sdk.AccAddressFromBech32(beneficiary)
		if err != nil {
			return err
		}

		if !a.keeper.IsAllowedAddress(ctx, beneficiaryAcc, *pool.Beneficiaries) {
			return types.ErrNotPoolBeneficiary
		}

		err = a.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, beneficiaryAcc, p.Amounts)
		if err != nil {
			return err
		}

		// update pool to reduce pool's balance
		pool.Balance = pool.Balance.Sub(sdk.Coins(p.Amounts).AmountOf(pool.Token))
	}

	a.keeper.SetSpendingPool(ctx, *pool)
	return nil
}
