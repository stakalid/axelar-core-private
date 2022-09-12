// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	multisig "github.com/axelarnetwork/axelar-core/x/multisig/exported"
	github_com_axelarnetwork_axelar_core_x_nexus_exported "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	"github.com/axelarnetwork/axelar-core/x/tss/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"sync"
)

// Ensure, that StakingKeeperMock does implement types.StakingKeeper.
// If this is not the case, regenerate this file with moq.
var _ types.StakingKeeper = &StakingKeeperMock{}

// StakingKeeperMock is a mock implementation of types.StakingKeeper.
//
// 	func TestSomethingThatUsesStakingKeeper(t *testing.T) {
//
// 		// make and configure a mocked types.StakingKeeper
// 		mockedStakingKeeper := &StakingKeeperMock{
// 			ValidatorFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, addr github_com_cosmos_cosmos_sdk_types.ValAddress) stakingtypes.ValidatorI {
// 				panic("mock out the Validator method")
// 			},
// 		}
//
// 		// use mockedStakingKeeper in code that requires types.StakingKeeper
// 		// and then make assertions.
//
// 	}
type StakingKeeperMock struct {
	// ValidatorFunc mocks the Validator method.
	ValidatorFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, addr github_com_cosmos_cosmos_sdk_types.ValAddress) stakingtypes.ValidatorI

	// calls tracks calls to the methods.
	calls struct {
		// Validator holds details about calls to the Validator method.
		Validator []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Addr is the addr argument value.
			Addr github_com_cosmos_cosmos_sdk_types.ValAddress
		}
	}
	lockValidator sync.RWMutex
}

// Validator calls ValidatorFunc.
func (mock *StakingKeeperMock) Validator(ctx github_com_cosmos_cosmos_sdk_types.Context, addr github_com_cosmos_cosmos_sdk_types.ValAddress) stakingtypes.ValidatorI {
	if mock.ValidatorFunc == nil {
		panic("StakingKeeperMock.ValidatorFunc: method is nil but StakingKeeper.Validator was just called")
	}
	callInfo := struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Addr github_com_cosmos_cosmos_sdk_types.ValAddress
	}{
		Ctx:  ctx,
		Addr: addr,
	}
	mock.lockValidator.Lock()
	mock.calls.Validator = append(mock.calls.Validator, callInfo)
	mock.lockValidator.Unlock()
	return mock.ValidatorFunc(ctx, addr)
}

// ValidatorCalls gets all the calls that were made to Validator.
// Check the length with:
//     len(mockedStakingKeeper.ValidatorCalls())
func (mock *StakingKeeperMock) ValidatorCalls() []struct {
	Ctx  github_com_cosmos_cosmos_sdk_types.Context
	Addr github_com_cosmos_cosmos_sdk_types.ValAddress
} {
	var calls []struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Addr github_com_cosmos_cosmos_sdk_types.ValAddress
	}
	mock.lockValidator.RLock()
	calls = mock.calls.Validator
	mock.lockValidator.RUnlock()
	return calls
}

// Ensure, that SnapshotterMock does implement types.Snapshotter.
// If this is not the case, regenerate this file with moq.
var _ types.Snapshotter = &SnapshotterMock{}

// SnapshotterMock is a mock implementation of types.Snapshotter.
//
// 	func TestSomethingThatUsesSnapshotter(t *testing.T) {
//
// 		// make and configure a mocked types.Snapshotter
// 		mockedSnapshotter := &SnapshotterMock{
// 			GetOperatorFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, proxy github_com_cosmos_cosmos_sdk_types.AccAddress) github_com_cosmos_cosmos_sdk_types.ValAddress {
// 				panic("mock out the GetOperator method")
// 			},
// 		}
//
// 		// use mockedSnapshotter in code that requires types.Snapshotter
// 		// and then make assertions.
//
// 	}
type SnapshotterMock struct {
	// GetOperatorFunc mocks the GetOperator method.
	GetOperatorFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, proxy github_com_cosmos_cosmos_sdk_types.AccAddress) github_com_cosmos_cosmos_sdk_types.ValAddress

	// calls tracks calls to the methods.
	calls struct {
		// GetOperator holds details about calls to the GetOperator method.
		GetOperator []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Proxy is the proxy argument value.
			Proxy github_com_cosmos_cosmos_sdk_types.AccAddress
		}
	}
	lockGetOperator sync.RWMutex
}

// GetOperator calls GetOperatorFunc.
func (mock *SnapshotterMock) GetOperator(ctx github_com_cosmos_cosmos_sdk_types.Context, proxy github_com_cosmos_cosmos_sdk_types.AccAddress) github_com_cosmos_cosmos_sdk_types.ValAddress {
	if mock.GetOperatorFunc == nil {
		panic("SnapshotterMock.GetOperatorFunc: method is nil but Snapshotter.GetOperator was just called")
	}
	callInfo := struct {
		Ctx   github_com_cosmos_cosmos_sdk_types.Context
		Proxy github_com_cosmos_cosmos_sdk_types.AccAddress
	}{
		Ctx:   ctx,
		Proxy: proxy,
	}
	mock.lockGetOperator.Lock()
	mock.calls.GetOperator = append(mock.calls.GetOperator, callInfo)
	mock.lockGetOperator.Unlock()
	return mock.GetOperatorFunc(ctx, proxy)
}

// GetOperatorCalls gets all the calls that were made to GetOperator.
// Check the length with:
//     len(mockedSnapshotter.GetOperatorCalls())
func (mock *SnapshotterMock) GetOperatorCalls() []struct {
	Ctx   github_com_cosmos_cosmos_sdk_types.Context
	Proxy github_com_cosmos_cosmos_sdk_types.AccAddress
} {
	var calls []struct {
		Ctx   github_com_cosmos_cosmos_sdk_types.Context
		Proxy github_com_cosmos_cosmos_sdk_types.AccAddress
	}
	mock.lockGetOperator.RLock()
	calls = mock.calls.GetOperator
	mock.lockGetOperator.RUnlock()
	return calls
}

// Ensure, that NexusMock does implement types.Nexus.
// If this is not the case, regenerate this file with moq.
var _ types.Nexus = &NexusMock{}

// NexusMock is a mock implementation of types.Nexus.
//
// 	func TestSomethingThatUsesNexus(t *testing.T) {
//
// 		// make and configure a mocked types.Nexus
// 		mockedNexus := &NexusMock{
// 			GetChainsFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context) []github_com_axelarnetwork_axelar_core_x_nexus_exported.Chain {
// 				panic("mock out the GetChains method")
// 			},
// 		}
//
// 		// use mockedNexus in code that requires types.Nexus
// 		// and then make assertions.
//
// 	}
type NexusMock struct {
	// GetChainsFunc mocks the GetChains method.
	GetChainsFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context) []github_com_axelarnetwork_axelar_core_x_nexus_exported.Chain

	// calls tracks calls to the methods.
	calls struct {
		// GetChains holds details about calls to the GetChains method.
		GetChains []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
		}
	}
	lockGetChains sync.RWMutex
}

// GetChains calls GetChainsFunc.
func (mock *NexusMock) GetChains(ctx github_com_cosmos_cosmos_sdk_types.Context) []github_com_axelarnetwork_axelar_core_x_nexus_exported.Chain {
	if mock.GetChainsFunc == nil {
		panic("NexusMock.GetChainsFunc: method is nil but Nexus.GetChains was just called")
	}
	callInfo := struct {
		Ctx github_com_cosmos_cosmos_sdk_types.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetChains.Lock()
	mock.calls.GetChains = append(mock.calls.GetChains, callInfo)
	mock.lockGetChains.Unlock()
	return mock.GetChainsFunc(ctx)
}

// GetChainsCalls gets all the calls that were made to GetChains.
// Check the length with:
//     len(mockedNexus.GetChainsCalls())
func (mock *NexusMock) GetChainsCalls() []struct {
	Ctx github_com_cosmos_cosmos_sdk_types.Context
} {
	var calls []struct {
		Ctx github_com_cosmos_cosmos_sdk_types.Context
	}
	mock.lockGetChains.RLock()
	calls = mock.calls.GetChains
	mock.lockGetChains.RUnlock()
	return calls
}

// Ensure, that MultiSigKeeperMock does implement types.MultiSigKeeper.
// If this is not the case, regenerate this file with moq.
var _ types.MultiSigKeeper = &MultiSigKeeperMock{}

// MultiSigKeeperMock is a mock implementation of types.MultiSigKeeper.
//
// 	func TestSomethingThatUsesMultiSigKeeper(t *testing.T) {
//
// 		// make and configure a mocked types.MultiSigKeeper
// 		mockedMultiSigKeeper := &MultiSigKeeperMock{
// 			GetActiveKeyIDsFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, chainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) []multisig.KeyID {
// 				panic("mock out the GetActiveKeyIDs method")
// 			},
// 			GetKeyFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, keyID multisig.KeyID) (multisig.Key, bool) {
// 				panic("mock out the GetKey method")
// 			},
// 		}
//
// 		// use mockedMultiSigKeeper in code that requires types.MultiSigKeeper
// 		// and then make assertions.
//
// 	}
type MultiSigKeeperMock struct {
	// GetActiveKeyIDsFunc mocks the GetActiveKeyIDs method.
	GetActiveKeyIDsFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, chainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) []multisig.KeyID

	// GetKeyFunc mocks the GetKey method.
	GetKeyFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, keyID multisig.KeyID) (multisig.Key, bool)

	// calls tracks calls to the methods.
	calls struct {
		// GetActiveKeyIDs holds details about calls to the GetActiveKeyIDs method.
		GetActiveKeyIDs []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// ChainName is the chainName argument value.
			ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
		}
		// GetKey holds details about calls to the GetKey method.
		GetKey []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// KeyID is the keyID argument value.
			KeyID multisig.KeyID
		}
	}
	lockGetActiveKeyIDs sync.RWMutex
	lockGetKey          sync.RWMutex
}

// GetActiveKeyIDs calls GetActiveKeyIDsFunc.
func (mock *MultiSigKeeperMock) GetActiveKeyIDs(ctx github_com_cosmos_cosmos_sdk_types.Context, chainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) []multisig.KeyID {
	if mock.GetActiveKeyIDsFunc == nil {
		panic("MultiSigKeeperMock.GetActiveKeyIDsFunc: method is nil but MultiSigKeeper.GetActiveKeyIDs was just called")
	}
	callInfo := struct {
		Ctx       github_com_cosmos_cosmos_sdk_types.Context
		ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
	}{
		Ctx:       ctx,
		ChainName: chainName,
	}
	mock.lockGetActiveKeyIDs.Lock()
	mock.calls.GetActiveKeyIDs = append(mock.calls.GetActiveKeyIDs, callInfo)
	mock.lockGetActiveKeyIDs.Unlock()
	return mock.GetActiveKeyIDsFunc(ctx, chainName)
}

// GetActiveKeyIDsCalls gets all the calls that were made to GetActiveKeyIDs.
// Check the length with:
//     len(mockedMultiSigKeeper.GetActiveKeyIDsCalls())
func (mock *MultiSigKeeperMock) GetActiveKeyIDsCalls() []struct {
	Ctx       github_com_cosmos_cosmos_sdk_types.Context
	ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
} {
	var calls []struct {
		Ctx       github_com_cosmos_cosmos_sdk_types.Context
		ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
	}
	mock.lockGetActiveKeyIDs.RLock()
	calls = mock.calls.GetActiveKeyIDs
	mock.lockGetActiveKeyIDs.RUnlock()
	return calls
}

// GetKey calls GetKeyFunc.
func (mock *MultiSigKeeperMock) GetKey(ctx github_com_cosmos_cosmos_sdk_types.Context, keyID multisig.KeyID) (multisig.Key, bool) {
	if mock.GetKeyFunc == nil {
		panic("MultiSigKeeperMock.GetKeyFunc: method is nil but MultiSigKeeper.GetKey was just called")
	}
	callInfo := struct {
		Ctx   github_com_cosmos_cosmos_sdk_types.Context
		KeyID multisig.KeyID
	}{
		Ctx:   ctx,
		KeyID: keyID,
	}
	mock.lockGetKey.Lock()
	mock.calls.GetKey = append(mock.calls.GetKey, callInfo)
	mock.lockGetKey.Unlock()
	return mock.GetKeyFunc(ctx, keyID)
}

// GetKeyCalls gets all the calls that were made to GetKey.
// Check the length with:
//     len(mockedMultiSigKeeper.GetKeyCalls())
func (mock *MultiSigKeeperMock) GetKeyCalls() []struct {
	Ctx   github_com_cosmos_cosmos_sdk_types.Context
	KeyID multisig.KeyID
} {
	var calls []struct {
		Ctx   github_com_cosmos_cosmos_sdk_types.Context
		KeyID multisig.KeyID
	}
	mock.lockGetKey.RLock()
	calls = mock.calls.GetKey
	mock.lockGetKey.RUnlock()
	return calls
}
