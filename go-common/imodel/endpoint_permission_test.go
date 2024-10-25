package imodel

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePermissions(t *testing.T) {
	type args struct {
		r io.ReadCloser
	}
	tests := []struct {
		name    string
		args    args
		want    []*EndpointPermission
		wantErr bool
	}{
		{
			name:    "no values",
			wantErr: true,
		},
		{
			name: "invalid yaml",
			args: args{
				r: func() io.ReadCloser {
					file, err := os.Open(filepath.Join("test-data", "invalid.yaml"))
					if err != nil {
						t.Error(err)
					}
					return file
				}(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "happy path",
			args: args{
				r: func() io.ReadCloser {
					file, err := os.Open(filepath.Join("test-data", "permissions_example.yaml"))
					if err != nil {
						t.Error(err)
					}
					return file
				}(),
			},
			want: []*EndpointPermission{
				{
					Path:     "/v1/healthz",
					Method:   "get",
					Insecure: true,
				},
				{
					Path:   "/v1/answers",
					Method: "post",
					Roles:  []string{"administrators", "campaign_owners"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePermissions(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePermissions() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestEndpointPermissions_Filter(t *testing.T) {
	type args struct {
		filter func(ep *EndpointPermission) bool
	}
	tests := []struct {
		name string
		x    EndpointPermissions
		args args
		want EndpointPermissions
	}{
		{
			name: "no parameters",
			args: args{},
			want: nil,
		},
		{
			name: "happy path",
			x: EndpointPermissions{
				{
					Path:     "/a/b/c",
					Method:   "get",
					Roles:    []string{"1", "2"},
					Insecure: true,
				},
				{
					Path:   "/d/e/f",
					Method: "post",
					Roles:  []string{"3", "4"},
				},
			},
			args: args{
				filter: func(ep *EndpointPermission) bool {
					return ep.Insecure
				},
			},
			want: EndpointPermissions{
				{
					Path:     "/a/b/c",
					Method:   "get",
					Roles:    []string{"1", "2"},
					Insecure: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.x.Filter(tt.args.filter), "Filter(%v)", tt.args.filter)
		})
	}
}

func TestEndpointPermissions_Allow(t *testing.T) {
	type args struct {
		path   string
		method string
		roles  []string
	}
	tests := []struct {
		name string
		x    EndpointPermissions
		args args
		want bool
	}{
		{
			name: "no values",
			want: false,
		},
		{
			name: "get /v1/questions matches",
			x: EndpointPermissions{
				{
					Path:   "/v1/questions",
					Method: http.MethodGet,
					Roles:  []string{"users", "administrators"},
				},
			},
			args: args{
				path:   "/v1/questions",
				method: http.MethodGet,
				roles:  []string{"administrators"},
			},
			want: true,
		},
		{
			name: "get /v1/questions roles dont match",
			x: EndpointPermissions{
				{
					Path:   "/v1/questions",
					Method: http.MethodGet,
				},
			},
			args: args{
				path:   "/v1/questions",
				method: http.MethodGet,
				roles:  []string{"administrators"},
			},
			want: false,
		},
		{
			name: "post /v1/questions no match",
			x: EndpointPermissions{
				{
					Path:     "/v1/questions/{id}/answers",
					Method:   http.MethodGet,
					Roles:    []string{"users", "administrators"},
					Insecure: false,
				},
			},
			args: args{
				path:   "/v1/questions/abc-123/answers",
				method: http.MethodPost,
				roles:  []string{"administrators"},
			},
			want: false,
		},
		{
			name: "post /v1/questions insecure",
			x: EndpointPermissions{
				{
					Path:     "/v1/questions/{id}/answers",
					Method:   http.MethodGet,
					Insecure: true,
				},
			},
			args: args{
				path:   "/v1/questions/abc-123/answers",
				method: http.MethodGet,
			},
			want: false,
		},
		{
			name: "get /v1/campaigns/{id} real path",
			x: EndpointPermissions{
				{
					Path:     "/v1/campaigns/{id}",
					Method:   http.MethodGet,
					Insecure: true,
					Roles:    []string{"users", "campaign_owners"},
				},
			},
			args: args{
				path:   "/v1/campaigns/f2ccc9d2-bc5b-4977-b4de-b417b835d853",
				method: http.MethodGet,
				roles:  []string{"campaign_owners"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.x.Allow(tt.args.path, tt.args.method, tt.args.roles...)
			assert.Equal(t, tt.want, got, "continue bad value")
		})
	}
}

func TestEndpointPermissions_Len(t *testing.T) {
	tests := []struct {
		name string
		x    EndpointPermissions
		want int
	}{
		{
			name: "no values",
			want: 0,
		},
		{
			name: "happy path",
			x: EndpointPermissions{
				{
					Path:     "/v1/questions",
					Method:   http.MethodGet,
					Insecure: true,
				},
				{
					Path:     "/v1/questions",
					Method:   http.MethodPost,
					Insecure: true,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.x.Len(), "Len()")
		})
	}
}
