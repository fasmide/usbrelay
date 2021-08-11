package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "on",
		Short: "turn on a relay",
		RunE: func(cmd *cobra.Command, args []string) error {
			relay, exists := relays[args[0]]
			if !exists {
				return fmt.Errorf("no such relay %s", args[0])
			}

			err := relay.Open()
			if err != nil {
				return fmt.Errorf("unable to open relay: %w", err)
			}

			relay.On()

			relay.Close()

			return nil
		},
	})
}
