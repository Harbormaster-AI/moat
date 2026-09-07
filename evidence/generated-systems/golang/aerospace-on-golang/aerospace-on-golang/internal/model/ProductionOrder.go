package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ProductionOrder Declaration
//==============================================================
type ProductionOrder struct {
    gorm.Model
     OrderNumber                                    string
    VariantId         *uint
    Variant           *AircraftVariant `gorm:"foreignKey:VariantId"`
    PlantId         *uint
    Plant           *Plant `gorm:"foreignKey:PlantId"`
    AircraftOrderId         *uint
    AircraftOrder           *AircraftOrder `gorm:"foreignKey:AircraftOrderId"`
    Status                      ProductionOrderStatus

// parent associations as their child

}

