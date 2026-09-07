package model

import (
    "gorm.io/gorm"
)

//==============================================================
// SalesCampaign Declaration
//==============================================================
type SalesCampaign struct {
    gorm.Model
     CampaignCode                                    string
    RegionId         *uint
    Region           *SalesRegion `gorm:"foreignKey:RegionId"`
    OperatorId         *uint
    Operator           *Operator `gorm:"foreignKey:OperatorId"`
     Quotes           []Quote `gorm:"foreignKey:QuotesFromSalesCampaignId"`
    Status                      SalesCampaignStatus

// parent associations as their child

}

