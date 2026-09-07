package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Incident Declaration
//==============================================================
type Incident struct {
    gorm.Model
     Location                                                            string
    Description                                    string
    ClaimId         *uint
    Claim           *Claim `gorm:"foreignKey:ClaimId"`
     InsuredObjects           []InsuredObject `gorm:"foreignKey:InsuredObjectsFromIncidentId"`
    IncidentType                      PerilType

// parent associations as their child

}

