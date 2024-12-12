package appointee

import (
	"math/big"

	"github.com/Layr-Labs/eigenlayer-cli/pkg/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type canCallConfig struct {
	Network                  string
	RPCUrl                   string
	AccountAddress           gethcommon.Address
	AppointeeAddress         gethcommon.Address
	Target                   gethcommon.Address
	Selector                 [4]byte
	PermissionManagerAddress gethcommon.Address
	ChainID                  *big.Int
	Environment              string
}

type listAppointeesConfig struct {
	Network                  string
	RPCUrl                   string
	AccountAddress           gethcommon.Address
	Target                   gethcommon.Address
	Selector                 [4]byte
	PermissionManagerAddress gethcommon.Address
	ChainID                  *big.Int
	Environment              string
}

type listAppointeePermissionsConfig struct {
	Network                  string
	RPCUrl                   string
	AccountAddress           gethcommon.Address
	AppointeeAddress         gethcommon.Address
	PermissionManagerAddress gethcommon.Address
	ChainID                  *big.Int
	Environment              string
}

type removeConfig struct {
	Network                  string
	RPCUrl                   string
	AccountAddress           gethcommon.Address
	AppointeeAddress         gethcommon.Address
	Target                   gethcommon.Address
	SignerConfig             types.SignerConfig
	Selector                 [4]byte
	PermissionManagerAddress gethcommon.Address
	ChainID                  *big.Int
	Environment              string
}

type setConfig struct {
	Network                  string
	RPCUrl                   string
	AccountAddress           gethcommon.Address
	AppointeeAddress         gethcommon.Address
	Target                   gethcommon.Address
	SignerConfig             types.SignerConfig
	Selector                 [4]byte
	PermissionManagerAddress gethcommon.Address
	ChainID                  *big.Int
	Environment              string
}
