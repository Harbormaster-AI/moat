package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Quote Declaration
//==============================================================
type Quote struct {
    gorm.Model
     QuoteNumber                                    string
    TotalAmount                                                            string
    AircraftOrderId         *uint
    AircraftOrder           *AircraftOrder `gorm:"foreignKey:AircraftOrderId"`

// parent associations as their child

}

