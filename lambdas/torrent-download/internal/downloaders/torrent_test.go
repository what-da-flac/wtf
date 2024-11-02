package downloaders

import (
	"testing"
)

func TestTorrentDownloader_checkLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "header",
			args: args{
				line: "ID   Done       Have  ETA           Up    Down  Ratio  Status       Name",
			},
			want: false,
		},
		{
			name: "partially",
			args: args{
				line: "1    99%   260.0 MB  0 sec        0.0  1335.0   0.00  Downloading  Coldplay - Moon Music (2024) FLAC [PMEDIA]",
			},
			want: false,
		},
		{
			name: "completed",
			args: args{
				line: "1   100%   261.8 MB  Done         0.0     8.0   0.00  Idle         Coldplay - Moon Music (2024) FLAC [PMEDIA]",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TorrentDownloader{}
			if got := x.checkLine(tt.args.line); got != tt.want {
				t.Errorf("checkLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
