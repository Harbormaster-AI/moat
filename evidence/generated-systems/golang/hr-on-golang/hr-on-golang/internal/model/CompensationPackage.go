package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// CompensationPackage Declaration
//==============================================================
type CompensationPackage struct {
    gorm.Model
     EffectiveFrom                                                            time.Time
    EffectiveTo                                                            time.Time
    Currency                                    string
    ContractId         *uint
    Contract           *EmploymentContract `gorm:"foreignKey:ContractId"`
     SalaryComponents           []SalaryComponent `gorm:"foreignKey:SalaryComponentsFromCompensationPackageId"`
     BonusPlans           []BonusPlan `gorm:"foreignKey:BonusPlansFromCompensationPackageId"`
     EquityGrants           []EquityGrant `gorm:"foreignKey:EquityGrantsFromCompensationPackageId"`

// parent associations as their child

}

