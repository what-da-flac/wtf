package downloaders

import (
	"testing"

	"github.com/what-da-flac/wtf/services/torrent-download/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestTorrentDownloader_readTorrentLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want *model.TorrentLine
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
			want: &model.TorrentLine{
				ID:      "1",
				Percent: 0.65,
			},
		},
		{
			name: "1%",
			args: args{
				line: `2     1%    1.56 MB  18 min       3.0   245.0   0.00  Up & Down    Drum Samples`,
			},
			want: &model.TorrentLine{
				ID:      "2",
				Percent: 0.01,
			},
		},
		{
			name: "0%",
			args: args{
				line: `101     0%       None  Unknown      0.0     0.0   0.00  Idle         The Cure - Songs Of A Lost World (2024) [24Bit-96kHz] FLAC [PMEDIA] ⭐️`,
			},
			want: &model.TorrentLine{
				ID:      "101",
				Percent: 0,
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
