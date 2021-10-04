package cmd

import (
	"errors"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start <zoom meeting name>",
		Short: "starts a zoom call",
		Run: func(cmd *cobra.Command, args []string) {
			err := handleStartCommand(args[0])
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), err.Error())
			}
		},
	}
}

func handleStartCommand(meetingName string) error {

	zoomsFile := viper.GetString(ZoomsFile)
	zooms, err := LoadZooms(zoomsFile)

	if err != nil {
		return err
	}

	link, found := zooms.Entries[meetingName]

	if found {
		z, err := convert(link)

		if err != nil {
			return err
		}

		err = open.Run(z)
		if err != nil {
			return errors.New(fmt.Sprintf("%s could not be open with zoom", z))
		}
	} else {
		return errors.New(fmt.Sprintf("%s not Found", meetingName))
	}
	return nil
}


func init() {
	startCmd := StartCmd()
	rootCmd.AddCommand(startCmd)
}
