package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Location Declaration
//==============================================================
type Location struct {
    gorm.Model
     Name                                    string
    Address                                                            string
    Timezone                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Departments           []Department `gorm:"foreignKey:DepartmentsFromLocationId"`
     Positions           []Position `gorm:"foreignKey:PositionsFromLocationId"`
     Employees           []Employee `gorm:"foreignKey:EmployeesFromLocationId"`

// parent associations as their child

}

