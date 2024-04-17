package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	ModuleInfo interface {
		Insert(movie *ModuleInfo) error
		Get(id int64) (*ModuleInfo, error)
		Update(movie *ModuleInfo) error
		Delete(id int64) error
	}

	TeacherInfo interface {
		GetAll() ([]*TeacherInfo, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		ModuleInfo:  ModuleInfoModel{DB: db},
		TeacherInfo: TeacherInfoModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		ModuleInfo: MockModuleInfoModel{},
	}
}
