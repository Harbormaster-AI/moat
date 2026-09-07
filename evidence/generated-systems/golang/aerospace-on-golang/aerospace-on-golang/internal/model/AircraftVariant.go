package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AircraftVariant Declaration
//==============================================================
type AircraftVariant struct {
    gorm.Model
     VariantCode                                    string
    RangeNm                                                            string
    MaxTakeoffWeightKg                                                            string
    Model_Id         *uint
    Model_           *AircraftModel `gorm:"foreignKey:Model_Id"`
    EngineTypeId         *uint
    EngineType           *EngineType `gorm:"foreignKey:EngineTypeId"`
    AvionicsSuiteId         *uint
    AvionicsSuite           *AvionicsSuite `gorm:"foreignKey:AvionicsSuiteId"`
    ApuId         *uint
    Apu           *APU `gorm:"foreignKey:ApuId"`
    LandingGearId         *uint
    LandingGear           *LandingGear `gorm:"foreignKey:LandingGearId"`
     CabinLayouts           []CabinLayout `gorm:"foreignKey:CabinLayoutsFromAircraftVariantId"`
     Options           []AircraftOption `gorm:"foreignKey:OptionsFromAircraftVariantId"`
     Packages           []AircraftPackage `gorm:"foreignKey:PackagesFromAircraftVariantId"`

// parent associations as their child

}

