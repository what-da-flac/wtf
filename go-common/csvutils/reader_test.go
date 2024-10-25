package csvutils

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSV(t *testing.T) {
	type args struct {
		csvSource func() (io.Reader, error)
	}
	tests := []struct {
		name    string
		args    args
		want    map[int][]string
		wantErr bool
	}{
		{
			name: "happy path",
			want: map[int][]string{
				0: {
					"FirstName",
					"LastName",
					"Email",
					"Phone",
					"Role",
					"Status",
					"Facilities",
				},
				1: {
					"John",
					"Stark",
					"541-654-6513",
					"CLERK",
					"johnstark@mailinator.com",
					"NEW",
					"f2;f1",
				},
			},
			args: args{
				csvSource: func() (io.Reader, error) {
					source := `FirstName,LastName,Email,Phone,Role,Status,Facilities
John,Stark,johnstark@mailinator.com,541-654-6513,CLERK,NEW,f2;f1`
					reader := bytes.NewBufferString(source)
					return reader, nil
				},
			},
			wantErr: false,
		},
		{
			name: "invalid column count",
			want: nil,
			args: args{
				csvSource: func() (io.Reader, error) {
					source := `FirstName,LastName,Email,Phone,Role,Facilities
John,Stark,johnstark@mailinator.com,541-654-6513,CLERK,f2;f1,`
					reader := bytes.NewBufferString(source)
					return reader, nil
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := func(index int, values []string) error {
				if index == 0 {
					return nil
				}
				expected := tt.want[index]
				assert.ElementsMatch(t, expected, values)
				return nil
			}
			reader, err := tt.args.csvSource()
			assert.NoError(t, err)
			if err := Read(reader, cb); (err != nil) != tt.wantErr {
				t.Errorf("ParseCSVRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
