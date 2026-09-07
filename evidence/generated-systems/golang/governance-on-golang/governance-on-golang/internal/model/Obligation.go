package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Obligation Declaration
//==============================================================
type Obligation struct {
    gorm.Model
     ReferenceNumber                                    string
    DescriptionText                                    string
    RegulationId         *uint
    Regulation           *Regulation `gorm:"foreignKey:RegulationId"`
     Controls           []Control `gorm:"foreignKey:ControlsFromObligationId"`
     Policies           []Policy `gorm:"foreignKey:PoliciesFromObligationId"`
     Contracts           []Contract `gorm:"foreignKey:ContractsFromObligationId"`
    ObligationType                      ObligationType
    ReviewFrequency                      ControlFrequency

// parent associations as their child

}

