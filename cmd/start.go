package cmd

import (
	"errors"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	zooms, err := loadZooms(zoomsFile)

	if err != nil {
		return err
	}

	link, found := zooms.Entries[meetingName]

	if found {
		z := convert(link[0])
		err = open.Run(z)
		if err != nil {
			return err
		}
	} else {
		return errors.New(fmt.Sprintf("%s not Found", meetingName))
	}
	return nil
}

type Zooms struct {
	Entries map[string][]string
}

func loadZooms(file string) (*Zooms, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	zooms := &Zooms{}
	err = yaml.Unmarshal(content, &zooms)
	if err != nil {
		return nil, err
	}
	return zooms, nil
}

func init() {
	startCmd := StartCmd()
	rootCmd.AddCommand(startCmd)
}
