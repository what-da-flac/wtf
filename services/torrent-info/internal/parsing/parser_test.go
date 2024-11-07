package parsing

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/what-da-flac/wtf/services/torrent-info/internal/domain"
)

func TestParseTorrent(t *testing.T) {
	type args struct {
		metadataFilename string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Torrent
		wantErr bool
	}{
		{
			name: "alien-romulus",
			args: args{
				metadataFilename: filepath.Join("test-data", "alien-romulus.txt"),
			},
			want: domain.Torrent{
				Name:       "Alien Romulus (2024) [1080p] [WEBRip] [5.1] [YTS.MX]",
				PieceCount: 1025,
				PieceSize:  "2.19 MiB",
				TotalSize:  "2.35 GB",
				Privacy:    "Public torrent",
				Files: []domain.File{
					{
						FileName: "Alien Romulus (2024) [1080p] [WEBRip] [5.1] [YTS.MX]/Alien.Romulus.2024.1080p.WEBRip.x264.AAC5.1-[YTS.MX].mp4",
						FileSize: "2.35 GB",
					},
					{
						FileName: "Alien Romulus (2024) [1080p] [WEBRip] [5.1] [YTS.MX]/Alien.Romulus.2024.1080p.WEBRip.x264.AAC5.1-[YTS.MX].srt",
						FileSize: "73.15 kB",
					},
					{
						FileName: "Alien Romulus (2024) [1080p] [WEBRip] [5.1] [YTS.MX]/Subs/English (CC).eng.srt",
						FileSize: "106.4 kB",
					},
					{
						FileName: "Alien Romulus (2024) [1080p] [WEBRip] [5.1] [YTS.MX]/Subs/English.eng.srt",
						FileSize: "73.15 kB",
					},
					{
						FileName: "Alien Romulus (2024) [1080p] [WEBRip] [5.1] [YTS.MX]/Subs/Español (Latinoamérica).spa.srt",
						FileSize: "75.67 kB",
					},
					{
						FileName: "Alien Romulus (2024) [1080p] [WEBRip] [5.1] [YTS.MX]/Subs/Français (Canada).fre.srt",
						FileSize: "78.53 kB",
					},
					{
						FileName: "Alien Romulus (2024) [1080p] [WEBRip] [5.1] [YTS.MX]/YTSProxies.com.txt",
						FileSize: "0.58 kB",
					},
					{
						FileName: "Alien Romulus (2024) [1080p] [WEBRip] [5.1] [YTS.MX]/www.YTS.MX.jpg",
						FileSize: "53.23 kB",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "number-of-the-beast",
			args: args{
				metadataFilename: filepath.Join("test-data", "iron-maiden-number-of-the-beast.txt"),
			},
			want: domain.Torrent{
				Name:       "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube",
				PieceCount: 885,
				PieceSize:  "1.00 MiB",
				TotalSize:  "928.0 MB",
				Privacy:    "Public torrent",
				Files: []domain.File{
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/00.-Iron Maiden - The Number Of The Beast.m3u",
						FileSize: "0.26 kB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/01.-Invaders.flac",
						FileSize: "79.80 MB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/02.-Children Of The Damned.flac",
						FileSize: "104.2 MB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/03.-The Prisoner.flac",
						FileSize: "138.6 MB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/04.-22 Acacia Avenue (The Continuing Saga Of Charlotte The Harlot).flac",
						FileSize: "149.9 MB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/05.-The Number Of The Beast.flac",
						FileSize: "112.5 MB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/06.-Run To The Hills.flac",
						FileSize: "90.68 MB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/07.-Gangland.flac",
						FileSize: "88.14 MB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/08.-Hallowed Be Thy Name.flac",
						FileSize: "162.0 MB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/DR12.txt",
						FileSize: "1.31 kB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/Spectrograms/01.-Invaders.flac.png",
						FileSize: "228.6 kB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/Spectrograms/02.-Children Of The Damned.flac.png",
						FileSize: "232.3 kB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/Spectrograms/03.-The Prisoner.flac.png",
						FileSize: "223.0 kB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/Spectrograms/04.-22 Acacia Avenue (The Continuing Saga Of Charlotte The Harlot).flac.png",
						FileSize: "223.3 kB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/Spectrograms/05.-The Number Of The Beast.flac.png",
						FileSize: "212.1 kB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/Spectrograms/06.-Run To The Hills.flac.png",
						FileSize: "213.2 kB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/Spectrograms/07.-Gangland.flac.png",
						FileSize: "213.7 kB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/Spectrograms/08.-Hallowed Be Thy Name.flac.png",
						FileSize: "205.8 kB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/folder.jpg",
						FileSize: "354.9 kB",
					},
					{
						FileName: "Iron Maiden - The Number Of The Beast (1982) (UK EMI 100 PBTHAL Vinyl 24-96) [FLAC] vtwin88cube/lineage.txt",
						FileSize: "0.56 kB",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(tt.args.metadataFilename)
			require.NoError(t, err)
			got, err := ParseTorrent(string(data))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTorrent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.PieceCount, got.PieceCount)
			assert.Equal(t, tt.want.PieceSize, got.PieceSize)
			assert.Equal(t, tt.want.TotalSize, got.TotalSize)
			assert.Equal(t, tt.want.Privacy, got.Privacy)
			assert.ElementsMatch(t, tt.want.Files, got.Files)
		})
	}
}
