package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InsertionOrder Declaration
//==============================================================
type InsertionOrder struct {
    gorm.Model
     IoNumber                                    string
    AgreedBudget                                                            string
    Flight                                                            string
    AdvertiserId         *uint
    Advertiser           *Advertiser `gorm:"foreignKey:AdvertiserId"`
    AgencyId         *uint
    Agency           *Agency `gorm:"foreignKey:AgencyId"`
    PublisherId         *uint
    Publisher           *Publisher `gorm:"foreignKey:PublisherId"`
     Campaigns           []Campaign `gorm:"foreignKey:CampaignsFromInsertionOrderId"`
    Status                      IOStatus

// parent associations as their child

}

