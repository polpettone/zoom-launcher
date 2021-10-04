package cmd

import (
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


func handleStartCommand(meetingName string)  error {

	zoomsFile := viper.GetString(ZoomsFile)
	zooms, err := loadZooms(zoomsFile)

	if err != nil {
		return  err
	}

	var keywords []string
	for k, _ := range zooms.Entries {
		keywords = append(keywords, k)
	}
	link := zooms.Entries[meetingName]
	z := convert(link[0])

	err = open.Run(z)

	if err != nil {
		return err
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
