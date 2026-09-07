package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AircraftModel Declaration
//==============================================================
type AircraftModel struct {
    gorm.Model
     Name                                    string
    ModelDesignation                                    string
    FamilyId         *uint
    Family           *AircraftFamily `gorm:"foreignKey:FamilyId"`
     Variants           []AircraftVariant `gorm:"foreignKey:VariantsFromAircraftModelId"`
     EngineTypes           []EngineType `gorm:"foreignKey:EngineTypesFromAircraftModelId"`
    AircraftType                      AircraftType

// parent associations as their child

}

