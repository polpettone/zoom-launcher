package cmd

import (
	"fmt"
	"strings"
)

func convert(httpLink string) string {
	split :=  strings.Split(httpLink, "/")
	zoomHost := split[2]
	meetingID := strings.Split(split[4], "?")[0]
	password := strings.Split(split[4], "=")[1]

	zoommtg := fmt.Sprintf("zoommtg://%s?action=join&confno=%s&pwd=%s", zoomHost, meetingID, password)
	return zoommtg
}
