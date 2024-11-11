package downloaders

import (
	"testing"

	"github.com/what-da-flac/wtf/go-common/loggers"

	"github.com/stretchr/testify/assert"
)

func TestTorrentDownloader_checkLine(t *testing.T) {
	logger := loggers.MustNewDevelopmentLogger()
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
			x := NewTorrentDownloader(logger, 0)
			if got := x.checkLine(tt.args.line); got != tt.want {
				t.Errorf("checkLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTorrentDownloader_readTorrentLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want *TorrentLine
	}{
		{
			name: "header",
			args: args{
				line: `ID   Done       Have  ETA           Up    Down  Ratio  Status       Name`,
			},
			want: nil,
		},
		{
			name: "65%",
			args: args{
				line: `     1    65%   644.3 MB  8 sec        0.0  43245.0   0.00  Downloading  The Cure - Songs Of A Lost World (2024) [24Bit-96kHz] FLAC [PMEDIA]`,
			},
			want: &TorrentLine{
				ID: "1",
			},
		},
		{
			name: "1%",
			args: args{
				line: `2     1%    1.56 MB  18 min       3.0   245.0   0.00  Up & Down    Drum Samples`,
			},
			want: &TorrentLine{
				ID: "2",
			},
		},
		{
			name: "0%",
			args: args{
				line: `101     0%       None  Unknown      0.0     0.0   0.00  Idle         The Cure - Songs Of A Lost World (2024) [24Bit-96kHz] FLAC [PMEDIA] ⭐️`,
			},
			want: &TorrentLine{
				ID: "101",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TorrentDownloader{}
			got := x.readTorrentLine(tt.args.line)
			assert.Equal(t, tt.want, got)
		})
	}
}
