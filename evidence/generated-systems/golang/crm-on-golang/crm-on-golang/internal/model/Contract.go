package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Contract Declaration
//==============================================================
type Contract struct {
    gorm.Model
     ContractNumber                                    string
    StartDate                                                            time.Time
    EndDate                                                            time.Time
    RenewalTermMonths                                                            string
    AutoRenew                                    bool
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    OwnerId         *uint
    Owner           *User `gorm:"foreignKey:OwnerId"`
     Orders           []Order `gorm:"foreignKey:OrdersFromContractId"`
     Cases           []Case_ `gorm:"foreignKey:CasesFromContractId"`
    Status                      ContractStatus

// parent associations as their child

}

