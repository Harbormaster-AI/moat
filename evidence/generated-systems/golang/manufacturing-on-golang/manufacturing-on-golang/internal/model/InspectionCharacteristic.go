package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InspectionCharacteristic Declaration
//==============================================================
type InspectionCharacteristic struct {
    gorm.Model
     CharacteristicCode                                    string
    Name                                    string
    LowerSpecLimit                                                            string
    UpperSpecLimit                                                            string
    Target                                                            string
    InspectionPlanId         *uint
    InspectionPlan           *InspectionPlan `gorm:"foreignKey:InspectionPlanId"`
    MeasurementType                      MeasurementType

// parent associations as their child

}

