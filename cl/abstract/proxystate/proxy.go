// Code generated; DO NOT EDIT.
package proxystate

import (
	"fmt"
	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon-lib/types/clonable"
	"github.com/ledgerwatch/erigon/cl/abstract"
	"github.com/ledgerwatch/erigon/cl/clparams"
	"github.com/ledgerwatch/erigon/cl/cltypes"
	"github.com/ledgerwatch/erigon/cl/cltypes/solid"
)

type BeaconStateProxy struct {
	Handler    InvocationHandler
	Underlying abstract.BeaconState
}

func (this *BeaconStateProxy) AddCurrentEpochAtteastation(arg0 *solid.PendingAttestation) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("AddCurrentEpochAtteastation", args)
	_ = rets
	if !intercept {
		this.Underlying.AddCurrentEpochAtteastation(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) AddCurrentEpochParticipationFlags(arg0 *cltypes.ParticipationFlags) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("AddCurrentEpochParticipationFlags", args)
	_ = rets
	if !intercept {
		this.Underlying.AddCurrentEpochParticipationFlags(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) AddEth1DataVote(arg0 *cltypes.Eth1Data) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("AddEth1DataVote", args)
	_ = rets
	if !intercept {
		this.Underlying.AddEth1DataVote(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) AddHistoricalRoot(arg0 common.Hash) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("AddHistoricalRoot", args)
	_ = rets
	if !intercept {
		this.Underlying.AddHistoricalRoot(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) AddHistoricalSummary(arg0 *cltypes.HistoricalSummary) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("AddHistoricalSummary", args)
	_ = rets
	if !intercept {
		this.Underlying.AddHistoricalSummary(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) AddInactivityScore(arg0 uint64) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("AddInactivityScore", args)
	_ = rets
	if !intercept {
		this.Underlying.AddInactivityScore(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) AddPreviousEpochAttestation(arg0 *solid.PendingAttestation) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("AddPreviousEpochAttestation", args)
	_ = rets
	if !intercept {
		this.Underlying.AddPreviousEpochAttestation(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) AddPreviousEpochParticipationAt(arg0 int, arg1 byte) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("AddPreviousEpochParticipationAt", args)
	_ = rets
	if !intercept {
		this.Underlying.AddPreviousEpochParticipationAt(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) AddPreviousEpochParticipationFlags(arg0 *cltypes.ParticipationFlags) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("AddPreviousEpochParticipationFlags", args)
	_ = rets
	if !intercept {
		this.Underlying.AddPreviousEpochParticipationFlags(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) AddValidator(arg0 solid.Validator, arg1 uint64) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("AddValidator", args)
	_ = rets
	if !intercept {
		this.Underlying.AddValidator(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) AppendValidator(arg0 solid.Validator) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("AppendValidator", args)
	_ = rets
	if !intercept {
		this.Underlying.AppendValidator(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) BaseReward(arg0 uint64) (ret0 uint64, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("BaseReward", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.BaseReward(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) BaseRewardPerIncrement() (ret0 uint64) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("BaseRewardPerIncrement", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.BaseRewardPerIncrement()
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) BeaconConfig() (ret0 *clparams.BeaconChainConfig) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("BeaconConfig", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.BeaconConfig()
		return
	}

	var ok bool

	ret0, ok = rets[0].(*clparams.BeaconChainConfig)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *clparams.BeaconChainConfig value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) BlockRoot() (ret0 [32]byte, ret1 error) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("BlockRoot", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.BlockRoot()
		return
	}

	var ok bool

	ret0, ok = rets[0].([32]byte)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid [32]byte value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) BlockRoots() (ret0 solid.HashVectorSSZ) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("BlockRoots", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.BlockRoots()
		return
	}

	var ok bool

	ret0, ok = rets[0].(solid.HashVectorSSZ)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid solid.HashVectorSSZ value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) Clone() (ret0 clonable.Clonable) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("Clone", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.Clone()
		return
	}

	var ok bool

	ret0, ok = rets[0].(clonable.Clonable)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid clonable.Clonable value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) CommitteeCount(arg0 uint64) (ret0 uint64) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("CommitteeCount", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.CommitteeCount(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) ComputeCommittee(arg0 []uint64, arg1 uint64, arg2 uint64, arg3 uint64) (ret0 []uint64, ret1 error) {
	args := []any{
		arg0,
		arg1,
		arg2,
		arg3,
	}
	rets, intercept := this.Handler.Invoke("ComputeCommittee", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ComputeCommittee(
			arg0,
			arg1,
			arg2,
			arg3,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].([]uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid []uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ComputeNextSyncCommittee() (ret0 *solid.SyncCommittee, ret1 error) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("ComputeNextSyncCommittee", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ComputeNextSyncCommittee()
		return
	}

	var ok bool

	ret0, ok = rets[0].(*solid.SyncCommittee)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *solid.SyncCommittee value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) CurrentEpochAttestations() (ret0 *solid.ListSSZ[*solid.PendingAttestation]) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("CurrentEpochAttestations", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.CurrentEpochAttestations()
		return
	}

	var ok bool

	ret0, ok = rets[0].(*solid.ListSSZ[*solid.PendingAttestation])
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *solid.ListSSZ[*solid.PendingAttestation] value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) CurrentEpochAttestationsLength() (ret0 int) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("CurrentEpochAttestationsLength", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.CurrentEpochAttestationsLength()
		return
	}

	var ok bool

	ret0, ok = rets[0].(int)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid int value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) CurrentJustifiedCheckpoint() (ret0 solid.Checkpoint) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("CurrentJustifiedCheckpoint", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.CurrentJustifiedCheckpoint()
		return
	}

	var ok bool

	ret0, ok = rets[0].(solid.Checkpoint)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid solid.Checkpoint value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) CurrentSyncCommittee() (ret0 *solid.SyncCommittee) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("CurrentSyncCommittee", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.CurrentSyncCommittee()
		return
	}

	var ok bool

	ret0, ok = rets[0].(*solid.SyncCommittee)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *solid.SyncCommittee value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) DecodeSSZ(arg0 []byte, arg1 int) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("DecodeSSZ", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.DecodeSSZ(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) EncodeSSZ(arg0 []byte) (ret0 []byte, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("EncodeSSZ", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.EncodeSSZ(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].([]byte)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid []byte value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) EncodingSizeSSZ() (ret0 int) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("EncodingSizeSSZ", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.EncodingSizeSSZ()
		return
	}

	var ok bool

	ret0, ok = rets[0].(int)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid int value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) EpochParticipation(arg0 bool) (ret0 *solid.BitList) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("EpochParticipation", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.EpochParticipation(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(*solid.BitList)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *solid.BitList value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) EpochParticipationForValidatorIndex(arg0 bool, arg1 int) (ret0 cltypes.ParticipationFlags) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("EpochParticipationForValidatorIndex", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.EpochParticipationForValidatorIndex(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(cltypes.ParticipationFlags)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid cltypes.ParticipationFlags value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) Eth1Data() (ret0 *cltypes.Eth1Data) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("Eth1Data", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.Eth1Data()
		return
	}

	var ok bool

	ret0, ok = rets[0].(*cltypes.Eth1Data)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *cltypes.Eth1Data value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) Eth1DataVotes() (ret0 *solid.ListSSZ[*cltypes.Eth1Data]) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("Eth1DataVotes", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.Eth1DataVotes()
		return
	}

	var ok bool

	ret0, ok = rets[0].(*solid.ListSSZ[*cltypes.Eth1Data])
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *solid.ListSSZ[*cltypes.Eth1Data] value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) Eth1DepositIndex() (ret0 uint64) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("Eth1DepositIndex", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.Eth1DepositIndex()
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) FinalizedCheckpoint() (ret0 solid.Checkpoint) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("FinalizedCheckpoint", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.FinalizedCheckpoint()
		return
	}

	var ok bool

	ret0, ok = rets[0].(solid.Checkpoint)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid solid.Checkpoint value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) ForEachBalance(arg0 func(v uint64, idx int, total int) bool) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ForEachBalance", args)
	_ = rets
	if !intercept {
		this.Underlying.ForEachBalance(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) ForEachSlashingSegment(arg0 func(idx int, v uint64, total int) bool) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ForEachSlashingSegment", args)
	_ = rets
	if !intercept {
		this.Underlying.ForEachSlashingSegment(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) ForEachValidator(arg0 func(v solid.Validator, idx int, total int) bool) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ForEachValidator", args)
	_ = rets
	if !intercept {
		this.Underlying.ForEachValidator(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) Fork() (ret0 *cltypes.Fork) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("Fork", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.Fork()
		return
	}

	var ok bool

	ret0, ok = rets[0].(*cltypes.Fork)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *cltypes.Fork value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) GenesisTime() (ret0 uint64) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("GenesisTime", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.GenesisTime()
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) GenesisValidatorsRoot() (ret0 common.Hash) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("GenesisValidatorsRoot", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.GenesisValidatorsRoot()
		return
	}

	var ok bool

	ret0, ok = rets[0].(common.Hash)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid common.Hash value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) GetActiveValidatorsIndices(arg0 uint64) (ret0 []uint64) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("GetActiveValidatorsIndices", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.GetActiveValidatorsIndices(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].([]uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid []uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) GetAttestationParticipationFlagIndicies(arg0 solid.AttestationData, arg1 uint64) (ret0 []uint8, ret1 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("GetAttestationParticipationFlagIndicies", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.GetAttestationParticipationFlagIndicies(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].([]uint8)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid []uint8 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) GetAttestingIndicies(arg0 solid.AttestationData, arg1 []byte, arg2 bool) (ret0 []uint64, ret1 error) {
	args := []any{
		arg0,
		arg1,
		arg2,
	}
	rets, intercept := this.Handler.Invoke("GetAttestingIndicies", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.GetAttestingIndicies(
			arg0,
			arg1,
			arg2,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].([]uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid []uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) GetBeaconCommitee(arg0 uint64, arg1 uint64) (ret0 []uint64, ret1 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("GetBeaconCommitee", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.GetBeaconCommitee(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].([]uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid []uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) GetBeaconProposerIndex() (ret0 uint64, ret1 error) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("GetBeaconProposerIndex", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.GetBeaconProposerIndex()
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) GetBlockRootAtSlot(arg0 uint64) (ret0 common.Hash, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("GetBlockRootAtSlot", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.GetBlockRootAtSlot(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(common.Hash)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid common.Hash value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) GetDomain(arg0 [4]byte, arg1 uint64) (ret0 []byte, ret1 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("GetDomain", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.GetDomain(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].([]byte)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid []byte value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) GetRandaoMix(arg0 int) (ret0 [32]byte) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("GetRandaoMix", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.GetRandaoMix(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].([32]byte)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid [32]byte value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) GetRandaoMixes(arg0 uint64) (ret0 [32]byte) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("GetRandaoMixes", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.GetRandaoMixes(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].([32]byte)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid [32]byte value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) GetTotalActiveBalance() (ret0 uint64) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("GetTotalActiveBalance", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.GetTotalActiveBalance()
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) GetValidatorChurnLimit() (ret0 uint64) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("GetValidatorChurnLimit", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.GetValidatorChurnLimit()
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) HashSSZ() (ret0 [32]byte, ret1 error) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("HashSSZ", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.HashSSZ()
		return
	}

	var ok bool

	ret0, ok = rets[0].([32]byte)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid [32]byte value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) IncrementSlashingSegmentAt(arg0 int, arg1 uint64) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("IncrementSlashingSegmentAt", args)
	_ = rets
	if !intercept {
		this.Underlying.IncrementSlashingSegmentAt(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) InitiateValidatorExit(arg0 uint64) (ret0 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("InitiateValidatorExit", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.InitiateValidatorExit(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) JustificationBits() (ret0 cltypes.JustificationBits) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("JustificationBits", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.JustificationBits()
		return
	}

	var ok bool

	ret0, ok = rets[0].(cltypes.JustificationBits)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid cltypes.JustificationBits value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) LatestBlockHeader() (ret0 cltypes.BeaconBlockHeader) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("LatestBlockHeader", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.LatestBlockHeader()
		return
	}

	var ok bool

	ret0, ok = rets[0].(cltypes.BeaconBlockHeader)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid cltypes.BeaconBlockHeader value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) LatestExecutionPayloadHeader() (ret0 *cltypes.Eth1Header) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("LatestExecutionPayloadHeader", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.LatestExecutionPayloadHeader()
		return
	}

	var ok bool

	ret0, ok = rets[0].(*cltypes.Eth1Header)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *cltypes.Eth1Header value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) NextSyncCommittee() (ret0 *solid.SyncCommittee) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("NextSyncCommittee", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.NextSyncCommittee()
		return
	}

	var ok bool

	ret0, ok = rets[0].(*solid.SyncCommittee)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *solid.SyncCommittee value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) NextWithdrawalIndex() (ret0 uint64) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("NextWithdrawalIndex", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.NextWithdrawalIndex()
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) NextWithdrawalValidatorIndex() (ret0 uint64) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("NextWithdrawalValidatorIndex", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.NextWithdrawalValidatorIndex()
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) PreviousEpochAttestations() (ret0 *solid.ListSSZ[*solid.PendingAttestation]) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("PreviousEpochAttestations", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.PreviousEpochAttestations()
		return
	}

	var ok bool

	ret0, ok = rets[0].(*solid.ListSSZ[*solid.PendingAttestation])
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *solid.ListSSZ[*solid.PendingAttestation] value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) PreviousEpochAttestationsLength() (ret0 int) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("PreviousEpochAttestationsLength", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.PreviousEpochAttestationsLength()
		return
	}

	var ok bool

	ret0, ok = rets[0].(int)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid int value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) PreviousJustifiedCheckpoint() (ret0 solid.Checkpoint) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("PreviousJustifiedCheckpoint", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.PreviousJustifiedCheckpoint()
		return
	}

	var ok bool

	ret0, ok = rets[0].(solid.Checkpoint)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid solid.Checkpoint value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) PreviousSlot() (ret0 uint64) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("PreviousSlot", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.PreviousSlot()
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) PreviousStateRoot() (ret0 common.Hash) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("PreviousStateRoot", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.PreviousStateRoot()
		return
	}

	var ok bool

	ret0, ok = rets[0].(common.Hash)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid common.Hash value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) RandaoMixes() (ret0 solid.HashVectorSSZ) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("RandaoMixes", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.RandaoMixes()
		return
	}

	var ok bool

	ret0, ok = rets[0].(solid.HashVectorSSZ)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid solid.HashVectorSSZ value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) ResetCurrentEpochAttestations() {
	args := []any{}
	rets, intercept := this.Handler.Invoke("ResetCurrentEpochAttestations", args)
	_ = rets
	if !intercept {
		this.Underlying.ResetCurrentEpochAttestations()
		return
	}

}

func (this *BeaconStateProxy) ResetEpochParticipation() {
	args := []any{}
	rets, intercept := this.Handler.Invoke("ResetEpochParticipation", args)
	_ = rets
	if !intercept {
		this.Underlying.ResetEpochParticipation()
		return
	}

}

func (this *BeaconStateProxy) ResetEth1DataVotes() {
	args := []any{}
	rets, intercept := this.Handler.Invoke("ResetEth1DataVotes", args)
	_ = rets
	if !intercept {
		this.Underlying.ResetEth1DataVotes()
		return
	}

}

func (this *BeaconStateProxy) ResetHistoricalSummaries() {
	args := []any{}
	rets, intercept := this.Handler.Invoke("ResetHistoricalSummaries", args)
	_ = rets
	if !intercept {
		this.Underlying.ResetHistoricalSummaries()
		return
	}

}

func (this *BeaconStateProxy) ResetPreviousEpochAttestations() {
	args := []any{}
	rets, intercept := this.Handler.Invoke("ResetPreviousEpochAttestations", args)
	_ = rets
	if !intercept {
		this.Underlying.ResetPreviousEpochAttestations()
		return
	}

}

func (this *BeaconStateProxy) SetActivationEligibilityEpochForValidatorAtIndex(arg0 int, arg1 uint64) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetActivationEligibilityEpochForValidatorAtIndex", args)
	_ = rets
	if !intercept {
		this.Underlying.SetActivationEligibilityEpochForValidatorAtIndex(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) SetActivationEpochForValidatorAtIndex(arg0 int, arg1 uint64) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetActivationEpochForValidatorAtIndex", args)
	_ = rets
	if !intercept {
		this.Underlying.SetActivationEpochForValidatorAtIndex(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) SetBlockRootAt(arg0 int, arg1 common.Hash) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetBlockRootAt", args)
	_ = rets
	if !intercept {
		this.Underlying.SetBlockRootAt(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) SetCurrentEpochParticipationFlags(arg0 *cltypes.ParticipationFlagsList) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetCurrentEpochParticipationFlags", args)
	_ = rets
	if !intercept {
		this.Underlying.SetCurrentEpochParticipationFlags(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetCurrentJustifiedCheckpoint(arg0 solid.Checkpoint) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetCurrentJustifiedCheckpoint", args)
	_ = rets
	if !intercept {
		this.Underlying.SetCurrentJustifiedCheckpoint(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetCurrentSyncCommittee(arg0 *solid.SyncCommittee) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetCurrentSyncCommittee", args)
	_ = rets
	if !intercept {
		this.Underlying.SetCurrentSyncCommittee(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetEffectiveBalanceForValidatorAtIndex(arg0 int, arg1 uint64) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetEffectiveBalanceForValidatorAtIndex", args)
	_ = rets
	if !intercept {
		this.Underlying.SetEffectiveBalanceForValidatorAtIndex(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) SetEpochParticipationForValidatorIndex(arg0 bool, arg1 int, arg2 *cltypes.ParticipationFlags) {
	args := []any{
		arg0,
		arg1,
		arg2,
	}
	rets, intercept := this.Handler.Invoke("SetEpochParticipationForValidatorIndex", args)
	_ = rets
	if !intercept {
		this.Underlying.SetEpochParticipationForValidatorIndex(
			arg0,
			arg1,
			arg2,
		)
		return
	}

}

func (this *BeaconStateProxy) SetEth1Data(arg0 *cltypes.Eth1Data) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetEth1Data", args)
	_ = rets
	if !intercept {
		this.Underlying.SetEth1Data(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetEth1DepositIndex(arg0 uint64) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetEth1DepositIndex", args)
	_ = rets
	if !intercept {
		this.Underlying.SetEth1DepositIndex(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetExitEpochForValidatorAtIndex(arg0 int, arg1 uint64) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetExitEpochForValidatorAtIndex", args)
	_ = rets
	if !intercept {
		this.Underlying.SetExitEpochForValidatorAtIndex(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) SetFinalizedCheckpoint(arg0 solid.Checkpoint) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetFinalizedCheckpoint", args)
	_ = rets
	if !intercept {
		this.Underlying.SetFinalizedCheckpoint(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetFork(arg0 *cltypes.Fork) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetFork", args)
	_ = rets
	if !intercept {
		this.Underlying.SetFork(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetInactivityScores(arg0 []uint64) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetInactivityScores", args)
	_ = rets
	if !intercept {
		this.Underlying.SetInactivityScores(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetJustificationBits(arg0 *cltypes.JustificationBits) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetJustificationBits", args)
	_ = rets
	if !intercept {
		this.Underlying.SetJustificationBits(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetLatestBlockHeader(arg0 *cltypes.BeaconBlockHeader) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetLatestBlockHeader", args)
	_ = rets
	if !intercept {
		this.Underlying.SetLatestBlockHeader(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetLatestExecutionPayloadHeader(arg0 *cltypes.Eth1Header) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetLatestExecutionPayloadHeader", args)
	_ = rets
	if !intercept {
		this.Underlying.SetLatestExecutionPayloadHeader(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetNextSyncCommittee(arg0 *solid.SyncCommittee) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetNextSyncCommittee", args)
	_ = rets
	if !intercept {
		this.Underlying.SetNextSyncCommittee(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetNextWithdrawalIndex(arg0 uint64) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetNextWithdrawalIndex", args)
	_ = rets
	if !intercept {
		this.Underlying.SetNextWithdrawalIndex(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetNextWithdrawalValidatorIndex(arg0 uint64) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetNextWithdrawalValidatorIndex", args)
	_ = rets
	if !intercept {
		this.Underlying.SetNextWithdrawalValidatorIndex(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetPreviousEpochAttestations(arg0 *solid.ListSSZ[*solid.PendingAttestation]) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetPreviousEpochAttestations", args)
	_ = rets
	if !intercept {
		this.Underlying.SetPreviousEpochAttestations(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetPreviousEpochParticipationFlags(arg0 *cltypes.ParticipationFlagsList) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetPreviousEpochParticipationFlags", args)
	_ = rets
	if !intercept {
		this.Underlying.SetPreviousEpochParticipationFlags(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetPreviousJustifiedCheckpoint(arg0 solid.Checkpoint) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetPreviousJustifiedCheckpoint", args)
	_ = rets
	if !intercept {
		this.Underlying.SetPreviousJustifiedCheckpoint(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetPreviousStateRoot(arg0 common.Hash) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetPreviousStateRoot", args)
	_ = rets
	if !intercept {
		this.Underlying.SetPreviousStateRoot(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetRandaoMixAt(arg0 int, arg1 common.Hash) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetRandaoMixAt", args)
	_ = rets
	if !intercept {
		this.Underlying.SetRandaoMixAt(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) SetSlashingSegmentAt(arg0 int, arg1 uint64) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetSlashingSegmentAt", args)
	_ = rets
	if !intercept {
		this.Underlying.SetSlashingSegmentAt(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) SetSlot(arg0 uint64) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetSlot", args)
	_ = rets
	if !intercept {
		this.Underlying.SetSlot(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetStateRootAt(arg0 int, arg1 common.Hash) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetStateRootAt", args)
	_ = rets
	if !intercept {
		this.Underlying.SetStateRootAt(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) SetValidatorBalance(arg0 int, arg1 uint64) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetValidatorBalance", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetValidatorBalance(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetValidatorInactivityScore(arg0 int, arg1 uint64) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetValidatorInactivityScore", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetValidatorInactivityScore(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetValidatorIsCurrentMatchingHeadAttester(arg0 int, arg1 bool) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetValidatorIsCurrentMatchingHeadAttester", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetValidatorIsCurrentMatchingHeadAttester(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetValidatorIsCurrentMatchingSourceAttester(arg0 int, arg1 bool) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetValidatorIsCurrentMatchingSourceAttester", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetValidatorIsCurrentMatchingSourceAttester(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetValidatorIsCurrentMatchingTargetAttester(arg0 int, arg1 bool) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetValidatorIsCurrentMatchingTargetAttester", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetValidatorIsCurrentMatchingTargetAttester(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetValidatorIsPreviousMatchingHeadAttester(arg0 int, arg1 bool) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetValidatorIsPreviousMatchingHeadAttester", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetValidatorIsPreviousMatchingHeadAttester(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetValidatorIsPreviousMatchingSourceAttester(arg0 int, arg1 bool) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetValidatorIsPreviousMatchingSourceAttester", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetValidatorIsPreviousMatchingSourceAttester(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetValidatorIsPreviousMatchingTargetAttester(arg0 int, arg1 bool) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetValidatorIsPreviousMatchingTargetAttester", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetValidatorIsPreviousMatchingTargetAttester(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetValidatorMinCurrentInclusionDelayAttestation(arg0 int, arg1 *solid.PendingAttestation) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetValidatorMinCurrentInclusionDelayAttestation", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetValidatorMinCurrentInclusionDelayAttestation(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetValidatorMinPreviousInclusionDelayAttestation(arg0 int, arg1 *solid.PendingAttestation) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetValidatorMinPreviousInclusionDelayAttestation", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetValidatorMinPreviousInclusionDelayAttestation(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetValidatorSlashed(arg0 int, arg1 bool) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetValidatorSlashed", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetValidatorSlashed(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetVersion(arg0 clparams.StateVersion) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SetVersion", args)
	_ = rets
	if !intercept {
		this.Underlying.SetVersion(
			arg0,
		)
		return
	}

}

func (this *BeaconStateProxy) SetWithdrawableEpochForValidatorAtIndex(arg0 int, arg1 uint64) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetWithdrawableEpochForValidatorAtIndex", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SetWithdrawableEpochForValidatorAtIndex(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SetWithdrawalCredentialForValidatorAtIndex(arg0 int, arg1 common.Hash) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SetWithdrawalCredentialForValidatorAtIndex", args)
	_ = rets
	if !intercept {
		this.Underlying.SetWithdrawalCredentialForValidatorAtIndex(
			arg0,
			arg1,
		)
		return
	}

}

func (this *BeaconStateProxy) SlashValidator(arg0 uint64, arg1 *uint64) (ret0 error) {
	args := []any{
		arg0,
		arg1,
	}
	rets, intercept := this.Handler.Invoke("SlashValidator", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SlashValidator(
			arg0,
			arg1,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SlashingSegmentAt(arg0 int) (ret0 uint64) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("SlashingSegmentAt", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.SlashingSegmentAt(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) Slot() (ret0 uint64) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("Slot", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.Slot()
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) StateRoots() (ret0 solid.HashVectorSSZ) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("StateRoots", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.StateRoots()
		return
	}

	var ok bool

	ret0, ok = rets[0].(solid.HashVectorSSZ)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid solid.HashVectorSSZ value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) SyncRewards() (ret0 uint64, ret1 uint64, ret2 error) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("SyncRewards", args)
	_ = rets
	if !intercept {
		ret0, ret1, ret2 = this.Underlying.SyncRewards()
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(uint64)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[1]))
	}
	ret2, ok = rets[2].(error)
	if rets[2] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[2]))
	}
	return ret0, ret1, ret2
}

func (this *BeaconStateProxy) UpgradeToAltair() (ret0 error) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("UpgradeToAltair", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.UpgradeToAltair()
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) UpgradeToBellatrix() (ret0 error) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("UpgradeToBellatrix", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.UpgradeToBellatrix()
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) UpgradeToCapella() (ret0 error) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("UpgradeToCapella", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.UpgradeToCapella()
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) UpgradeToDeneb() (ret0 error) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("UpgradeToDeneb", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.UpgradeToDeneb()
		return
	}

	var ok bool

	ret0, ok = rets[0].(error)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) ValidatorBalance(arg0 int) (ret0 uint64, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorBalance", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorBalance(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorEffectiveBalance(arg0 int) (ret0 uint64, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorEffectiveBalance", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorEffectiveBalance(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorExitEpoch(arg0 int) (ret0 uint64, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorExitEpoch", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorExitEpoch(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorForValidatorIndex(arg0 int) (ret0 solid.Validator, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorForValidatorIndex", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorForValidatorIndex(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(solid.Validator)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid solid.Validator value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorInactivityScore(arg0 int) (ret0 uint64, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorInactivityScore", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorInactivityScore(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorIndexByPubkey(arg0 [48]byte) (ret0 uint64, ret1 bool) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorIndexByPubkey", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorIndexByPubkey(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(bool)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid bool value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorIsCurrentMatchingHeadAttester(arg0 int) (ret0 bool, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorIsCurrentMatchingHeadAttester", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorIsCurrentMatchingHeadAttester(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(bool)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid bool value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorIsCurrentMatchingSourceAttester(arg0 int) (ret0 bool, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorIsCurrentMatchingSourceAttester", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorIsCurrentMatchingSourceAttester(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(bool)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid bool value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorIsCurrentMatchingTargetAttester(arg0 int) (ret0 bool, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorIsCurrentMatchingTargetAttester", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorIsCurrentMatchingTargetAttester(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(bool)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid bool value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorIsPreviousMatchingHeadAttester(arg0 int) (ret0 bool, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorIsPreviousMatchingHeadAttester", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorIsPreviousMatchingHeadAttester(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(bool)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid bool value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorIsPreviousMatchingSourceAttester(arg0 int) (ret0 bool, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorIsPreviousMatchingSourceAttester", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorIsPreviousMatchingSourceAttester(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(bool)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid bool value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorIsPreviousMatchingTargetAttester(arg0 int) (ret0 bool, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorIsPreviousMatchingTargetAttester", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorIsPreviousMatchingTargetAttester(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(bool)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid bool value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorLength() (ret0 int) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("ValidatorLength", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.ValidatorLength()
		return
	}

	var ok bool

	ret0, ok = rets[0].(int)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid int value", rets[0]))
	}
	return ret0
}

func (this *BeaconStateProxy) ValidatorMinCurrentInclusionDelayAttestation(arg0 int) (ret0 *solid.PendingAttestation, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorMinCurrentInclusionDelayAttestation", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorMinCurrentInclusionDelayAttestation(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(*solid.PendingAttestation)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *solid.PendingAttestation value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorMinPreviousInclusionDelayAttestation(arg0 int) (ret0 *solid.PendingAttestation, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorMinPreviousInclusionDelayAttestation", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorMinPreviousInclusionDelayAttestation(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(*solid.PendingAttestation)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid *solid.PendingAttestation value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorPublicKey(arg0 int) (ret0 common.Bytes48, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorPublicKey", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorPublicKey(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(common.Bytes48)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid common.Bytes48 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) ValidatorWithdrawableEpoch(arg0 int) (ret0 uint64, ret1 error) {
	args := []any{
		arg0,
	}
	rets, intercept := this.Handler.Invoke("ValidatorWithdrawableEpoch", args)
	_ = rets
	if !intercept {
		ret0, ret1 = this.Underlying.ValidatorWithdrawableEpoch(
			arg0,
		)
		return
	}

	var ok bool

	ret0, ok = rets[0].(uint64)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid uint64 value", rets[0]))
	}
	ret1, ok = rets[1].(error)
	if rets[1] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid error value", rets[1]))
	}
	return ret0, ret1
}

func (this *BeaconStateProxy) Version() (ret0 clparams.StateVersion) {
	args := []any{}
	rets, intercept := this.Handler.Invoke("Version", args)
	_ = rets
	if !intercept {
		ret0 = this.Underlying.Version()
		return
	}

	var ok bool

	ret0, ok = rets[0].(clparams.StateVersion)
	if rets[0] != nil && !ok {
		panic(fmt.Sprintf("%+v is not a valid clparams.StateVersion value", rets[0]))
	}
	return ret0
}