package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Plant Declaration
//==============================================================
type Plant struct {
    gorm.Model
     Name                                    string
    PlantCode                                    string
    Address                                                            string
    ManufacturerId         *uint
    Manufacturer           *AerospaceManufacturer `gorm:"foreignKey:ManufacturerId"`
     ProductionLines           []ProductionLine `gorm:"foreignKey:ProductionLinesFromPlantId"`
     Warehouses           []Warehouse `gorm:"foreignKey:WarehousesFromPlantId"`

// parent associations as their child

}

