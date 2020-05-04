package repository_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	uuid "github.com/farkroft/go.uuid"

	"gitlab.com/auth-service/application/request"

	"github.com/jinzhu/gorm"
	"gitlab.com/auth-service/application/repository"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRegisterUserRepositoryShouldSuccessAndReturnUserModel(t *testing.T) {
	db, sqlMocks, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error when stub db connection %v", err)
	}
	dbMock, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("error when mock postgres %v", err)
	}
	repo := repository.UserRepo{DB: dbMock}

	userReq := request.UserRequest{
		Username: "fajarar77@gmail.com",
		Password: "password",
	}

	id := uuid.NewV4()

	sqlMocks.ExpectBegin()
	expectedQuery := regexp.QuoteMeta("INSERT INTO \"users\" (\"id\",\"created_at\",\"updated_at\",\"deleted_at\",\"username\",\"password\") VALUES ($1,$2,$3,$4,$5,$6) RETURNING \"users\".\"id\"")
	row := sqlmock.NewRows([]string{"id"}).AddRow(id.String())
	sqlMocks.ExpectQuery(expectedQuery).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), nil, userReq.Username, userReq.Password).WillReturnRows(row)
	sqlMocks.ExpectCommit()

	user, err := repo.RegisterUser(userReq)
	if err != nil {
		t.Errorf("repo return err %v", err)
	}
	err = sqlMocks.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations %v", err)
	}
	assert.Equal(t, user.ID, id)
	assert.Equal(t, user.Username, userReq.Username)
	assert.Equal(t, user.Password, userReq.Password)
}
