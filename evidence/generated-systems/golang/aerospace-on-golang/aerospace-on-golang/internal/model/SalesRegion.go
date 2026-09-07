package model

import (
    "gorm.io/gorm"
)

//==============================================================
// SalesRegion Declaration
//==============================================================
type SalesRegion struct {
    gorm.Model
     Name                                    string
    RegionCode                                    string
     Operators           []Operator `gorm:"foreignKey:OperatorsFromSalesRegionId"`
     SalesCampaigns           []SalesCampaign `gorm:"foreignKey:SalesCampaignsFromSalesRegionId"`

// parent associations as their child

}

