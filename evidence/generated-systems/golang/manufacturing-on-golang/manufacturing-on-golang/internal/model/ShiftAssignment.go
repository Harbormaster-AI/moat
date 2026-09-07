package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ShiftAssignment Declaration
//==============================================================
type ShiftAssignment struct {
    gorm.Model
     AssignmentDate                                                            time.Time
    ShiftId         *uint
    Shift           *Shift `gorm:"foreignKey:ShiftId"`
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    WorkCenterId         *uint
    WorkCenter           *WorkCenter `gorm:"foreignKey:WorkCenterId"`

// parent associations as their child

}

