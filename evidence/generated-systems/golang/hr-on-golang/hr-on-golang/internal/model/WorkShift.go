package model

import (
    "gorm.io/gorm"
)

//==============================================================
// WorkShift Declaration
//==============================================================
type WorkShift struct {
    gorm.Model
     StartTime                                                            string
    EndTime                                                            string
    BreakMinutes                                                            string
    WorkScheduleId         *uint
    WorkSchedule           *WorkSchedule `gorm:"foreignKey:WorkScheduleId"`
    DayOfWeek                      DayOfWeek

// parent associations as their child

}

