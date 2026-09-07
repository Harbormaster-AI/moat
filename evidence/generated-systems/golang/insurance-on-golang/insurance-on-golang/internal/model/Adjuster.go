package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Adjuster Declaration
//==============================================================
type Adjuster struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    LicenseNumber                                    string
     Claims           []Claim `gorm:"foreignKey:ClaimsFromAdjusterId"`
     ServiceProviders           []ServiceProvider `gorm:"foreignKey:ServiceProvidersFromAdjusterId"`
    AdjusterType                      AdjusterType

// parent associations as their child

}

