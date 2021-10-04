package cmd

import "testing"

func Test_convert(t *testing.T) {

	tests := []struct {
		name string
		httpsLink string
		want string
	}{
		{
			name: "one",
			httpsLink: "https://zoom.us/j/meetingID?pwd=password",
			want: "zoommtg://zoom.us?action=join&confno=meetingID&pwd=password",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.httpsLink); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
