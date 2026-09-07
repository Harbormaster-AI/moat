package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AircraftFamily Declaration
//==============================================================
type AircraftFamily struct {
    gorm.Model
     Name                                    string
    FamilyCode                                    string
    ProgramId         *uint
    Program           *AircraftProgram `gorm:"foreignKey:ProgramId"`
     AircraftModels           []AircraftModel `gorm:"foreignKey:AircraftModelsFromAircraftFamilyId"`

// parent associations as their child

}

