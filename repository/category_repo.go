package repository

import (
	"github.com/rikyhidayat21/gomet-unit-test/entity"
)

// kontrak
type CategoryRepository interface {
	FindById(id string) *entity.Category
}