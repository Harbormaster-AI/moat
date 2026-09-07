package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Report Declaration
//==============================================================
type Report struct {
    gorm.Model
     ReportName                                    string
    GeneratedAt                                                            time.Time
    FileUrl                                                            string
    AdAccountId         *uint
    AdAccount           *AdAccount `gorm:"foreignKey:AdAccountId"`
    CampaignId         *uint
    Campaign           *Campaign `gorm:"foreignKey:CampaignId"`
    LineItemId         *uint
    LineItem           *LineItem `gorm:"foreignKey:LineItemId"`
    ReportType                      ReportType

// parent associations as their child

}

