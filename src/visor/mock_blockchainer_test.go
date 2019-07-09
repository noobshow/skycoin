// Code generated by mockery v1.0.0. DO NOT EDIT.

package visor

import blockdb "github.com/skycoin/skycoin/src/visor/blockdb"
import cipher "github.com/skycoin/skycoin/src/cipher"
import coin "github.com/skycoin/skycoin/src/coin"
import dbutil "github.com/skycoin/skycoin/src/visor/dbutil"
import mock "github.com/stretchr/testify/mock"
import params "github.com/skycoin/skycoin/src/params"

// MockBlockchainer is an autogenerated mock type for the Blockchainer type
type MockBlockchainer struct {
	mock.Mock
}

// ExecuteBlock provides a mock function with given fields: tx, sb
func (_m *MockBlockchainer) ExecuteBlock(tx *dbutil.Tx, sb *coin.SignedBlock) error {
	ret := _m.Called(tx, sb)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, *coin.SignedBlock) error); ok {
		r0 = rf(tx, sb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBlocks provides a mock function with given fields: tx, seqs
func (_m *MockBlockchainer) GetBlocks(tx *dbutil.Tx, seqs []uint64) ([]coin.SignedBlock, error) {
	ret := _m.Called(tx, seqs)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, []uint64) []coin.SignedBlock); ok {
		r0 = rf(tx, seqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, []uint64) error); ok {
		r1 = rf(tx, seqs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlocksInRange provides a mock function with given fields: tx, start, end
func (_m *MockBlockchainer) GetBlocksInRange(tx *dbutil.Tx, start uint64, end uint64) ([]coin.SignedBlock, error) {
	ret := _m.Called(tx, start, end)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, uint64, uint64) []coin.SignedBlock); ok {
		r0 = rf(tx, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, uint64, uint64) error); ok {
		r1 = rf(tx, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGenesisBlock provides a mock function with given fields: tx
func (_m *MockBlockchainer) GetGenesisBlock(tx *dbutil.Tx) (*coin.SignedBlock, error) {
	ret := _m.Called(tx)

	var r0 *coin.SignedBlock
	if rf, ok := ret.Get(0).(func(*dbutil.Tx) *coin.SignedBlock); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBlocks provides a mock function with given fields: tx, n
func (_m *MockBlockchainer) GetLastBlocks(tx *dbutil.Tx, n uint64) ([]coin.SignedBlock, error) {
	ret := _m.Called(tx, n)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, uint64) []coin.SignedBlock); ok {
		r0 = rf(tx, n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, uint64) error); ok {
		r1 = rf(tx, n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSignedBlockByHash provides a mock function with given fields: tx, hash
func (_m *MockBlockchainer) GetSignedBlockByHash(tx *dbutil.Tx, hash cipher.SHA256) (*coin.SignedBlock, error) {
	ret := _m.Called(tx, hash)

	var r0 *coin.SignedBlock
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, cipher.SHA256) *coin.SignedBlock); ok {
		r0 = rf(tx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, cipher.SHA256) error); ok {
		r1 = rf(tx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSignedBlockBySeq provides a mock function with given fields: tx, seq
func (_m *MockBlockchainer) GetSignedBlockBySeq(tx *dbutil.Tx, seq uint64) (*coin.SignedBlock, error) {
	ret := _m.Called(tx, seq)

	var r0 *coin.SignedBlock
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, uint64) *coin.SignedBlock); ok {
		r0 = rf(tx, seq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, uint64) error); ok {
		r1 = rf(tx, seq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Head provides a mock function with given fields: tx
func (_m *MockBlockchainer) Head(tx *dbutil.Tx) (*coin.SignedBlock, error) {
	ret := _m.Called(tx)

	var r0 *coin.SignedBlock
	if rf, ok := ret.Get(0).(func(*dbutil.Tx) *coin.SignedBlock); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeadSeq provides a mock function with given fields: tx
func (_m *MockBlockchainer) HeadSeq(tx *dbutil.Tx) (uint64, bool, error) {
	ret := _m.Called(tx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*dbutil.Tx) uint64); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(*dbutil.Tx) bool); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*dbutil.Tx) error); ok {
		r2 = rf(tx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Len provides a mock function with given fields: tx
func (_m *MockBlockchainer) Len(tx *dbutil.Tx) (uint64, error) {
	ret := _m.Called(tx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*dbutil.Tx) uint64); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBlock provides a mock function with given fields: tx, txns, currentTime
func (_m *MockBlockchainer) NewBlock(tx *dbutil.Tx, txns coin.Transactions, currentTime uint64) (*coin.Block, error) {
	ret := _m.Called(tx, txns, currentTime)

	var r0 *coin.Block
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, coin.Transactions, uint64) *coin.Block); ok {
		r0 = rf(tx, txns, currentTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, coin.Transactions, uint64) error); ok {
		r1 = rf(tx, txns, currentTime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Time provides a mock function with given fields: tx
func (_m *MockBlockchainer) Time(tx *dbutil.Tx) (uint64, error) {
	ret := _m.Called(tx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*dbutil.Tx) uint64); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionFee provides a mock function with given fields: tx, hours
func (_m *MockBlockchainer) TransactionFee(tx *dbutil.Tx, hours uint64) coin.FeeCalculator {
	ret := _m.Called(tx, hours)

	var r0 coin.FeeCalculator
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, uint64) coin.FeeCalculator); ok {
		r0 = rf(tx, hours)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(coin.FeeCalculator)
		}
	}

	return r0
}

// Unspent provides a mock function with given fields:
func (_m *MockBlockchainer) Unspent() blockdb.UnspentPooler {
	ret := _m.Called()

	var r0 blockdb.UnspentPooler
	if rf, ok := ret.Get(0).(func() blockdb.UnspentPooler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(blockdb.UnspentPooler)
		}
	}

	return r0
}

// VerifyBlock provides a mock function with given fields: tx, sb
func (_m *MockBlockchainer) VerifyBlock(tx *dbutil.Tx, sb *coin.SignedBlock) error {
	ret := _m.Called(tx, sb)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, *coin.SignedBlock) error); ok {
		r0 = rf(tx, sb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyBlockTxnConstraints provides a mock function with given fields: tx, txn
func (_m *MockBlockchainer) VerifyBlockTxnConstraints(tx *dbutil.Tx, txn coin.Transaction) error {
	ret := _m.Called(tx, txn)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, coin.Transaction) error); ok {
		r0 = rf(tx, txn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifySingleTxnHardConstraints provides a mock function with given fields: tx, txn, signed
func (_m *MockBlockchainer) VerifySingleTxnHardConstraints(tx *dbutil.Tx, txn coin.Transaction, signed TxnSignedFlag) error {
	ret := _m.Called(tx, txn, signed)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, coin.Transaction, TxnSignedFlag) error); ok {
		r0 = rf(tx, txn, signed)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifySingleTxnSoftHardConstraints provides a mock function with given fields: tx, txn, distParams, verifyParams, signed
func (_m *MockBlockchainer) VerifySingleTxnSoftHardConstraints(tx *dbutil.Tx, txn coin.Transaction, distParams params.Distribution, verifyParams params.VerifyTxn, signed TxnSignedFlag) (*coin.SignedBlock, coin.UxArray, error) {
	ret := _m.Called(tx, txn, distParams, verifyParams, signed)

	var r0 *coin.SignedBlock
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, coin.Transaction, params.Distribution, params.VerifyTxn, TxnSignedFlag) *coin.SignedBlock); ok {
		r0 = rf(tx, txn, distParams, verifyParams, signed)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.SignedBlock)
		}
	}

	var r1 coin.UxArray
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, coin.Transaction, params.Distribution, params.VerifyTxn, TxnSignedFlag) coin.UxArray); ok {
		r1 = rf(tx, txn, distParams, verifyParams, signed)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(coin.UxArray)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*dbutil.Tx, coin.Transaction, params.Distribution, params.VerifyTxn, TxnSignedFlag) error); ok {
		r2 = rf(tx, txn, distParams, verifyParams, signed)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
