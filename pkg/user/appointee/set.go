package appointee

import (
	"github.com/Layr-Labs/eigenlayer-cli/pkg/internal/common/flags"
	"github.com/Layr-Labs/eigenlayer-cli/pkg/telemetry"
	"github.com/urfave/cli/v2"
)

func SetCmd() *cli.Command {
	setCmd := &cli.Command{
		Name:      "set",
		Usage:     "user appointee set <AccountAddress> <AppointeeAddress> <TargetAddress> <Selector>",
		UsageText: "Grant a user a permission.",
		Description: `
		Grant a user a permission.'.
		`,
		After: telemetry.AfterRunAction(),
		Flags: []cli.Flag{
			&flags.VerboseFlag,
			&AccountAddressFlag,
			&AppointeeAddressFlag,
			&TargetAddressFlag,
			&SelectorFlag,
		},
	}

	return setCmd
}
