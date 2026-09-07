package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AdAccount Declaration
//==============================================================
type AdAccount struct {
    gorm.Model
     Name                                    string
    AccountCode                                    string
    DefaultCurrency                                    string
    DefaultTimezone                                    string
    AdvertiserId         *uint
    Advertiser           *Advertiser `gorm:"foreignKey:AdvertiserId"`
     Users           []User `gorm:"foreignKey:UsersFromAdAccountId"`
     Campaigns           []Campaign `gorm:"foreignKey:CampaignsFromAdAccountId"`
    BillingProfileId         *uint
    BillingProfile           *BillingProfile `gorm:"foreignKey:BillingProfileId"`
    DspId         *uint
    Dsp           *DSP `gorm:"foreignKey:DspId"`
     PerformanceMetrics           []PerformanceMetric `gorm:"foreignKey:PerformanceMetricsFromAdAccountId"`

// parent associations as their child

}

