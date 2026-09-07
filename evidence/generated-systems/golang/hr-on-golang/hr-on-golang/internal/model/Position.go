package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Position Declaration
//==============================================================
type Position struct {
    gorm.Model
     PositionCode                                    string
    Fte                                                            string
    DepartmentId         *uint
    Department           *Department `gorm:"foreignKey:DepartmentId"`
    JobProfileId         *uint
    JobProfile           *JobProfile `gorm:"foreignKey:JobProfileId"`
    CostCenterId         *uint
    CostCenter           *CostCenter `gorm:"foreignKey:CostCenterId"`
    LocationId         *uint
    Location           *Location `gorm:"foreignKey:LocationId"`
    ManagerPositionId         *uint
    ManagerPosition           *Position `gorm:"foreignKey:ManagerPositionId"`
     DirectReports           []Position `gorm:"foreignKey:DirectReportsFromPositionId"`
     Assignments           []EmploymentAssignment `gorm:"foreignKey:AssignmentsFromPositionId"`
    Status                      PositionStatus
    WorkLocationType                      WorkLocationType

// parent associations as their child

}

