package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Distributor Declaration
//==============================================================
type Distributor struct {
    gorm.Model
     Name                                    string
    LicenseNumber                                    string
    Region                                    string
     Insurers           []Insurer `gorm:"foreignKey:InsurersFromDistributorId"`
     Agents           []Agent `gorm:"foreignKey:AgentsFromDistributorId"`
     Policies           []Policy `gorm:"foreignKey:PoliciesFromDistributorId"`
    DistributorType                      DistributionChannelType

// parent associations as their child

}

