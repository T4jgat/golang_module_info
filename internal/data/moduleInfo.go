package data

import (
	"encoding/json"
	"fmt"
	"time"
)

type ModuleInfo struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"-"`
	ModuleName     string    `json:"module_name"`
	ModuleDuration int32     `json:"module_duration"`
	Runtime        int32     `json:"runtime"`
	ExamType       []string  `json:"exam_type"`
	Version        int32     `json:"version"`
}

func (m ModuleInfo) MarshalJSON() ([]byte, error) {
	var runtime string

	if m.Runtime != 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtime)
	}

	type ModuleInfoAlias ModuleInfo

	aux := struct {
		ModuleInfoAlias
		Runtime string `json:"runtime,omitempty"`
	}{
		ModuleInfoAlias: ModuleInfoAlias(m),
		Runtime:         runtime,
	}

	return json.Marshal(aux)
}
