package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Role Declaration
//==============================================================
type Role struct {
    gorm.Model
     Name                                    string
    Responsibility                                    string
     Assignments           []RoleAssignment `gorm:"foreignKey:AssignmentsFromRoleId"`

// parent associations as their child

}

