package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ScheduleException Declaration
//==============================================================
type ScheduleException struct {
    gorm.Model
     Date                                                            time.Time
    Reason                                    string
    Hours                                                            string
    WorkScheduleId         *uint
    WorkSchedule           *WorkSchedule `gorm:"foreignKey:WorkScheduleId"`
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`

// parent associations as their child

}

