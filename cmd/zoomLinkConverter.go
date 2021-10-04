package cmd

import (
	"errors"
	"fmt"
	"strings"
)

func convert(httpLink string) (string, error) {
	split :=  strings.Split(httpLink, "/")

	if len(split) < 5 {
		return "", errors.New("invalid Zoom link")
	}

	zoomHost := split[2]
	meetingID := strings.Split(split[4], "?")[0]
	password := strings.Split(split[4], "=")[1]

	zoommtg := fmt.Sprintf("zoommtg://%s?action=join&confno=%s&pwd=%s", zoomHost, meetingID, password)
	return zoommtg, nil
}
