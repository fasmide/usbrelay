package cmd

import (
	"fmt"
	"os"

	"github.com/fasmide/usbrelay/relay"
	"github.com/spf13/cobra"
)

var relays map[string]*relay.Relay

var rootCmd = &cobra.Command{
	Use:   "relay",
	Short: "Relay turns on or off cheap relays",
}

// Execute root command
func Execute() {
	var err error

	relays, err = relay.ByPort()
	if err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
