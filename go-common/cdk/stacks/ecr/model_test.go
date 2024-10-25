package ecr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel_Validate(t *testing.T) {
	type fields struct {
		EmptyOnDelete   bool
		Mutable         bool
		Name            string
		RemoveOnDestroy bool
		UseDefaults     bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    Model
		wantErr bool
	}{
		{
			name: "defaults",
			fields: fields{
				Name:        "test",
				UseDefaults: true,
			},
			want: Model{
				EmptyOnDelete:   true,
				Mutable:         true,
				Name:            "test",
				RemoveOnDestroy: true,
				UseDefaults:     true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Model{
				EmptyOnDelete:   tt.fields.EmptyOnDelete,
				Mutable:         tt.fields.Mutable,
				Name:            tt.fields.Name,
				RemoveOnDestroy: tt.fields.RemoveOnDestroy,
				UseDefaults:     tt.fields.UseDefaults,
			}
			err := x.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, *x)
		})
	}
}
