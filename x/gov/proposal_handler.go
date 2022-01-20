package gov

import (
	"fmt"

	kiratypes "github.com/KiraCore/sekai/types"
	"github.com/KiraCore/sekai/x/gov/keeper"
	"github.com/KiraCore/sekai/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
)

type ApplyWhitelistAccountPermissionProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyWhitelistAccountPermissionProposalHandler(keeper keeper.Keeper) *ApplyWhitelistAccountPermissionProposalHandler {
	return &ApplyWhitelistAccountPermissionProposalHandler{keeper: keeper}
}

func (a ApplyWhitelistAccountPermissionProposalHandler) ProposalType() string {
	return kiratypes.WhitelistAccountPermissionProposalType
}

func (a ApplyWhitelistAccountPermissionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.WhitelistAccountPermissionProposal)

	actor, found := a.keeper.GetNetworkActorByAddress(ctx, p.Address)
	if found {
		if actor.Permissions.IsWhitelisted(types.PermValue(p.Permission)) {
			return sdkerrors.Wrap(types.ErrWhitelisting, "permission already whitelisted")
		}
	} else {
		actor = types.NewDefaultActor(p.Address)
	}

	return a.keeper.AddWhitelistPermission(ctx, actor, types.PermValue(p.Permission))
}

type ApplyBlacklistAccountPermissionProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyBlacklistAccountPermissionProposalHandler(keeper keeper.Keeper) *ApplyBlacklistAccountPermissionProposalHandler {
	return &ApplyBlacklistAccountPermissionProposalHandler{keeper: keeper}
}

func (a ApplyBlacklistAccountPermissionProposalHandler) ProposalType() string {
	return kiratypes.BlacklistAccountPermissionProposalType
}

func (a ApplyBlacklistAccountPermissionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.BlacklistAccountPermissionProposal)

	actor, found := a.keeper.GetNetworkActorByAddress(ctx, p.Address)
	if found {
		if actor.Permissions.IsWhitelisted(types.PermValue(p.Permission)) {
			return sdkerrors.Wrap(types.ErrWhitelisting, "permission already blacklisted")
		}
	} else {
		actor = types.NewDefaultActor(p.Address)
	}

	return a.keeper.AddBlacklistPermission(ctx, actor, types.PermValue(p.Permission))
}

type ApplyRemoveWhitelistedAccountPermissionProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyRemoveWhitelistedAccountPermissionProposalHandler(keeper keeper.Keeper) *ApplyRemoveWhitelistedAccountPermissionProposalHandler {
	return &ApplyRemoveWhitelistedAccountPermissionProposalHandler{keeper: keeper}
}

func (a ApplyRemoveWhitelistedAccountPermissionProposalHandler) ProposalType() string {
	return kiratypes.RemoveWhitelistedAccountPermissionProposalType
}

func (a ApplyRemoveWhitelistedAccountPermissionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.RemoveWhitelistedAccountPermissionProposal)

	actor, found := a.keeper.GetNetworkActorByAddress(ctx, p.Address)
	if found {
		if actor.Permissions.IsWhitelisted(types.PermValue(p.Permission)) {
			return sdkerrors.Wrap(types.ErrWhitelisting, "permission already blacklisted")
		}
	} else {
		actor = types.NewDefaultActor(p.Address)
	}

	return a.keeper.RemoveWhitelistedAccountPermission(ctx, actor, types.PermValue(p.Permission))
}

type ApplyRemoveBlacklistedAccountPermissionProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyRemoveBlacklistedAccountPermissionProposalHandler(keeper keeper.Keeper) *ApplyRemoveBlacklistedAccountPermissionProposalHandler {
	return &ApplyRemoveBlacklistedAccountPermissionProposalHandler{keeper: keeper}
}

func (a ApplyRemoveBlacklistedAccountPermissionProposalHandler) ProposalType() string {
	return kiratypes.RemoveBlacklistedAccountPermissionProposalType
}

func (a ApplyRemoveBlacklistedAccountPermissionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.RemoveBlacklistedAccountPermissionProposal)

	actor, found := a.keeper.GetNetworkActorByAddress(ctx, p.Address)
	if found {
		if actor.Permissions.IsWhitelisted(types.PermValue(p.Permission)) {
			return sdkerrors.Wrap(types.ErrWhitelisting, "permission already blacklisted")
		}
	} else {
		actor = types.NewDefaultActor(p.Address)
	}

	return a.keeper.RemoveBlacklistedAccountPermission(ctx, actor, types.PermValue(p.Permission))
}

type ApplyAssignRoleToAccountProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyAssignRoleToAccountProposalHandler(keeper keeper.Keeper) *ApplyAssignRoleToAccountProposalHandler {
	return &ApplyAssignRoleToAccountProposalHandler{keeper: keeper}
}

func (a ApplyAssignRoleToAccountProposalHandler) ProposalType() string {
	return kiratypes.RemoveBlacklistedAccountPermissionProposalType
}

func (a ApplyAssignRoleToAccountProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.RemoveBlacklistedAccountPermissionProposal)

	actor, found := a.keeper.GetNetworkActorByAddress(ctx, p.Address)
	if found {
		if actor.Permissions.IsWhitelisted(types.PermValue(p.Permission)) {
			return sdkerrors.Wrap(types.ErrWhitelisting, "permission already blacklisted")
		}
	} else {
		actor = types.NewDefaultActor(p.Address)
	}

	return a.keeper.AssignRoleToAccount(ctx, actor, types.PermValue(p.Permission))
}

type ApplyUnassignRoleFromAccountProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUnassignRoleFromAccountProposalHandler(keeper keeper.Keeper) *ApplyAssignRoleToAccountProposalHandler {
	return &ApplyAssignRoleFromAccountProposalHandler{keeper: keeper}
}

func (a ApplyAssignRoleFromAccountProposalHandler) ProposalType() string {
	return kiratypes.RemoveBlacklistedAccountPermissionProposalType
}

func (a ApplyAssignRoleFromAccountProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.ApplyUnassignRoleFromAccountProposal)

	actor, found := a.keeper.GetNetworkActorByAddress(ctx, p.Address)
	if found {
		if actor.Permissions.IsWhitelisted(types.PermValue(p.Permission)) {
			return sdkerrors.Wrap(types.ErrWhitelisting, "permission already blacklisted")
		}
	} else {
		actor = types.NewDefaultActor(p.Address)
	}

	return a.keeper.UnassignRoleFromAccount(ctx, actor, types.PermValue(p.Permission))
}

type ApplySetNetworkPropertyProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplySetNetworkPropertyProposalHandler(keeper keeper.Keeper) *ApplySetNetworkPropertyProposalHandler {
	return &ApplySetNetworkPropertyProposalHandler{keeper: keeper}
}

func (a ApplySetNetworkPropertyProposalHandler) ProposalType() string {
	return kiratypes.SetNetworkPropertyProposalType
}

func (a ApplySetNetworkPropertyProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.SetNetworkPropertyProposal)

	property, err := a.keeper.GetNetworkProperty(ctx, p.NetworkProperty)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	if property == p.Value {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "network property already set as proposed value")
	}

	return a.keeper.SetNetworkProperty(ctx, p.NetworkProperty, p.Value)
}

type ApplyUpsertDataRegistryProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUpsertDataRegistryProposalHandler(keeper keeper.Keeper) *ApplyUpsertDataRegistryProposalHandler {
	return &ApplyUpsertDataRegistryProposalHandler{keeper: keeper}
}

func (a ApplyUpsertDataRegistryProposalHandler) ProposalType() string {
	return kiratypes.UpsertDataRegistryProposalType
}

func (a ApplyUpsertDataRegistryProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.UpsertDataRegistryProposal)
	entry := types.NewDataRegistryEntry(p.Hash, p.Reference, p.Encoding, p.Size_)
	a.keeper.UpsertDataRegistryEntry(ctx, p.Key, entry)
	return nil
}

type ApplySetPoorNetworkMessagesProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplySetPoorNetworkMessagesProposalHandler(keeper keeper.Keeper) *ApplySetPoorNetworkMessagesProposalHandler {
	return &ApplySetPoorNetworkMessagesProposalHandler{keeper: keeper}
}

func (a ApplySetPoorNetworkMessagesProposalHandler) ProposalType() string {
	return kiratypes.SetPoorNetworkMessagesProposalType
}

func (a ApplySetPoorNetworkMessagesProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.SetPoorNetworkMessagesProposal)
	msgs := types.AllowedMessages{Messages: p.Messages}
	a.keeper.SavePoorNetworkMessages(ctx, &msgs)
	return nil
}

type CreateRoleProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyCreateRoleProposalHandler(keeper keeper.Keeper) *CreateRoleProposalHandler {
	return &CreateRoleProposalHandler{keeper: keeper}
}

func (c CreateRoleProposalHandler) ProposalType() string {
	return kiratypes.CreateRoleProposalType
}

func (c CreateRoleProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.CreateRoleProposal)

	// check sid is good variable naming form
	if !keeper.ValidateRoleSidKey(p.RoleSid) {
		return errors.Wrap(types.ErrInvalidRoleSid, fmt.Sprintf("invalid role sid configuration: sid=%s", p.RoleSid))
	}

	_, err := c.keeper.GetRoleBySid(ctx, p.RoleSid)
	if err == nil {
		return types.ErrRoleExist
	}

	roleId := c.keeper.CreateRole(ctx, p.RoleSid, p.RoleDescription)

	for _, w := range p.WhitelistedPermissions {
		err := c.keeper.WhitelistRolePermission(ctx, roleId, w)
		if err != nil {
			return err
		}
	}

	for _, b := range p.BlacklistedPermissions {
		err := c.keeper.BlacklistRolePermission(ctx, roleId, b)
		if err != nil {
			return err
		}
	}
	return nil
}

type ApplyRemoveRoleProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyRemoveRoleProposalHandler(keeper keeper.Keeper) *ApplyRemoveRoleProposalHandler {
	return &ApplyRemoveRoleProposalHandler{keeper: keeper}
}

func (c ApplyRemoveRoleProposalHandler) ProposalType() string {
	return kiratypes.RemoveRoleProposalType
}

func (c ApplyRemoveRoleProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.RemoveRoleProposal)
	return nil
}

type ApplyWhitelistRolePermissionProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyWhitelistRolePermissionProposalHandler(keeper keeper.Keeper) *ApplyWhitelistRolePermissionProposalHandler {
	return &ApplyWhitelistRolePermissionProposalHandler{keeper: keeper}
}

func (c ApplyWhitelistRolePermissionProposalHandler) ProposalType() string {
	return kiratypes.WhitelistRolePermissionProposalType
}

func (c ApplyWhitelistRolePermissionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.WhitelistRolePermissionProposal)
	return nil
}

type ApplyBlacklistRolePermissionProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyBlacklistRolePermissionProposalHandler(keeper keeper.Keeper) *ApplyBlacklistRolePermissionProposalHandler {
	return &ApplyBlacklistRolePermissionProposalHandler{keeper: keeper}
}

func (c ApplyBlacklistRolePermissionProposalHandler) ProposalType() string {
	return kiratypes.BlacklistRolePermissionProposalType
}

func (c ApplyBlacklistRolePermissionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.BlacklistRolePermissionProposal)
	return nil
}

type ApplyRemoveWhitelistedRolePermissionProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyRemoveWhitelistedRolePermissionProposalHandler(keeper keeper.Keeper) *ApplyRemoveWhitelistedRolePermissionProposalHandler {
	return &ApplyRemoveWhitelistedRolePermissionProposalHandler{keeper: keeper}
}

func (c ApplyRemoveWhitelistedRolePermissionProposalHandler) ProposalType() string {
	return kiratypes.RemoveWhitelistedRolePermissionProposalType
}

func (c ApplyRemoveWhitelistedRolePermissionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.RemoveWhitelistedRolePermissionProposal)
	return nil
}

type ApplyRemoveBlacklistedRolePermissionProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyRemoveBlacklistedRolePermissionProposalHandler(keeper keeper.Keeper) *ApplyRemoveBlacklistedRolePermissionProposalHandler {
	return &ApplyRemoveBlacklistedRolePermissionProposalHandler{keeper: keeper}
}

func (c ApplyRemoveBlacklistedRolePermissionProposalHandler) ProposalType() string {
	return kiratypes.RemoveBlacklistedRolePermissionProposalType
}

func (c ApplyRemoveBlacklistedRolePermissionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.RemoveBlacklistedRolePermissionProposal)
	return nil
}

type SetProposalDurationsProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplySetProposalDurationsProposalHandler(keeper keeper.Keeper) *SetProposalDurationsProposalHandler {
	return &SetProposalDurationsProposalHandler{keeper: keeper}
}

func (c SetProposalDurationsProposalHandler) ProposalType() string {
	return kiratypes.SetProposalDurationsProposalType
}

func (c SetProposalDurationsProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.SetProposalDurationsProposal)
	for i, pt := range p.TypeofProposals {
		err := c.keeper.SetProposalDuration(ctx, pt, p.ProposalDurations[i])
		if err != nil {
			return nil
		}
	}
	return nil
}
