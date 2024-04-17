package data

import (
	"database/sql"
	"time"
)

type TeacherInfo struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated-at"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	email      string    `json:"email"`
	ModuleInfo int64     `json:"module_info_fk" gorm:"foreignKey:ModuleInfoID"`
}

type TeacherInfoModel struct {
	DB *sql.DB
}

func (m TeacherInfoModel) GetAll() ([]*TeacherInfo, error) {
	query := `SELECT id, name, surname, email, module_info_fk FROM teacher_info`
	var teachers []*TeacherInfo

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var teacher TeacherInfo

		err := rows.Scan(
			&teacher.ID,
			&teacher.Name,
			&teacher.Surname,
			&teacher.email,
			&teacher.ModuleInfo,
		)
		if err != nil {
			return nil, err
		}

		teachers = append(teachers, &teacher)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}
