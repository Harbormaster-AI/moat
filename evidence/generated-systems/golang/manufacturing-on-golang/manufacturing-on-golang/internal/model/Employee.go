package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Employee Declaration
//==============================================================
type Employee struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    WorkCenterId         *uint
    WorkCenter           *WorkCenter `gorm:"foreignKey:WorkCenterId"`
     ShiftAssignments           []ShiftAssignment `gorm:"foreignKey:ShiftAssignmentsFromEmployeeId"`
     CorrectiveActions           []CorrectiveAction `gorm:"foreignKey:CorrectiveActionsFromEmployeeId"`
    Role                      EmployeeRole
    SkillLevel                      SkillLevel

// parent associations as their child

}

