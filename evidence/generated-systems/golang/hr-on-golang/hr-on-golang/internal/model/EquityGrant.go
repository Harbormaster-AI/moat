package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// EquityGrant Declaration
//==============================================================
type EquityGrant struct {
    gorm.Model
     GrantId                                    string
    GrantedUnits                                                            string
    VestingStart                                                            time.Time
    CompensationPackageId         *uint
    CompensationPackage           *CompensationPackage `gorm:"foreignKey:CompensationPackageId"`
    GrantType                      EquityType

// parent associations as their child

}

