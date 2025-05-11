package main

import (
	"testing"
)

func TestHasEnoughQuality(t *testing.T) {
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
			if got := HasEnoughQuality(tt.args.bitRate, tt.args.minBitrate); got != tt.want {
				t.Errorf("HasEnoughQuality() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateMinValue(t *testing.T) {
	type args struct {
		currentBitRate int
		wantedBitRate  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "192k to 320k",
			args: args{
				currentBitRate: 192000,
				wantedBitRate:  320000,
			},
			want: 192000,
		},
		{
			name: "320k to 320k",
			args: args{
				currentBitRate: 320000,
				wantedBitRate:  320000,
			},
			want: 320000,
		},
		{
			name: "880k to 320k",
			args: args{
				currentBitRate: 880000,
				wantedBitRate:  320000,
			},
			want: 320000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMinValue(tt.args.currentBitRate, tt.args.wantedBitRate); got != tt.want {
				t.Errorf("CalculateMinValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
