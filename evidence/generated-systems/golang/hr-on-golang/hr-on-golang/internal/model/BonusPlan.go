package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BonusPlan Declaration
//==============================================================
type BonusPlan struct {
    gorm.Model
     Name                                    string
    TargetPercentage                                                            string
     CompensationPackages           []CompensationPackage `gorm:"foreignKey:CompensationPackagesFromBonusPlanId"`

// parent associations as their child

}

