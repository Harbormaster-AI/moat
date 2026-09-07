package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Campaign Declaration
//==============================================================
type Campaign struct {
    gorm.Model
     Name                                    string
    TotalBudget                                                            string
    Flight                                                            string
    AdAccountId         *uint
    AdAccount           *AdAccount `gorm:"foreignKey:AdAccountId"`
     LineItems           []LineItem `gorm:"foreignKey:LineItemsFromCampaignId"`
     Kpis           []KPI `gorm:"foreignKey:KpisFromCampaignId"`
     TrackingPixels           []TrackingPixel `gorm:"foreignKey:TrackingPixelsFromCampaignId"`
     Audiences           []AudienceSegment `gorm:"foreignKey:AudiencesFromCampaignId"`
     Reports           []Report `gorm:"foreignKey:ReportsFromCampaignId"`
    InsertionOrderId         *uint
    InsertionOrder           *InsertionOrder `gorm:"foreignKey:InsertionOrderId"`
    Objective                      ObjectiveType
    Status                      CampaignStatus

// parent associations as their child

}

