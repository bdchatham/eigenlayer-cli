package operator

import (
	"math/big"

	"github.com/Layr-Labs/eigenlayer-cli/pkg/types"
	"github.com/ethereum/go-ethereum/common"
)

type DeregisterConfig struct {
	avsAddress               common.Address
	operatorSetIds           []uint32
	operatorAddress          common.Address
	network                  string
	environment              string
	broadcast                bool
	rpcUrl                   string
	chainID                  *big.Int
	signerConfig             *types.SignerConfig
	output                   string
	outputType               string
	delegationManagerAddress common.Address
	isSilent                 bool
}

type RegisterConfig struct {
	avsAddress               common.Address
	operatorSetIds           []uint32
	operatorAddress          common.Address
	network                  string
	environment              string
	broadcast                bool
	rpcUrl                   string
	chainID                  *big.Int
	signerConfig             *types.SignerConfig
	output                   string
	outputType               string
	delegationManagerAddress common.Address
	isSilent                 bool
}
