package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ConversionEvent Declaration
//==============================================================
type ConversionEvent struct {
    gorm.Model
     Timestamp                                                            time.Time
    Value                                                            string
    CampaignId         *uint
    Campaign           *Campaign `gorm:"foreignKey:CampaignId"`
    LineItemId         *uint
    LineItem           *LineItem `gorm:"foreignKey:LineItemId"`
    TrackingPixelId         *uint
    TrackingPixel           *TrackingPixel `gorm:"foreignKey:TrackingPixelId"`
    EventType                      ConversionEventType
    AttributionModel                      AttributionModel

// parent associations as their child

}

