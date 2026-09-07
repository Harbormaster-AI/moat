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
     Title                                    string
    EffectiveDate                                                            time.Time
    ExpiryDate                                                            time.Time
    RepositoryUrl                                                            string
    ThirdPartyId         *uint
    ThirdParty           *ThirdParty `gorm:"foreignKey:ThirdPartyId"`
     Obligations           []Obligation `gorm:"foreignKey:ObligationsFromContractId"`
     DataProcessingActivities           []DataProcessingActivity `gorm:"foreignKey:DataProcessingActivitiesFromContractId"`
    MatterId         *uint
    Matter           *Matter `gorm:"foreignKey:MatterId"`
    Status                      ContractStatus

// parent associations as their child

}

