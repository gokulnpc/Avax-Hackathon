// Code generated
// This file is a generated precompile contract test with the skeleton of test functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package icalculator

import (
	"math/big"
	"testing"

	"github.com/ava-labs/subnet-evm/core/state"
	"github.com/ava-labs/subnet-evm/precompile/testutils"
	"github.com/ava-labs/subnet-evm/vmerrs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var (
	_ = vmerrs.ErrOutOfGas
	_ = big.NewInt
	_ = common.Big0
	_ = require.New
)

// These tests are run against the precompile contract directly with
// the given input and expected output. They're just a guide to
// help you write your own tests. These tests are for general cases like
// allowlist, readOnly behaviour, and gas cost. You should write your own
// tests for specific cases.
var (
	tests = map[string]testutils.PrecompileTest{
		"insufficient gas for add should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := AddInput{}
				input, err := PackAdd(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: AddGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
		"insufficient gas for nextTwo should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput *big.Int
				testInput = new(big.Int)
				input, err := PackNextTwo(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: NextTwoGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
		"insufficient gas for repeat should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := RepeatInput{}
				input, err := PackRepeat(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: RepeatGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
	}
)

// TestICalculatorRun tests the Run function of the precompile contract.
func TestICalculatorRun(t *testing.T) {
	// Run tests.
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.Run(t, Module, state.NewTestStateDB(t))
		})
	}
}

func BenchmarkICalculator(b *testing.B) {
	// Benchmark tests.
	for name, test := range tests {
		b.Run(name, func(b *testing.B) {
			test.Bench(b, Module, state.NewTestStateDB(b))
		})
	}
}