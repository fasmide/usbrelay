package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

const exampleOff = `  First relay of device 1-4:
    relay off 1-4 1
  All relays of device 1-4:
    relay off 1-4 -a
`

func init() {
	var all bool
	cmd := &cobra.Command{
		Use:     "off",
		Short:   "turn off a relay",
		Example: exampleOff,
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
				relay.AllOff()
				return nil
			}

			index, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("unable to parse relay index: %w", err)
			}

			relay.Off(byte(index))
			relay.Close()

			return nil
		},
	}

	cmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "all relays")
	rootCmd.AddCommand(cmd)
}
