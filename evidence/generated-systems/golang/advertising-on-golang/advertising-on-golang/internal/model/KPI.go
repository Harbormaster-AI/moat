package model

import (
    "gorm.io/gorm"
)

//==============================================================
// KPI Declaration
//==============================================================
type KPI struct {
    gorm.Model
     TargetValue                                                            string
    CampaignId         *uint
    Campaign           *Campaign `gorm:"foreignKey:CampaignId"`
    MetricType                      MetricType

// parent associations as their child

}

