package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ForecastLine Declaration
//==============================================================
type ForecastLine struct {
    gorm.Model
     Period                                                            time.Time
    Quantity                                                            string
    Confidence                                                            string
    ForecastId         *uint
    Forecast           *Forecast `gorm:"foreignKey:ForecastId"`
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`

// parent associations as their child

}

