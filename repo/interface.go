package repo

import (
	"app/Go/model"
	"context"
)

type StudentRepo interface {
	GetOneByID(ctx context.Context, id int) model.Student
	GetAll(ctx context.Context) []model.Student
	InsertOne(ctx context.Context, c model.Student) error
}
