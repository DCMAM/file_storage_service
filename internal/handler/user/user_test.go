package user

import (
	"bytes"
	"file_storage_service/internal/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Login(t *testing.T) {
	type mockFields struct {
		userUsecase *MockUserProvider
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
					"username": "",
					"password": ""
				`)),
			},
			mock: func(m mockFields) {},
		},
		{
			name: "invalid_username",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"username": "",
					"password": ""
				}`)),
			},
			mock: func(m mockFields) {},
		},
		{
			name: "invalid_password",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"username": "username123",
					"password": ""
				}`)),
			},
			mock: func(m mockFields) {},
		},
		{
			name: "error_when_calling_login_usecase",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"username": "username123",
					"password": "#password123"
				}`)),
			},
			mock: func(m mockFields) {
				m.userUsecase.EXPECT().Login("username123", "#password123").Return(models.User{
					UserID:   int64(1),
					Name:     "name of user",
					Username: "username123",
					Password: "#password123",
					Token:    "_eajdwgwgwe",
				}, assert.AnError)
			},
		},
		{
			name: "Success",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"username": "username123",
					"password": "#password123"
				}`)),
			},
			mock: func(m mockFields) {
				m.userUsecase.EXPECT().Login("username123", "#password123").Return(models.User{
					UserID:   int64(1),
					Name:     "name of user",
					Username: "username123",
					Password: "#password123",
					Token:    "_eajdwgwgwe",
				}, nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFields := mockFields{
				userUsecase: NewMockUserProvider(ctrl),
			}
			test.mock(mockFields)

			controller := &Handler{
				user: mockFields.userUsecase,
			}

			mockWriter := httptest.NewRecorder()
			mockRequest, _ := http.NewRequest("POST", "/user/login", test.args.body)

			controller.Login(mockWriter, mockRequest)
		})
	}
}

func Test_Register(t *testing.T) {
	type mockFields struct {
		userUsecase *MockUserProvider
	}
	type args struct {
		body io.Reader
	}
	tests := []struct {
		name string
		args args
		mock func(mockFields)
	}{
		{
			name: "invalid_JSON_body",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"name": "",
					"username": "",
					"password": ""
				`)),
			},
			mock: func(m mockFields) {},
		},
		{
			name: "invalid_name",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"name": "",
					"username": "",
					"password": ""
				}`)),
			},
			mock: func(m mockFields) {},
		},
		{
			name: "invalid_username",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"name": "name of user",
					"username": "",
					"password": ""
				}`)),
			},
			mock: func(m mockFields) {},
		},
		{
			name: "invalid_password",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"name": "name of user",
					"username": "username123",
					"password": ""
				}`)),
			},
			mock: func(m mockFields) {},
		},
		{
			name: "error_when_calling_register_usecase",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"name": "name of user",
					"username": "username123",
					"password": "#password123"
				}`)),
			},
			mock: func(m mockFields) {
				m.userUsecase.EXPECT().Register("name of user", "username123", "#password123").Return(assert.AnError)
			},
		},
		{
			name: "Success",
			args: args{
				body: bytes.NewBuffer([]byte(`{
					"name": "name of user",
					"username": "username123",
					"password": "#password123"
				}`)),
			},
			mock: func(m mockFields) {
				m.userUsecase.EXPECT().Register("name of user", "username123", "#password123").Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFields := mockFields{
				userUsecase: NewMockUserProvider(ctrl),
			}
			test.mock(mockFields)

			controller := &Handler{
				user: mockFields.userUsecase,
			}

			mockWriter := httptest.NewRecorder()
			mockRequest, _ := http.NewRequest("POST", "/user/register", test.args.body)

			controller.Register(mockWriter, mockRequest)
		})
	}
}
