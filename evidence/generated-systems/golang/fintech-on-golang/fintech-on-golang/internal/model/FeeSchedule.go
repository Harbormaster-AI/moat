package model

import (
    "gorm.io/gorm"
)

//==============================================================
// FeeSchedule Declaration
//==============================================================
type FeeSchedule struct {
    gorm.Model
     Name                                    string
    Amount                                                            string
    Percentage                                                            string
    Minimum                                                            string
    Maximum                                                            string
    PricingPlanId         *uint
    PricingPlan           *PricingPlan `gorm:"foreignKey:PricingPlanId"`
    FeeType                      FeeType
    CalculationMethod                      FeeCalculationMethod

// parent associations as their child

}

