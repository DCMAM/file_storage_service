package local_storage

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUseCase_DownloadFile(t *testing.T) {
	mockFile, mockErr := os.Open("test/123")

	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr error
	}{
		{
			name: "error",
			args: args{
				path: "test/123",
			},
			want:    mockFile,
			wantErr: mockErr,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := Repository{}

			got, gotError := repo.DonwloadFile(test.args.path)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, gotError)
		})
	}
}
