package file

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"file_storage_service/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_DownloadFile(t *testing.T) {
	mockFile, _ := os.Open("/test/123")

	type mockFields struct {
		fileUsecase *MockfileProvicer
	}
	type args struct {
		body io.Reader
	}
	tests := []struct {
		name       string
		args       args
		mock       func(mockFields)
		expResp    interface{}
		expCode    int
		expMessage string
	}{
		{
			name: "invalid_JSON_body",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"path": "/test/123"
				`)),
			},
			mock: func(m mockFields) {},
		},
		{
			name: "empty_path_request",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"path": ""
				}`)),
			},
			mock: func(m mockFields) {},
		},
		{
			name: "error_when_calling_usecase",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"path": "test/123"
				}`)),
			},
			mock: func(m mockFields) {
				m.fileUsecase.EXPECT().DonwloadFile("test/123").Return(mockFile, assert.AnError)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFields := mockFields{
				fileUsecase: NewMockfileProvicer(ctrl),
			}
			test.mock(mockFields)

			controller := &Handler{
				file: mockFields.fileUsecase,
			}

			mockWriter := httptest.NewRecorder()
			mockRequest, _ := http.NewRequest("GET", "/file/", test.args.body)

			controller.DonwloadFile(mockWriter, mockRequest)
		})
	}
}

func Test_GetAllFiles(t *testing.T) {
	type mockFields struct {
		fileUsecase *MockfileProvicer
	}
	type args struct {
		body io.Reader
	}
	tests := []struct {
		name       string
		mock       func(mockFields)
		expResp    interface{}
		expCode    int
		expMessage string
	}{
		{
			name: "error_when_calling_get_all_files_usecase",
			mock: func(m mockFields) {
				m.fileUsecase.EXPECT().GetAllFiles().Return([]models.File{}, assert.AnError)
			},
		},
		{
			name: "Success",
			mock: func(m mockFields) {
				m.fileUsecase.EXPECT().GetAllFiles().Return([]models.File{
					{
						FileID:     int64(1),
						URL:        "URL/url",
						Uploader:   "uplader 1",
						UploadTime: "01/01/2023",
					},
				}, nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFields := mockFields{
				fileUsecase: NewMockfileProvicer(ctrl),
			}
			test.mock(mockFields)

			controller := &Handler{
				file: mockFields.fileUsecase,
			}

			mockWriter := httptest.NewRecorder()
			mockRequest, _ := http.NewRequest("GET", "/file/", nil)

			controller.GetAllFiles(mockWriter, mockRequest)
		})
	}
}

func Test_UploadFiles(t *testing.T) {
	type mockFields struct {
		fileUsecase *MockfileProvicer
	}

	tests := []struct {
		name       string
		mock       func(mockFields)
		expResp    interface{}
		expCode    int
		expMessage string
	}{
		// TODO: mock method for parsing file
		{
			name: "error_parsing_file",
			mock: func(m mockFields) {},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFields := mockFields{
				fileUsecase: NewMockfileProvicer(ctrl),
			}
			test.mock(mockFields)

			controller := &Handler{
				file: mockFields.fileUsecase,
			}

			mockWriter := httptest.NewRecorder()
			mockRequest, _ := http.NewRequest("POST", "/file/", nil)

			controller.UploadFile(mockWriter, mockRequest)
		})
	}
}
