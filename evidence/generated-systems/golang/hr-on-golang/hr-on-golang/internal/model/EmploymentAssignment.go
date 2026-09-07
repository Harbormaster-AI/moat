package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// EmploymentAssignment Declaration
//==============================================================
type EmploymentAssignment struct {
    gorm.Model
     StartDate                                                            time.Time
    EndDate                                                            time.Time
    Primary                                    bool
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    PositionId         *uint
    Position           *Position `gorm:"foreignKey:PositionId"`
    SupervisorId         *uint
    Supervisor           *Employee `gorm:"foreignKey:SupervisorId"`
    AssignmentType                      AssignmentType
    Status                      AssignmentStatus

// parent associations as their child

}

