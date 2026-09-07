package model

import (
    "gorm.io/gorm"
)

//==============================================================
// EngineType Declaration
//==============================================================
type EngineType struct {
    gorm.Model
     EngineModelCode                                    string
    MaxThrustKn                                                            string
    SupplierId         *uint
    Supplier           *Supplier `gorm:"foreignKey:SupplierId"`
     CompatibleModels           []AircraftModel `gorm:"foreignKey:CompatibleModelsFromEngineTypeId"`
    Category                      EngineCategory

// parent associations as their child

}

