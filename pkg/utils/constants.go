package utils

import "github.com/ethereum/go-ethereum/common"

const (
	EmojiCheckMark = "✅"
	EmojiCrossMark = "❌"
	EmojiWarning   = "⚠️"
	EmojiInfo      = "ℹ️"
	EmojiWait      = "⏳"
	EmojiLink      = "🔗"
	EmojiInternet  = "🌐"

	MainnetChainId = 1
	HoleskyChainId = 17000
	SepoliaChainId = 11155111
	AnvilChainId   = 31337

	MainnetNetworkName = "mainnet"
	HoleskyNetworkName = "holesky"
	SepoliaNetworkName = "seopolia"
	AnvilNetworkName   = "anvil"
	UnknownNetworkName = "unknown"

	CallData string = "calldata"
	Pretty   string = "pretty"
	Json     string = "json"
)

var (
	ZeroAddress = common.HexToAddress("")
)
