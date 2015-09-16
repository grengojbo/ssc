package card

import (
	"database/sql"
	"time"
)

type Departments struct {
	DeptID     int64          `json:"id"`
	DeptName   string         `json:"-"`
	Code       sql.NullString `json:"code"`
	Invalidate time.Time      `json:"invalidate"`
	Status     int            `json:"status"`
	Supdeptid  int64          `json:"supdeptid"`
	Type       string         `json:"type"`
}

func (this *Departments) TableName() string {
	return "departments"
}
