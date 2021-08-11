package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "state",
		Short: "show states of relays",
		RunE: func(cmd *cobra.Command, args []string) error {
			for port, relay := range relays {
				err := relay.Open()
				if err != nil {
					return fmt.Errorf("unable to open port %s: %w", port, err)
				}
				defer relay.Close()

				state, err := relay.State()
				if err != nil {
					return fmt.Errorf("unable to ask for state: %w", err)
				}

				fmt.Printf("%s: %+v", port, state)
			}
			return nil
		},
	})
}
