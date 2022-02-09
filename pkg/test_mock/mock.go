package test_mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func NewMock() (*gorm.DB, sqlmock.Sqlmock) {
	// mockDB, mock, err := sqlmock.New()
	// db := sql.NewDb(mockDB, "sqlmock")

	// if err != nil {
	// 	log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }

	// return db, mock
	return nil, nil
}
