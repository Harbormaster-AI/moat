package model

import (
    "gorm.io/gorm"
)

//==============================================================
// WorkSchedule Declaration
//==============================================================
type WorkSchedule struct {
    gorm.Model
     Name                                    string
    StandardHoursPerWeek                                                            string
     Contracts           []EmploymentContract `gorm:"foreignKey:ContractsFromWorkScheduleId"`
     Shifts           []WorkShift `gorm:"foreignKey:ShiftsFromWorkScheduleId"`
     Exceptions           []ScheduleException `gorm:"foreignKey:ExceptionsFromWorkScheduleId"`
    ScheduleType                      ScheduleType

// parent associations as their child

}

