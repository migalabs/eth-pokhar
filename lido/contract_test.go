package lido

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLidoContract_GetOperatorsIndexes(t *testing.T) {
	// Create a new instance of LidoContract
	lidoContract, err := NewLidoContract("https://eth-mainnet.g.alchemy.com/v2/jidLpMbJztihH_LSlPrmO8Ffq4sZK110")
	// Check if there was an error
	assert.NoError(t, err)

	// Call the GetOperatorsIndexes function
	indexes, err := lidoContract.GetOperatorsIndexes()

	// Check if there was an error
	assert.NoError(t, err)
	fmt.Println(indexes)
	// Check if the returned indexes are not nil
	assert.NotNil(t, indexes)
}

func TestLidoContract_GetOperatorsData(t *testing.T) {
	// Create a new instance of LidoContract
	lidoContract := &LidoContract{}

	// Create some sample indexes
	indexes := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}

	// Call the GetOperatorsData function
	operatorsData, err := lidoContract.GetOperatorsData(indexes)

	// Check if there was an error
	assert.NoError(t, err)

	// Check if the returned operatorsData are not nil
	assert.NotNil(t, operatorsData)
}

// func TestLidoContract_GetOperatorsKeys(t *testing.T) {
// 	// Create a new instance of LidoContract
// 	lidoContract := &LidoContract{}

// 	// Create some sample operators
// 	operators := []NodeOperator{
// 		{Address: "0x123", Name: "Operator 1"},
// 		{Address: "0x456", Name: "Operator 2"},
// 	}

// 	// Call the GetOperatorsKeys function
// 	operatorsKeys, err := lidoContract.GetOperatorsKeys(operators)

// 	// Check if there was an error
// 	assert.NoError(t, err)

// 	// Check if the returned operatorsKeys are not nil
// 	assert.NotNil(t, operatorsKeys)
// }
