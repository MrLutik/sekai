package keeper

import (
	"fmt"
	"strings"

	"github.com/KiraCore/sekai/x/multistaking/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: set id and validator mapping

func (k Keeper) GetLastPoolId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyLastPoolId)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetLastPoolId(ctx sdk.Context, id uint64) {
	idBz := sdk.Uint64ToBigEndian(id)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyLastPoolId, idBz)
}

func (k Keeper) GetAllStakingPools(ctx sdk.Context) []types.StakingPool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStakingPool)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	pools := []types.StakingPool{}
	for ; iterator.Valid(); iterator.Next() {
		pool := types.StakingPool{}
		k.cdc.MustUnmarshal(iterator.Value(), &pool)
		pools = append(pools, pool)
	}
	return pools
}

func (k Keeper) GetStakingPoolByValidator(ctx sdk.Context, validator string) (pool types.StakingPool, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.KeyPrefixStakingPool), []byte(validator)...)
	bz := store.Get(key)
	if bz == nil {
		return pool, false
	}
	k.cdc.MustUnmarshal(bz, &pool)
	return pool, true
}

func (k Keeper) SetStakingPool(ctx sdk.Context, pool types.StakingPool) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.KeyPrefixStakingPool), []byte(pool.Validator)...)
	store.Set(key, k.cdc.MustMarshal(&pool))
}

func (k Keeper) RemoveStakingPool(ctx sdk.Context, pool types.StakingPool) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.KeyPrefixStakingPool), []byte(pool.Validator)...)
	store.Delete(key)
}

func getPoolPrefix(poolID uint64) string {
	return fmt.Sprintf("v%d/", poolID)
}
func getPoolCoins(pool types.StakingPool, coins sdk.Coins) sdk.Coins {
	prefix := getPoolPrefix(pool.Id)
	poolCoins := sdk.Coins{}
	for _, coin := range coins {
		poolCoins = poolCoins.Add(sdk.NewCoin(prefix+coin.Denom, coin.Amount.ToDec().Mul(sdk.OneDec().Sub(pool.Slashed)).RoundInt()))
	}
	return poolCoins
}
func getShareDenom(poolID uint64, denom string) string {
	prefix := getPoolPrefix(poolID)
	return prefix + denom
}
func getNativeDenom(poolID uint64, denom string) string {
	return strings.TrimPrefix(denom, getPoolPrefix(poolID))
}
