package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// RoleAssignment Declaration
//==============================================================
type RoleAssignment struct {
    gorm.Model
     EffectiveFrom                                                            time.Time
    EffectiveTo                                                            time.Time
    PersonId         *uint
    Person           *Person `gorm:"foreignKey:PersonId"`
    RoleId         *uint
    Role           *Role `gorm:"foreignKey:RoleId"`
    GovernanceBodyId         *uint
    GovernanceBody           *GovernanceBody `gorm:"foreignKey:GovernanceBodyId"`
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`

// parent associations as their child

}

