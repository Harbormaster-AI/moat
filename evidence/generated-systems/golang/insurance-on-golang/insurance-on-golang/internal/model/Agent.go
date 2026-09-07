package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Agent Declaration
//==============================================================
type Agent struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    LicenseId                                    string
    DistributorId         *uint
    Distributor           *Distributor `gorm:"foreignKey:DistributorId"`
     Policies           []Policy `gorm:"foreignKey:PoliciesFromAgentId"`
     Customers           []Customer `gorm:"foreignKey:CustomersFromAgentId"`
    Status                      ProducerStatus

// parent associations as their child

}

