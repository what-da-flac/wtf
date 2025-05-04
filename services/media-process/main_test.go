package main

import (
	"testing"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func TestHasAudioEnoughQuality(t *testing.T) {
	type args struct {
		audio      golang.Audio
		minBitrate int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path",
			args: args{
				audio: golang.Audio{
					BitRate: 880000,
				},
				minBitrate: 320000,
			},
			want: true,
		},
		{
			name: "sad path",
			args: args{
				audio: golang.Audio{
					BitRate: 192000,
				},
				minBitrate: 320000,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasAudioEnoughQuality(tt.args.audio, tt.args.minBitrate); got != tt.want {
				t.Errorf("HasAudioEnoughQuality() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateBitrate(t *testing.T) {
	type args struct {
		audio   golang.Audio
		bitRate int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "192k to 320k",
			args: args{
				audio: golang.Audio{
					BitRate: 192000,
				},
				bitRate: 320000,
			},
			want: 192000,
		},
		{
			name: "320k to 320k",
			args: args{
				audio: golang.Audio{
					BitRate: 320000,
				},
				bitRate: 320000,
			},
			want: 320000,
		},
		{
			name: "880k to 320k",
			args: args{
				audio: golang.Audio{
					BitRate: 880000,
				},
				bitRate: 320000,
			},
			want: 320000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateBitrate(tt.args.audio, tt.args.bitRate); got != tt.want {
				t.Errorf("CalculateBitrate() = %v, want %v", got, tt.want)
			}
		})
	}
}
