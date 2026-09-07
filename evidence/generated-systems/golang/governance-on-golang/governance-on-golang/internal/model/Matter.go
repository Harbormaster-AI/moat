package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Matter Declaration
//==============================================================
type Matter struct {
    gorm.Model
     MatterName                                    string
    LeadCounsel                                    string
     LegalHolds           []LegalHold `gorm:"foreignKey:LegalHoldsFromMatterId"`
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     DataBreaches           []DataBreach `gorm:"foreignKey:DataBreachesFromMatterId"`
     Contracts           []Contract `gorm:"foreignKey:ContractsFromMatterId"`
    MatterType                      MatterType
    Status                      MatterStatus

// parent associations as their child

}

