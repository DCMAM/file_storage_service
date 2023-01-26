package file

import (
	"file_storage_service/internal/models"
	"mime/multipart"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUseCase_DownloadFile(t *testing.T) {
	mockFile, _ := os.Open("/test/123")

	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr error
		mock    func(mockRepoProvider *MockfileDBProvider, mockFilerepoProvider *MockfileProvider)
	}{
		{
			name: "error",
			args: args{
				path: "test/123",
			},
			mock: func(mockRepoProvider *MockfileDBProvider, mockFilerepoProvider *MockfileProvider) {
				mockFilerepoProvider.EXPECT().DonwloadFile("test/123").Return(mockFile, assert.AnError)
			},
			want:    mockFile,
			wantErr: assert.AnError,
		},
		{
			name: "success",
			args: args{
				path: "test/123",
			},
			mock: func(mockRepoProvider *MockfileDBProvider, mockFilerepoProvider *MockfileProvider) {
				mockFilerepoProvider.EXPECT().DonwloadFile("test/123").Return(mockFile, nil)
			},
			want:    mockFile,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockfileDBProvider := NewMockfileDBProvider(ctrl)
			mockfileProvider := NewMockfileProvider(ctrl)

			test.mock(mockfileDBProvider, mockfileProvider)

			uc := Usecase{
				fileDB: mockfileDBProvider,
				file:   mockfileProvider,
			}

			got, gotError := uc.DonwloadFile(test.args.path)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, gotError)
		})
	}
}

func TestUseCase_GetAllFiles(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.File
		wantErr error
		mock    func(mockRepoProvider *MockfileDBProvider)
	}{
		{
			name: "got_error_when_get_data",
			mock: func(mockRepoProvider *MockfileDBProvider) {
				mockRepoProvider.EXPECT().GetAll().Return([]models.File{}, assert.AnError)
			},
			want:    nil,
			wantErr: assert.AnError,
		},
		{
			name: "successfully_get_data",
			mock: func(mockRepoProvider *MockfileDBProvider) {
				mockRepoProvider.EXPECT().GetAll().Return([]models.File{
					{
						FileID:     int64(1),
						URL:        "URL",
						Uploader:   "uploader 1",
						UploadTime: "01/01/2023",
					},
				}, nil)
			},
			want: []models.File{
				{
					FileID:     int64(1),
					URL:        "URL",
					Uploader:   "uploader 1",
					UploadTime: "01/01/2023",
				},
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockfileDBProvider := NewMockfileDBProvider(ctrl)

			test.mock(mockfileDBProvider)

			uc := Usecase{
				fileDB: mockfileDBProvider,
			}

			got, gotError := uc.GetAllFiles()
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, gotError)
		})
	}
}

func TestUseCase_UploadFile(t *testing.T) {
	type args struct {
		file     multipart.File
		username string
	}

	// TODO: find a way to mock file in parameter

	tests := []struct {
		name    string
		args    args
		wantErr error
		mock    func(mockRepoProvider *MockfileDBProvider, mockFilerepoProvider *MockfileProvider)
	}{
		{
			name: "got_errow_when_upload_file",
			args: args{},
			mock: func(mockRepoProvider *MockfileDBProvider, mockFilerepoProvider *MockfileProvider) {
				mockFilerepoProvider.EXPECT().UploadFile(gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "got_errow_when_set_file",
			args: args{
				username: "name of user",
			},
			mock: func(mockRepoProvider *MockfileDBProvider, mockFilerepoProvider *MockfileProvider) {
				mockFilerepoProvider.EXPECT().UploadFile(gomock.Any(), gomock.Any()).Return(nil)
				mockRepoProvider.EXPECT().SetFile(gomock.Any(), "name of user").Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "success",
			args: args{
				username: "name of user",
			},
			mock: func(mockRepoProvider *MockfileDBProvider, mockFilerepoProvider *MockfileProvider) {
				mockFilerepoProvider.EXPECT().UploadFile(gomock.Any(), gomock.Any()).Return(nil)
				mockRepoProvider.EXPECT().SetFile(gomock.Any(), "name of user").Return(nil)
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockfileDBProvider := NewMockfileDBProvider(ctrl)
			mockfileProvider := NewMockfileProvider(ctrl)

			test.mock(mockfileDBProvider, mockfileProvider)

			uc := Usecase{
				fileDB: mockfileDBProvider,
				file:   mockfileProvider,
			}

			gotError := uc.UploadFile(test.args.file, test.args.username)
			assert.Equal(t, test.wantErr, gotError)
		})
	}
}
