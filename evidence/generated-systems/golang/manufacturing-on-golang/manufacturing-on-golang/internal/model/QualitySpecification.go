package model

import (
    "gorm.io/gorm"
)

//==============================================================
// QualitySpecification Declaration
//==============================================================
type QualitySpecification struct {
    gorm.Model
     SpecCode                                    string
    Name                                    string
    Version                                    string
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`

// parent associations as their child

}

