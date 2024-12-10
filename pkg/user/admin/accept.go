package admin

import (
	"sort"

	"github.com/Layr-Labs/eigenlayer-cli/pkg/internal/common/flags"
	"github.com/Layr-Labs/eigenlayer-cli/pkg/telemetry"
	"github.com/urfave/cli/v2"
)

func AcceptCmd() *cli.Command {
	acceptCmd := &cli.Command{
		Name:      "accept-admin",
		Usage:     "user admin accept-admin --account-address <AccountAddress>",
		UsageText: "Accepts a user to become admin who is currently pending admin acceptance.",
		Description: `
		Accepts a user to become admin who is currently pending admin acceptance.
		`,
		After: telemetry.AfterRunAction(),
		Flags: acceptFlags(),
	}

	return acceptCmd
}

func acceptFlags() []cli.Flag {
	cmdFlags := []cli.Flag{
		&flags.VerboseFlag,
		&AccountAddressFlag,
	}
	sort.Sort(cli.FlagsByName(cmdFlags))
	return append(cmdFlags, flags.GetSignerFlags()...)
}