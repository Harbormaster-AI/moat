package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// CreativeApproval Declaration
//==============================================================
type CreativeApproval struct {
    gorm.Model
     Reviewer                                    string
    ReviewedAt                                                            time.Time
    CreativeAssetId         *uint
    CreativeAsset           *CreativeAsset `gorm:"foreignKey:CreativeAssetId"`
    PublisherId         *uint
    Publisher           *Publisher `gorm:"foreignKey:PublisherId"`
    Status                      CreativeApprovalStatus

// parent associations as their child

}

