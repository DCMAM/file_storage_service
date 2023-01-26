package database

// import (
// 	"context"
// 	"file_storage_service/internal/models"
// 	"testing"

// 	sqlmock "github.com/DATA-DOG/go-sqlmock"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestRepository_GetUserByUsername(t *testing.T) {
// 	type mockFields struct {
// 		db sqlmock.Sqlmock
// 	}

// 	tests := []struct {
// 		name         string
// 		username     string
// 		mock         func(mock mockFields)
// 		expectedResp models.User
// 		expectedErr  error
// 	}{
// 		{
// 			name:     "success",
// 			username: "username123",
// 			mock: func(mock mockFields) {
// 				rows := sqlmock.NewRows([]string{
// 					"user_id", "name", "username", "password",
// 				}).AddRow(
// 					int64(1), "name of user", "username123", "#password123",
// 				)

// 				mock.db.ExpectQuery(`
// 					SELECT
// 						user_id,
// 						name,
// 						username,
// 						password
// 					FROM users
// 					WHERE username = ?
// 				`).WillReturnRows(rows)
// 			},
// 			expectedResp: models.User{
// 				UserID:   int64(1),
// 				Name:     "name of user",
// 				Username: "username123",
// 				Password: "#password123",
// 			},
// 			expectedErr: nil,
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)

// 			dbMocked, mockDB, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
// 			mockSQLDB := sqldb.NewFromDB(dbMocked, dbMocked, "postgre")

// 			defer func() {
// 				ctrl.Finish()
// 				_ = dbMocked.Close()
// 			}()

// 			mockFields := mockFields{
// 				db:    mockDB,
// 				infra: NewMockinfraProvider(ctrl),
// 			}

// 			test.mock(mockFields)
// 			repo := &Repository{
// 				db:    mockSQLDB,
// 				infra: mockFields.infra,
// 			}

// 			got, err := repo.GetArticleRatingByArticleIDAndUserID(context.Background(), test.param)
// 			assert.Equal(t, test.expectedResp, got)
// 			assert.Equal(t, test.expectedErr, err)
// 			assert.Nil(t, mockDB.ExpectationsWereMet())
// 		})
// 	}
// }
