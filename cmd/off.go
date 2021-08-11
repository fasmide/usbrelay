package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "off",
		Short: "turn off a relay",
		RunE: func(cmd *cobra.Command, args []string) error {
			relay, exists := relays[args[0]]
			if !exists {
				return fmt.Errorf("no such relay %s", args[0])
			}

			err := relay.Open()
			if err != nil {
				return fmt.Errorf("unable to open relay: %w", err)
			}

			relay.Off()

			relay.Close()

			return nil
		},
	})
}
