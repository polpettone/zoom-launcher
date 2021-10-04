package cmd

import (
	"errors"
	"testing"
)

func Test_convert(t *testing.T) {

	tests := []struct {
		name      string
		httpsLink string
		want      string
		wantErr   error
	}{

		{
			name:      "one",
			httpsLink: "https://zoom.us/j/meetingID?pwd=password",
			want:      "zoommtg://zoom.us?action=join&confno=meetingID&pwd=password",
			wantErr:   nil,
		},

		{
			name:      "invalid zoom link",
			httpsLink: "invalidZoomHttpsLink",
			want:      "",
			wantErr:   errors.New("invalid Zoom link"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, gotErr := convert(tt.httpsLink)
			if got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
			if tt.wantErr != nil && gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("convert() = %v, wantErr %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}
