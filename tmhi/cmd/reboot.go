package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot the local T-Mobile gateway",
	RunE: func(cmd *cobra.Command, args []string) error {

		confirmed := false
		survey.AskOne(&survey.Confirm{
			Message: "Rebooting the gateway means losing Internet connectivity and this\ntool no longer being able to connect for a minute or two. Continue?",
		},
			&confirmed)

		if !confirmed {
			return nil
		}

		fmt.Println("Rebooting the TMHI gateway...")

		_, err := gatewayPost("gateway/reset?set=reboot", true, []byte(""))
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("rebooted")
		} else if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(rebootCmd)
}
