package repository_test

import (
	"log"
	"regexp"
	"testing"
	"time"

	uuid "github.com/farkroft/go.uuid"

	"gitlab.com/auth-service/application/request"

	"github.com/jinzhu/gorm"
	"gitlab.com/auth-service/application/repository"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRegisterUserRepositoryShouldSuccessAndReturnUserModel(t *testing.T) {
	db, sqlMocks, _ := sqlmock.New()
	dbMock, _ := gorm.Open("postgres", db)
	repo := repository.UserRepo{DB: dbMock}

	var Now = func() time.Time {
		str := "2002-10-02T10:00:00+07:00"
		res, _ := time.Parse(time.RFC3339, str)
		return res
	}

	userReq := request.UserRequest{
		Username: "fajarar77@gmail.com",
		Password: "password",
	}

	id := uuid.NewV4()
	sqlMocks.ExpectBegin()
	expectedQuery := regexp.QuoteMeta("INSERT INTO \"users\" (\"id\",\"created_at\",\"updated_at\",\"deleted_at\",\"username\",\"password\") VALUES ($1,$2,$3,$4,$5,$6) RETURNING \"users\".\"id\"")
	row := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "username", "password"}).AddRow(id, Now(), Now(), nil, userReq.Username, userReq.Password)
	sqlMocks.ExpectQuery(expectedQuery).WithArgs(id, Now(), Now(), nil, userReq.Username, userReq.Password).WillReturnRows(row)
	sqlMocks.ExpectCommit()

	_, err := repo.RegisterUser(userReq)
	if err != nil {
		log.Println("====REPO====")
		log.Println(err)
	}
	err = sqlMocks.ExpectationsWereMet()
}
