package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PriceBookEntry Declaration
//==============================================================
type PriceBookEntry struct {
    gorm.Model
     UnitPrice                                                            string
    EffectiveDate                                                            time.Time
    ExpirationDate                                                            time.Time
    AsActive                                    bool
    PriceBookId         *uint
    PriceBook           *PriceBook `gorm:"foreignKey:PriceBookId"`
    ProductId         *uint
    Product           *Product `gorm:"foreignKey:ProductId"`

// parent associations as their child

}

