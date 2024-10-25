package locations

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceKm(t *testing.T) {
	homePoint := Point{
		Lat: 19.5839973,
		Lng: -99.2354309,
	}
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name   string
		args   args
		wantKm float64
	}{
		{
			name: "home to drugstore",
			args: args{
				p1: homePoint,
				p2: Point{
					Lat: 19.5863285,
					Lng: -99.2329723,
				},
			},
			wantKm: 0,
		},
		{
			name: "home to school",
			args: args{
				p1: homePoint,
				p2: Point{
					Lat: 19.5659852,
					Lng: -99.2233111,
				},
			},
			wantKm: 2,
		},
		{
			name: "home to ottawa",
			args: args{
				p1: homePoint,
				p2: Point{
					Lat: 45.2496825,
					Lng: -76.1298907,
				},
			},
			wantKm: 3557,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DistanceKm(tt.args.p1, tt.args.p2)
			got = math.Floor(got)
			assert.Equal(t, tt.wantKm, got)
		})
	}
}
