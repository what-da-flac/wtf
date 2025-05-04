package main

import (
	"testing"
)

func TestHasAudioEnoughQuality(t *testing.T) {
	type args struct {
		bitRate    int
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
				bitRate:    880000,
				minBitrate: 320000,
			},
			want: true,
		},
		{
			name: "sad path",
			args: args{
				bitRate:    192000,
				minBitrate: 320000,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasAudioEnoughQuality(tt.args.bitRate, tt.args.minBitrate); got != tt.want {
				t.Errorf("HasAudioEnoughQuality() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateNumber(t *testing.T) {
	type args struct {
		bitRate    int
		dstBitRate int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "192k to 320k",
			args: args{
				bitRate:    192000,
				dstBitRate: 320000,
			},
			want: 192000,
		},
		{
			name: "320k to 320k",
			args: args{
				bitRate:    320000,
				dstBitRate: 320000,
			},
			want: 320000,
		},
		{
			name: "880k to 320k",
			args: args{
				bitRate:    880000,
				dstBitRate: 320000,
			},
			want: 320000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateNumber(tt.args.bitRate, tt.args.dstBitRate); got != tt.want {
				t.Errorf("CalculateNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
