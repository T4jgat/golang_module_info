package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	ModuleInfo interface {
		Insert(movie *ModuleInfo) error
		Get(id int64) (*ModuleInfo, error)
		Update(movie *ModuleInfo) error
		Delete(id int64) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		ModuleInfo: ModuleInfoModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		ModuleInfo: MockModuleInfoModel{},
	}
}
