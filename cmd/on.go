package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

const exampleOn = `  First relay of device 1-4:
    relay on 1-4 1
  All relays of device 1-4:
    relay on 1-4 -a
`

func init() {
	var all bool
	cmd := &cobra.Command{
		Use:     "on BUSID [index of relay]",
		Short:   "turn on a relay",
		Example: exampleOn,
		Args:    cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			relay, exists := relays[args[0]]
			if !exists {
				return fmt.Errorf("no such relay %s", args[0])
			}

			err := relay.Open()
			if err != nil {
				return fmt.Errorf("unable to open relay: %w", err)
			}
			defer relay.Close()

			if all {
				relay.AllOn()
				return nil
			}

			index, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("unable to parse relay index: %w", err)
			}

			relay.On(byte(index))
			relay.Close()

			return nil
		},
	}
	cmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "all relays")
	rootCmd.AddCommand(cmd)
}
