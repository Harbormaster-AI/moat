package model

import (
    "gorm.io/gorm"
)

//==============================================================
// TrackingPixel Declaration
//==============================================================
type TrackingPixel struct {
    gorm.Model
     Name                                    string
    Url                                                            string
    CampaignId         *uint
    Campaign           *Campaign `gorm:"foreignKey:CampaignId"`
    AdvertiserId         *uint
    Advertiser           *Advertiser `gorm:"foreignKey:AdvertiserId"`
     ConversionEvents           []ConversionEvent `gorm:"foreignKey:ConversionEventsFromTrackingPixelId"`
    EventType                      ConversionEventType
    PixelType                      PixelType

// parent associations as their child

}

