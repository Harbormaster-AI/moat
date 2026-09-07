package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Person Declaration
//==============================================================
type Person struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    Email                                                            string
    Department                                    string
     RoleAssignments           []RoleAssignment `gorm:"foreignKey:RoleAssignmentsFromPersonId"`
     OwnedPolicies           []Policy `gorm:"foreignKey:OwnedPoliciesFromPersonId"`
     CorrectiveActions           []CorrectiveAction `gorm:"foreignKey:CorrectiveActionsFromPersonId"`

// parent associations as their child

}

