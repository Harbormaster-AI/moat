package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InspectionPlan Declaration
//==============================================================
type InspectionPlan struct {
    gorm.Model
     PlanNumber                                    string
    Revision                                    string
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`
     Characteristics           []InspectionCharacteristic `gorm:"foreignKey:CharacteristicsFromInspectionPlanId"`
    SamplingPlan                      SamplingPlanType
    Status                      QualityPlanStatus

// parent associations as their child

}

