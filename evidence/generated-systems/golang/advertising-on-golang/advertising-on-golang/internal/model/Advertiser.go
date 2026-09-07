package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Advertiser Declaration
//==============================================================
type Advertiser struct {
    gorm.Model
     Name                                    string
    LegalName                                    string
    Industry                                    string
    Website                                    string
    AgencyId         *uint
    Agency           *Agency `gorm:"foreignKey:AgencyId"`
     AdAccounts           []AdAccount `gorm:"foreignKey:AdAccountsFromAdvertiserId"`
     BillingProfiles           []BillingProfile `gorm:"foreignKey:BillingProfilesFromAdvertiserId"`
     Campaigns           []Campaign `gorm:"foreignKey:CampaignsFromAdvertiserId"`
     TrackingPixels           []TrackingPixel `gorm:"foreignKey:TrackingPixelsFromAdvertiserId"`

// parent associations as their child

}

