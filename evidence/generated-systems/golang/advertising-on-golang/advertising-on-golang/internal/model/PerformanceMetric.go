package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PerformanceMetric Declaration
//==============================================================
type PerformanceMetric struct {
    gorm.Model
     Date                                                            time.Time
    Value                                                            string
    AdAccountId         *uint
    AdAccount           *AdAccount `gorm:"foreignKey:AdAccountId"`
    CampaignId         *uint
    Campaign           *Campaign `gorm:"foreignKey:CampaignId"`
    LineItemId         *uint
    LineItem           *LineItem `gorm:"foreignKey:LineItemId"`
    PlacementId         *uint
    Placement           *Placement `gorm:"foreignKey:PlacementId"`
    CreativeAssetId         *uint
    CreativeAsset           *CreativeAsset `gorm:"foreignKey:CreativeAssetId"`
    MetricType                      MetricType

// parent associations as their child

}

