package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CostCenter Declaration
//==============================================================
type CostCenter struct {
    gorm.Model
     Code                                    string
    Name                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Departments           []Department `gorm:"foreignKey:DepartmentsFromCostCenterId"`
     Positions           []Position `gorm:"foreignKey:PositionsFromCostCenterId"`
     Employees           []Employee `gorm:"foreignKey:EmployeesFromCostCenterId"`

// parent associations as their child

}

