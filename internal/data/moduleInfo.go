package data

import (
	"database/sql"
	"errors"
	"github.com/T4jgat/module_info/internal/validator"
	"github.com/lib/pq"
	"time"
)

type ModuleInfo struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"-"`
	ModuleName     string    `json:"module_name"`
	ModuleDuration int32     `json:"module_duration"`
	ExamType       []string  `json:"exam_type"`
	Version        int32     `json:"version"`
}

type ModuleInfoModel struct {
	DB *sql.DB
}

func (m ModuleInfoModel) Insert(modelInfo *ModuleInfo) error {
	query := `
		INSERT INTO module_info (module_name, module_duration, exam_type)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, version
		`

	args := []any{modelInfo.ModuleName, modelInfo.ModuleDuration, pq.Array(modelInfo.ExamType)}

	return m.DB.QueryRow(query, args...).Scan(&modelInfo.ID, &modelInfo.CreatedAt, &modelInfo.Version)
}

func (m ModuleInfoModel) Get(id int64) (*ModuleInfo, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, updated_at, module_name, module_duration, exam_type, version
		FROM module_info
		WHERE id=$1
		`
	var moduleInfo ModuleInfo

	err := m.DB.QueryRow(query, id).Scan(
		&moduleInfo.ID,
		&moduleInfo.CreatedAt,
		&moduleInfo.UpdatedAt,
		&moduleInfo.ModuleName,
		&moduleInfo.ModuleDuration,
		pq.Array(&moduleInfo.ExamType),
		&moduleInfo.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &moduleInfo, nil
}

func (m ModuleInfoModel) Update(moduleInfo *ModuleInfo) error {
	query := `
		UPDATE module_info
		SET module_name = $1, module_duration = $2, exam_type = $3, version = version + 1
		WHERE id = $4
		RETURNING version
		`

	args := []any{
		moduleInfo.ModuleName,
		moduleInfo.ModuleDuration,
		pq.Array(moduleInfo.ExamType),
		moduleInfo.ID,
	}

	return m.DB.QueryRow(query, args...).Scan(&moduleInfo.Version)
}

func (m ModuleInfoModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM module_info
		WHERE id = $1
		`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func ValidateModuleInfo(moduleValidator *validator.Validator, moduleInfo *ModuleInfo) {
	moduleValidator.Check(moduleInfo.ModuleName != "", "module_name", "must be provided")
	moduleValidator.Check(len(moduleInfo.ModuleName) <= 500, "module_name", "must not be more than 500 bytes long")

	moduleValidator.Check(moduleInfo.ModuleDuration != 0, "module_duration", "must be provided")
	moduleValidator.Check(moduleInfo.ModuleDuration >= 10, "module_duration", "must be greater than 10")

	//moduleValidator.Check(moduleInfo.Runtime != 0, "runtime", "must be provided")
	//moduleValidator.Check(moduleInfo.Runtime > 0, "runtime", "must be a positive integer")

	moduleValidator.Check(moduleInfo.ExamType != nil, "exam_type", "must be provided")
	moduleValidator.Check(len(moduleInfo.ExamType) >= 1, "exam_type", "must contain at least 1 type")
	moduleValidator.Check(len(moduleInfo.ExamType) <= 5, "exam_type", "must not contain more than 5 types")
}

type MockModuleInfoModel struct{}

func (m MockModuleInfoModel) Insert(moduleInfo *ModuleInfo) error {
	return nil
}

func (m MockModuleInfoModel) Get(id int64) (*ModuleInfo, error) {
	return &ModuleInfo{}, nil
}

func (m MockModuleInfoModel) Update(moduleInfo *ModuleInfo) error {
	return nil
}

func (m MockModuleInfoModel) Delete(id int64) error {
	return nil
}
