package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Underwriter Declaration
//==============================================================
type Underwriter struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    EmployeeId                                    string
    AuthorityLimit                                                            string
     Decisions           []UnderwritingDecision `gorm:"foreignKey:DecisionsFromUnderwriterId"`
    InsurerId         *uint
    Insurer           *Insurer `gorm:"foreignKey:InsurerId"`

// parent associations as their child

}

