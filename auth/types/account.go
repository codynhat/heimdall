package types

import (
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	yaml "gopkg.in/yaml.v2"

	"github.com/maticnetwork/heimdall/types"
)

var cdc = amino.NewCodec()

func init() {
	cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{}, secp256k1.PubKeyAminoName, nil)
	cdc.RegisterConcrete(secp256k1.PrivKeySecp256k1{}, secp256k1.PrivKeyAminoName, nil)
}

// Account is an interface used to store coins at a given address within state.
// It presumes a notion of sequence numbers for replay protection,
// a notion of account numbers for replay protection for previously pruned accounts,
// and a pubkey for authentication purposes.
//
// Many complex conditions can be used in the concrete struct which implements Account.
type Account interface {
	GetAddress() types.HeimdallAddress
	SetAddress(types.HeimdallAddress) error // errors if already set.

	GetPubKey() crypto.PubKey // can return nil.
	SetPubKey(crypto.PubKey) error

	GetAccountNumber() uint64
	SetAccountNumber(uint64) error

	GetSequence() uint64
	SetSequence(uint64) error

	GetCoins() types.Coins
	SetCoins(types.Coins) error

	// Calculates the amount of coins that can be sent to other accounts given
	// the current time.
	SpendableCoins(blockTime time.Time) types.Coins

	// Ensure that account implements stringer
	String() string
}

// VestingAccount defines an account type that vests coins via a vesting schedule.
type VestingAccount interface {
	Account

	// Delegation and undelegation accounting that returns the resulting base
	// coins amount.
	TrackDelegation(blockTime time.Time, amount types.Coins)
	TrackUndelegation(amount types.Coins)

	GetVestedCoins(blockTime time.Time) types.Coins
	GetVestingCoins(blockTime time.Time) types.Coins

	GetStartTime() int64
	GetEndTime() int64

	GetOriginalVesting() types.Coins
	GetDelegatedFree() types.Coins
	GetDelegatedVesting() types.Coins
}

//-----------------------------------------------------------------------------
// BaseAccount

var _ Account = (*BaseAccount)(nil)

// BaseAccount - a base account structure.
// This can be extended by embedding within in your AppAccount.
// However one doesn't have to use BaseAccount as long as your struct
// implements Account.
type BaseAccount struct {
	Address       types.HeimdallAddress `json:"address" yaml:"address"`
	Coins         types.Coins           `json:"coins" yaml:"coins"`
	PubKey        crypto.PubKey         `json:"public_key" yaml:"public_key"`
	AccountNumber uint64                `json:"account_number" yaml:"account_number"`
	Sequence      uint64                `json:"sequence" yaml:"sequence"`
}

// NewBaseAccount creates a new BaseAccount object
func NewBaseAccount(address types.HeimdallAddress, coins types.Coins,
	pubKey crypto.PubKey, accountNumber uint64, sequence uint64) *BaseAccount {

	return &BaseAccount{
		Address:       address,
		Coins:         coins,
		PubKey:        pubKey,
		AccountNumber: accountNumber,
		Sequence:      sequence,
	}
}

// String implements fmt.Stringer
func (acc BaseAccount) String() string {
	var pubkey string

	if acc.PubKey != nil {
		// pubkey = sdk.MustBech32ifyAccPub(acc.PubKey)
		var pubObject secp256k1.PubKeySecp256k1
		cdc.MustUnmarshalBinaryBare(acc.PubKey.Bytes(), &pubObject)
		pubkey = "0x" + hex.EncodeToString(pubObject[:])
	}

	return fmt.Sprintf(`Account:
  Address:       %s
  Pubkey:        %s
  Coins:         %s
  AccountNumber: %d
  Sequence:      %d`,
		acc.Address, pubkey, acc.Coins, acc.AccountNumber, acc.Sequence,
	)
}

// ProtoBaseAccount - a prototype function for BaseAccount
func ProtoBaseAccount() Account {
	return &BaseAccount{}
}

// NewBaseAccountWithAddress - returns a new base account with a given address
func NewBaseAccountWithAddress(addr types.HeimdallAddress) BaseAccount {
	return BaseAccount{
		Address: addr,
	}
}

// GetAddress - Implements sdk.Account.
func (acc BaseAccount) GetAddress() types.HeimdallAddress {
	return acc.Address
}

// SetAddress - Implements sdk.Account.
func (acc BaseAccount) SetAddress(addr types.HeimdallAddress) error {
	if len(acc.Address) != 0 && !acc.Address.Empty() {
		return errors.New("cannot override BaseAccount address")
	}
	acc.Address = addr
	return nil
}

// GetPubKey - Implements sdk.Account.
func (acc BaseAccount) GetPubKey() crypto.PubKey {
	return acc.PubKey
}

// SetPubKey - Implements sdk.Account.
func (acc BaseAccount) SetPubKey(pubKey crypto.PubKey) error {
	acc.PubKey = pubKey
	return nil
}

// GetCoins - Implements sdk.Account.
func (acc BaseAccount) GetCoins() types.Coins {
	return acc.Coins
}

// SetCoins - Implements sdk.Account.
func (acc BaseAccount) SetCoins(coins types.Coins) error {
	acc.Coins = coins
	return nil
}

// GetAccountNumber - Implements Account
func (acc BaseAccount) GetAccountNumber() uint64 {
	return acc.AccountNumber
}

// SetAccountNumber - Implements Account
func (acc BaseAccount) SetAccountNumber(accNumber uint64) error {
	acc.AccountNumber = accNumber
	return nil
}

// GetSequence - Implements sdk.Account.
func (acc BaseAccount) GetSequence() uint64 {
	return acc.Sequence
}

// SetSequence - Implements sdk.Account.
func (acc BaseAccount) SetSequence(seq uint64) error {
	acc.Sequence = seq
	return nil
}

// SpendableCoins returns the total set of spendable coins. For a base account,
// this is simply the base coins.
func (acc BaseAccount) SpendableCoins(_ time.Time) types.Coins {
	return acc.GetCoins()
}

// MarshalYAML returns the YAML representation of an account.
func (acc BaseAccount) MarshalYAML() (interface{}, error) {
	var bs []byte
	var err error
	var pubkey string

	if acc.PubKey != nil {
		pubkey, err = sdk.Bech32ifyAccPub(acc.PubKey)
		if err != nil {
			return nil, err
		}
	}

	bs, err = yaml.Marshal(struct {
		Address       types.HeimdallAddress
		Coins         types.Coins
		PubKey        string
		AccountNumber uint64
		Sequence      uint64
	}{
		Address:       acc.Address,
		Coins:         acc.Coins,
		PubKey:        pubkey,
		AccountNumber: acc.AccountNumber,
		Sequence:      acc.Sequence,
	})
	if err != nil {
		return nil, err
	}

	return string(bs), err
}

//-----------------------------------------------------------------------------
// Base Vesting Account

// BaseVestingAccount implements the VestingAccount interface. It contains all
// the necessary fields needed for any vesting account implementation.
type BaseVestingAccount struct {
	*BaseAccount

	OriginalVesting  types.Coins `json:"original_vesting"`  // coins in account upon initialization
	DelegatedFree    types.Coins `json:"delegated_free"`    // coins that are vested and delegated
	DelegatedVesting types.Coins `json:"delegated_vesting"` // coins that vesting and delegated

	EndTime int64 `json:"end_time"` // when the coins become unlocked
}

// NewBaseVestingAccount creates a new BaseVestingAccount object
func NewBaseVestingAccount(baseAccount *BaseAccount, originalVesting types.Coins,
	delegatedFree types.Coins, delegatedVesting types.Coins, endTime int64) *BaseVestingAccount {

	return &BaseVestingAccount{
		BaseAccount:      baseAccount,
		OriginalVesting:  originalVesting,
		DelegatedFree:    delegatedFree,
		DelegatedVesting: delegatedVesting,
		EndTime:          endTime,
	}
}

// String implements fmt.Stringer
func (bva BaseVestingAccount) String() string {
	var pubkey string

	if bva.PubKey != nil {
		pubkey = sdk.MustBech32ifyAccPub(bva.PubKey)
	}

	return fmt.Sprintf(`Vesting Account:
  Address:          %s
  Pubkey:           %s
  Coins:            %s
  AccountNumber:    %d
  Sequence:         %d
  OriginalVesting:  %s
  DelegatedFree:    %s
  DelegatedVesting: %s
  EndTime:          %d `,
		bva.Address, pubkey, bva.Coins, bva.AccountNumber, bva.Sequence,
		bva.OriginalVesting, bva.DelegatedFree, bva.DelegatedVesting, bva.EndTime,
	)
}

// spendableCoins returns all the spendable coins for a vesting account given a
// set of vesting coins.
//
// CONTRACT: The account's coins, delegated vesting coins, vestingCoins must be
// sorted.
func (bva BaseVestingAccount) spendableCoins(vestingCoins types.Coins) types.Coins {
	var spendableCoins types.Coins
	bc := bva.GetCoins()

	for _, coin := range bc {
		// zip/lineup all coins by their denomination to provide O(n) time
		baseAmt := coin.Amount
		vestingAmt := vestingCoins.AmountOf(coin.Denom)
		delVestingAmt := bva.DelegatedVesting.AmountOf(coin.Denom)

		// compute min((BC + DV) - V, BC) per the specification
		min := types.MinInt(baseAmt.Add(delVestingAmt).Sub(vestingAmt), baseAmt)
		spendableCoin := types.NewCoin(coin.Denom, min)

		if !spendableCoin.IsZero() {
			spendableCoins = spendableCoins.Add(types.Coins{spendableCoin})
		}
	}

	return spendableCoins
}

// trackDelegation tracks a delegation amount for any given vesting account type
// given the amount of coins currently vesting. It returns the resulting base
// coins.
//
// CONTRACT: The account's coins, delegation coins, vesting coins, and delegated
// vesting coins must be sorted.
func (bva *BaseVestingAccount) trackDelegation(vestingCoins, amount types.Coins) {
	bc := bva.GetCoins()

	for _, coin := range amount {
		// zip/lineup all coins by their denomination to provide O(n) time

		baseAmt := bc.AmountOf(coin.Denom)
		vestingAmt := vestingCoins.AmountOf(coin.Denom)
		delVestingAmt := bva.DelegatedVesting.AmountOf(coin.Denom)

		// Panic if the delegation amount is zero or if the base coins does not
		// exceed the desired delegation amount.
		if coin.Amount.IsZero() || baseAmt.LT(coin.Amount) {
			panic("delegation attempt with zero coins or insufficient funds")
		}

		// compute x and y per the specification, where:
		// X := min(max(V - DV, 0), D)
		// Y := D - X
		x := types.MinInt(types.MaxInt(vestingAmt.Sub(delVestingAmt), types.ZeroInt()), coin.Amount)
		y := coin.Amount.Sub(x)

		if !x.IsZero() {
			xCoin := types.NewCoin(coin.Denom, x)
			bva.DelegatedVesting = bva.DelegatedVesting.Add(types.Coins{xCoin})
		}

		if !y.IsZero() {
			yCoin := types.NewCoin(coin.Denom, y)
			bva.DelegatedFree = bva.DelegatedFree.Add(types.Coins{yCoin})
		}

		bva.Coins = bva.Coins.Sub(types.Coins{coin})
	}
}

// TrackUndelegation tracks an undelegation amount by setting the necessary
// values by which delegated vesting and delegated vesting need to decrease and
// by which amount the base coins need to increase. The resulting base coins are
// returned.
//
// NOTE: The undelegation (bond refund) amount may exceed the delegated
// vesting (bond) amount due to the way undelegation truncates the bond refund,
// which can increase the validator's exchange rate (tokens/shares) slightly if
// the undelegated tokens are non-integral.
//
// CONTRACT: The account's coins and undelegation coins must be sorted.
func (bva *BaseVestingAccount) TrackUndelegation(amount types.Coins) {
	for _, coin := range amount {
		// panic if the undelegation amount is zero
		if coin.Amount.IsZero() {
			panic("undelegation attempt with zero coins")
		}
		delegatedFree := bva.DelegatedFree.AmountOf(coin.Denom)
		delegatedVesting := bva.DelegatedVesting.AmountOf(coin.Denom)

		// compute x and y per the specification, where:
		// X := min(DF, D)
		// Y := min(DV, D - X)
		x := types.MinInt(delegatedFree, coin.Amount)
		y := types.MinInt(delegatedVesting, coin.Amount.Sub(x))

		if !x.IsZero() {
			xCoin := types.NewCoin(coin.Denom, x)
			bva.DelegatedFree = bva.DelegatedFree.Sub(types.Coins{xCoin})
		}

		if !y.IsZero() {
			yCoin := types.NewCoin(coin.Denom, y)
			bva.DelegatedVesting = bva.DelegatedVesting.Sub(types.Coins{yCoin})
		}

		bva.Coins = bva.Coins.Add(types.Coins{coin})
	}
}

// GetOriginalVesting returns a vesting account's original vesting amount
func (bva BaseVestingAccount) GetOriginalVesting() types.Coins {
	return bva.OriginalVesting
}

// GetDelegatedFree returns a vesting account's delegation amount that is not
// vesting.
func (bva BaseVestingAccount) GetDelegatedFree() types.Coins {
	return bva.DelegatedFree
}

// GetDelegatedVesting returns a vesting account's delegation amount that is
// still vesting.
func (bva BaseVestingAccount) GetDelegatedVesting() types.Coins {
	return bva.DelegatedVesting
}

//-----------------------------------------------------------------------------
// Continuous Vesting Account

var _ VestingAccount = (*ContinuousVestingAccount)(nil)

// ContinuousVestingAccount implements the VestingAccount interface. It
// continuously vests by unlocking coins linearly with respect to time.
type ContinuousVestingAccount struct {
	*BaseVestingAccount

	StartTime int64 `json:"start_time"` // when the coins start to vest
}

// NewContinuousVestingAccountRaw creates a new ContinuousVestingAccount object from BaseVestingAccount
func NewContinuousVestingAccountRaw(bva *BaseVestingAccount,
	startTime int64) *ContinuousVestingAccount {

	return &ContinuousVestingAccount{
		BaseVestingAccount: bva,
		StartTime:          startTime,
	}
}

// NewContinuousVestingAccount returns a new ContinuousVestingAccount
func NewContinuousVestingAccount(
	baseAcc *BaseAccount, StartTime, EndTime int64,
) *ContinuousVestingAccount {

	baseVestingAcc := &BaseVestingAccount{
		BaseAccount:     baseAcc,
		OriginalVesting: baseAcc.Coins,
		EndTime:         EndTime,
	}

	return &ContinuousVestingAccount{
		StartTime:          StartTime,
		BaseVestingAccount: baseVestingAcc,
	}
}

func (cva ContinuousVestingAccount) String() string {
	var pubkey string

	if cva.PubKey != nil {
		pubkey = sdk.MustBech32ifyAccPub(cva.PubKey)
	}

	return fmt.Sprintf(`Continuous Vesting Account:
  Address:          %s
  Pubkey:           %s
  Coins:            %s
  AccountNumber:    %d
  Sequence:         %d
  OriginalVesting:  %s
  DelegatedFree:    %s
  DelegatedVesting: %s
  StartTime:        %d
  EndTime:          %d `,
		cva.Address, pubkey, cva.Coins, cva.AccountNumber, cva.Sequence,
		cva.OriginalVesting, cva.DelegatedFree, cva.DelegatedVesting,
		cva.StartTime, cva.EndTime,
	)
}

// GetVestedCoins returns the total number of vested coins. If no coins are vested,
// nil is returned.
func (cva ContinuousVestingAccount) GetVestedCoins(blockTime time.Time) types.Coins {
	var vestedCoins types.Coins

	// We must handle the case where the start time for a vesting account has
	// been set into the future or when the start of the chain is not exactly
	// known.
	if blockTime.Unix() <= cva.StartTime {
		return vestedCoins
	} else if blockTime.Unix() >= cva.EndTime {
		return cva.OriginalVesting
	}

	// calculate the vesting scalar
	x := blockTime.Unix() - cva.StartTime
	y := cva.EndTime - cva.StartTime
	s := types.NewDec(x).Quo(types.NewDec(y))

	for _, ovc := range cva.OriginalVesting {
		vestedAmt := ovc.Amount.ToDec().Mul(s).RoundInt()
		vestedCoins = append(vestedCoins, types.NewCoin(ovc.Denom, vestedAmt))
	}

	return vestedCoins
}

// GetVestingCoins returns the total number of vesting coins. If no coins are
// vesting, nil is returned.
func (cva ContinuousVestingAccount) GetVestingCoins(blockTime time.Time) types.Coins {
	return cva.OriginalVesting.Sub(cva.GetVestedCoins(blockTime))
}

// SpendableCoins returns the total number of spendable coins per denom for a
// continuous vesting account.
func (cva ContinuousVestingAccount) SpendableCoins(blockTime time.Time) types.Coins {
	return cva.spendableCoins(cva.GetVestingCoins(blockTime))
}

// TrackDelegation tracks a desired delegation amount by setting the appropriate
// values for the amount of delegated vesting, delegated free, and reducing the
// overall amount of base coins.
func (cva *ContinuousVestingAccount) TrackDelegation(blockTime time.Time, amount types.Coins) {
	cva.trackDelegation(cva.GetVestingCoins(blockTime), amount)
}

// GetStartTime returns the time when vesting starts for a continuous vesting
// account.
func (cva *ContinuousVestingAccount) GetStartTime() int64 {
	return cva.StartTime
}

// GetEndTime returns the time when vesting ends for a continuous vesting account.
func (cva *ContinuousVestingAccount) GetEndTime() int64 {
	return cva.EndTime
}

//-----------------------------------------------------------------------------
// Delayed Vesting Account

var _ VestingAccount = (*DelayedVestingAccount)(nil)

// DelayedVestingAccount implements the VestingAccount interface. It vests all
// coins after a specific time, but non prior. In other words, it keeps them
// locked until a specified time.
type DelayedVestingAccount struct {
	*BaseVestingAccount
}

// NewDelayedVestingAccountRaw creates a new DelayedVestingAccount object from BaseVestingAccount
func NewDelayedVestingAccountRaw(bva *BaseVestingAccount) *DelayedVestingAccount {
	return &DelayedVestingAccount{
		BaseVestingAccount: bva,
	}
}

// NewDelayedVestingAccount returns a DelayedVestingAccount
func NewDelayedVestingAccount(baseAcc *BaseAccount, EndTime int64) *DelayedVestingAccount {
	baseVestingAcc := &BaseVestingAccount{
		BaseAccount:     baseAcc,
		OriginalVesting: baseAcc.Coins,
		EndTime:         EndTime,
	}

	return &DelayedVestingAccount{baseVestingAcc}
}

// GetVestedCoins returns the total amount of vested coins for a delayed vesting
// account. All coins are only vested once the schedule has elapsed.
func (dva DelayedVestingAccount) GetVestedCoins(blockTime time.Time) types.Coins {
	if blockTime.Unix() >= dva.EndTime {
		return dva.OriginalVesting
	}

	return nil
}

// GetVestingCoins returns the total number of vesting coins for a delayed
// vesting account.
func (dva DelayedVestingAccount) GetVestingCoins(blockTime time.Time) types.Coins {
	return dva.OriginalVesting.Sub(dva.GetVestedCoins(blockTime))
}

// SpendableCoins returns the total number of spendable coins for a delayed
// vesting account.
func (dva DelayedVestingAccount) SpendableCoins(blockTime time.Time) types.Coins {
	return dva.spendableCoins(dva.GetVestingCoins(blockTime))
}

// TrackDelegation tracks a desired delegation amount by setting the appropriate
// values for the amount of delegated vesting, delegated free, and reducing the
// overall amount of base coins.
func (dva *DelayedVestingAccount) TrackDelegation(blockTime time.Time, amount types.Coins) {
	dva.trackDelegation(dva.GetVestingCoins(blockTime), amount)
}

// GetStartTime returns zero since a delayed vesting account has no start time.
func (dva *DelayedVestingAccount) GetStartTime() int64 {
	return 0
}

// GetEndTime returns the time when vesting ends for a delayed vesting account.
func (dva *DelayedVestingAccount) GetEndTime() int64 {
	return dva.EndTime
}

//
// light base account
//

// LightBaseAccount - a base account structure.
type LightBaseAccount struct {
	Address       types.HeimdallAddress `json:"address" yaml:"address"`
	AccountNumber uint64                `json:"account_number" yaml:"account_number"`
	Sequence      uint64                `json:"sequence" yaml:"sequence"`
}
