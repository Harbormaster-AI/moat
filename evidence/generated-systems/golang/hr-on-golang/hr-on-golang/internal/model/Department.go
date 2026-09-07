package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Department Declaration
//==============================================================
type Department struct {
    gorm.Model
     Name                                    string
    Code                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    ManagerId         *uint
    Manager           *Employee `gorm:"foreignKey:ManagerId"`
     Positions           []Position `gorm:"foreignKey:PositionsFromDepartmentId"`
     Employees           []Employee `gorm:"foreignKey:EmployeesFromDepartmentId"`
    CostCenterId         *uint
    CostCenter           *CostCenter `gorm:"foreignKey:CostCenterId"`

// parent associations as their child

}

