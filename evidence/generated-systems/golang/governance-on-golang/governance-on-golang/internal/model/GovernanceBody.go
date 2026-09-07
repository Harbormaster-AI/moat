package model

import (
    "gorm.io/gorm"
)

//==============================================================
// GovernanceBody Declaration
//==============================================================
type GovernanceBody struct {
    gorm.Model
     Name                                    string
    CharterUrl                                                            string
    Chair                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     RoleAssignments           []RoleAssignment `gorm:"foreignKey:RoleAssignmentsFromGovernanceBodyId"`
     Policies           []Policy `gorm:"foreignKey:PoliciesFromGovernanceBodyId"`
    BodyType                      GovernanceBodyType

// parent associations as their child

}

