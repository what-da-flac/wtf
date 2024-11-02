package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
	"github.com/what-da-flac/wtf/go-common/amazon"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/lambdas/torrent-download/internal/environment"
	"github.com/what-da-flac/wtf/lambdas/torrent-download/internal/interfaces"
	"github.com/what-da-flac/wtf/lambdas/torrent-download/mocks"
	"github.com/what-da-flac/wtf/openapi/models"
	"go.uber.org/zap"
)

func Test_uploadResultToS3(t *testing.T) {
	zl, err := zap.NewDevelopment()
	require.NoError(t, err)
	logger := zl.Sugar()
	type args struct {
		logger    ifaces.Logger
		uploader  interfaces.Uploader
		config    *environment.Config
		torrent   *models.Torrent
		targetDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				logger: logger,
				uploader: &mocks.UploaderMock{
					UploadFunc: func(_ *os.File, _ string, key string, _ amazon.Content) error {
						assert.Equal(t, "31a3f3c5-45e7-4100-880b-d97fc19f908e/dummy.txt", key)
						return nil
					},
				},
				config: environment.New(),
				torrent: &models.Torrent{
					Id:   "31a3f3c5-45e7-4100-880b-d97fc19f908e",
					Name: "name-of-torrent-file",
				},
				targetDir: "test-data",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uploadResultToS3(tt.args.logger, tt.args.uploader, tt.args.config, tt.args.torrent, tt.args.targetDir)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
