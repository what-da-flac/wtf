package pgrepo

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/what-da-flac/wtf/openapi/models"

	"github.com/stretchr/testify/assert"
)

func TestUserDto_toProto(t *testing.T) {
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	type fields struct {
		Id        string
		Created   time.Time
		Email     string
		Name      string
		Image     *string
		LastLogin *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.User
	}{
		{
			name: "happy path",
			fields: fields{
				Id:      "1",
				Created: now,
				Email:   "2@test.com",
				Name:    "3",
				Image: func() *string {
					val := "4"
					return &val
				}(),
				LastLogin: &now,
			},
			want: &models.User{
				Id:        "1",
				Created:   now,
				Email:     "2@test.com",
				Name:      "3",
				Image:     aws.String("4"),
				LastLogin: now,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &UserDto{
				Id:        tt.fields.Id,
				Created:   tt.fields.Created,
				Email:     tt.fields.Email,
				Name:      tt.fields.Name,
				ImageUrl:  tt.fields.Image,
				LastLogin: tt.fields.LastLogin,
			}
			got := x.toProto()
			assert.EqualValues(t, tt.want.Id, got.Id)
			assert.EqualValues(t, tt.want.Email, got.Email)
			assert.EqualValues(t, tt.want.Name, got.Name)
			assert.EqualValues(t, tt.want.Image, got.Image)
			assert.EqualValues(t, tt.want.LastLogin.Truncate(time.Second), got.LastLogin.Truncate(time.Second))
		})
	}
}
