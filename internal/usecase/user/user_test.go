package user

import (
	"errors"
	"file_storage_service/internal/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUseCase_Login(t *testing.T) {
	mockPassword, _ := bcrypt.GenerateFromPassword([]byte("#password123"), bcrypt.DefaultCost)

	type args struct {
		name     string
		username string
		password string
	}

	tests := []struct {
		name    string
		args    args
		want    models.User
		wantErr error
		mock    func(mockUserServiceProvider *MockUserProvider)
	}{
		{
			name: "got_error_when_get_user_data_by_username",
			args: args{
				name:     "name",
				username: "name123",
				password: "#password123",
			},
			mock: func(mockUserServiceProvider *MockUserProvider) {
				mockUserServiceProvider.EXPECT().GetUserByUsername("name123").Return(models.User{}, assert.AnError)
			},
			want:    models.User{},
			wantErr: assert.AnError,
		},
		{
			name: "got_error_no_match_data",
			args: args{
				name:     "name",
				username: "name123",
				password: "#password123",
			},
			mock: func(mockUserServiceProvider *MockUserProvider) {
				mockUserServiceProvider.EXPECT().GetUserByUsername("name123").Return(models.User{}, nil)
			},
			want:    models.User{},
			wantErr: errors.New("username didn't match with any data in database"),
		},
		{
			name: "got_error_password_did_not_match",
			args: args{
				name:     "name",
				username: "name123",
				password: "#password123",
			},
			mock: func(mockUserServiceProvider *MockUserProvider) {
				mockUserServiceProvider.EXPECT().GetUserByUsername("name123").Return(models.User{
					UserID:   1,
					Name:     "name of user",
					Username: "user123",
					Password: "not_matching_password",
				}, nil)
			},
			want:    models.User{},
			wantErr: errors.New("password didn't match"),
		},
		{
			name: "successfully_login",
			args: args{
				name:     "name",
				username: "name123",
				password: "#password123",
			},
			mock: func(mockUserServiceProvider *MockUserProvider) {
				mockUserServiceProvider.EXPECT().GetUserByUsername("name123").Return(models.User{
					UserID:   1,
					Name:     "name of user",
					Username: "user123",
					Password: string(mockPassword),
				}, nil)
			},
			want: models.User{ // not testing returned data because helper service for token not using depedency injection yet
				UserID:   1,
				Name:     "name of user",
				Username: "user123",
				Password: "#password123",
				Token:    gomock.Any().String(),
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepoProvider := NewMockUserProvider(ctrl)

			test.mock(mockUserRepoProvider)

			uc := Usecase{
				user: mockUserRepoProvider,
			}

			_, gotError := uc.Login(test.args.username, test.args.password)
			// not testing returned data because helper service for token not using depedency injection yet
			// assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, gotError)
		})
	}
}

func TestUseCase_Register(t *testing.T) {
	type args struct {
		name     string
		username string
		password string
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		mock    func(mockUserServiceProvider *MockUserProvider)
	}{
		{
			name: "got_error_when_get_user_data_by_username",
			args: args{
				name:     "name",
				username: "name123",
				password: "#password123",
			},
			mock: func(mockUserServiceProvider *MockUserProvider) {
				mockUserServiceProvider.EXPECT().GetUserByUsername("name123").Return(models.User{}, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "got_error_that_user_already_exist",
			args: args{
				name:     "name",
				username: "name123",
				password: "#password123",
			},
			mock: func(mockUserServiceProvider *MockUserProvider) {
				mockUserServiceProvider.EXPECT().GetUserByUsername("name123").Return(models.User{
					Username: "name123",
				}, nil)
			},
			wantErr: errors.New("User already exist"),
		},
		{
			name: "got_error_when_register_user",
			args: args{
				name:     "name",
				username: "name123",
				password: "#password123",
			},
			mock: func(mockUserServiceProvider *MockUserProvider) {
				mockUserServiceProvider.EXPECT().GetUserByUsername("name123").Return(models.User{}, nil)
				mockUserServiceProvider.EXPECT().RegisterUser("name", "name123", gomock.Any()).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "successfully_register",
			args: args{
				name:     "name",
				username: "name123",
				password: "#password123",
			},
			mock: func(mockUserServiceProvider *MockUserProvider) {
				mockUserServiceProvider.EXPECT().GetUserByUsername("name123").Return(models.User{}, nil)
				mockUserServiceProvider.EXPECT().RegisterUser("name", "name123", gomock.Any()).Return(nil)
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepoProvider := NewMockUserProvider(ctrl)

			test.mock(mockUserRepoProvider)

			uc := Usecase{
				user: mockUserRepoProvider,
			}

			got := uc.Register(test.args.name, test.args.username, test.args.password)
			assert.Equal(t, test.wantErr, got)
		})
	}
}
