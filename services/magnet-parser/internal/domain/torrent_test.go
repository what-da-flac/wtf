package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func TestTorrent_ToDomain(t *testing.T) {
	type fields struct {
		Name       string
		Hash       string
		PieceCount int
		PieceSize  string
		TotalSize  string
		Privacy    string
		Trackers   []Tracker
		Files      []File
	}
	tests := []struct {
		name   string
		fields fields
		want   golang.Torrent
	}{
		{
			name: "happy path",
			fields: fields{
				Name:       "Alien Romulus (2024) [1080p] [WEBRip] [5.1] [YTS.MX]",
				Hash:       "579bb4dc2f8aed43ced4e415e90716d17962137f",
				PieceCount: 1025,
				PieceSize:  "2.19 MiB",
				TotalSize:  "2.35 GB",
				Privacy:    "Public torrent",
				Files: []File{
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
			want: golang.Torrent{
				Filename: "",
				Files: []golang.TorrentFile{
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
				Name:       "Alien Romulus (2024) [1080p] [WEBRip] [5.1] [YTS.MX]",
				PieceCount: 1025,
				PieceSize:  "2.19 MiB",
				Privacy:    "Public torrent",
				TotalSize:  "2.35 GB",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Torrent{
				Name:       tt.fields.Name,
				PieceCount: tt.fields.PieceCount,
				PieceSize:  tt.fields.PieceSize,
				TotalSize:  tt.fields.TotalSize,
				Privacy:    tt.fields.Privacy,
				Trackers:   tt.fields.Trackers,
				Files:      tt.fields.Files,
			}
			got := x.ToDomain()
			assert.Equal(t, tt.want, *got)
		})
	}
}
