package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewTorrentLine(t *testing.T) {
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
				ID:      "1",
				Eta:     "8 sec",
				Percent: 0.65,
			},
		},
		{
			name: "1%",
			args: args{
				line: `2     1%    1.56 MB  18 min       3.0   245.0   0.00  Up & Down    Drum Samples`,
			},
			want: &TorrentLine{
				ID:      "2",
				Eta:     "18 min",
				Percent: 0.01,
			},
		},
		{
			name: "0%",
			args: args{
				line: `101     0%       None  Unknown      0.0     0.0   0.00  Idle         The Cure - Songs Of A Lost World (2024) [24Bit-96kHz] FLAC [PMEDIA] ⭐️`,
			},
			want: &TorrentLine{
				ID:      "101",
				Eta:     "Unknown",
				Percent: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTorrentLine(tt.args.line)
			assert.Equal(t, tt.want, got)
		})
	}
}
