package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CreativeFile Declaration
//==============================================================
type CreativeFile struct {
    gorm.Model
     Uri                                                            string
    FileSizeKB                                                            string
    MimeType                                    string
    Checksum                                    string
    CreativeAssetId         *uint
    CreativeAsset           *CreativeAsset `gorm:"foreignKey:CreativeAssetId"`

// parent associations as their child

}

