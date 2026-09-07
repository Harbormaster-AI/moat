package model

import (
    "gorm.io/gorm"
)

//==============================================================
// LineItem Declaration
//==============================================================
type LineItem struct {
    gorm.Model
     Name                                    string
    BidAmount                                                            string
    DailyBudget                                                            string
    FrequencyCap                                                            string
    CampaignId         *uint
    Campaign           *Campaign `gorm:"foreignKey:CampaignId"`
     Placements           []Placement `gorm:"foreignKey:PlacementsFromLineItemId"`
    TargetingProfileId         *uint
    TargetingProfile           *TargetingProfile `gorm:"foreignKey:TargetingProfileId"`
    DealId         *uint
    Deal           *Deal `gorm:"foreignKey:DealId"`
     Creatives           []CreativeAsset `gorm:"foreignKey:CreativesFromLineItemId"`
     PerformanceMetrics           []PerformanceMetric `gorm:"foreignKey:PerformanceMetricsFromLineItemId"`
    Status                      LineItemStatus
    PricingModel                      PricingModel
    BidStrategy                      BidStrategyType
    Pacing                      PacingType

// parent associations as their child

}

