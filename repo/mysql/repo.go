package mysql

import (
	"app/Go/model"
	"app/Go/repo"
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB

}
func (s *studentRepository) GetOneByID (ctx context.Context, id int) model.Student {
	s.db
}

var instance studentRepository

func NewStudentRepository(db *gorm.DB) repo.StudentRepo {
	if instance.db == nil {
		instance.db = db

	}
	return instance
}


