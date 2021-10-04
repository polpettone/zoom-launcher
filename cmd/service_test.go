package cmd

import (
	"reflect"
	"testing"
)

const testCase0 = `
zooms:
    room1: link1
    room2: link2
`

func Test_convertRawZoomsToZooms(t *testing.T) {

	tests := []struct {
		name    string
		content []byte
		want    *Zooms
		wantErr bool
	}{

		{
			name: "happy path",
			content: []byte(testCase0),
			want: &Zooms{
				Zooms: map[string]string{"room1":"link1", "room2":"link2"},
			},
			wantErr: false,
		},
		{
			name: "error",
			content: []byte("broken"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertRawZoomsToZooms(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertRawZoomsToZooms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertRawZoomsToZooms() got = %v, want %v", got, tt.want)
			}
		})
	}
}
