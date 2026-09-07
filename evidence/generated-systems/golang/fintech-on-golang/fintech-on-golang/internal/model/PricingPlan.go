package model

import (
    "gorm.io/gorm"
)

//==============================================================
// PricingPlan Declaration
//==============================================================
type PricingPlan struct {
    gorm.Model
     Name                                    string
    PlanCode                                    string
    BaseCurrency                                    string
    ProductOfferingId         *uint
    ProductOffering           *ProductOffering `gorm:"foreignKey:ProductOfferingId"`
     FeeSchedules           []FeeSchedule `gorm:"foreignKey:FeeSchedulesFromPricingPlanId"`
     Limits           []UsageLimit `gorm:"foreignKey:LimitsFromPricingPlanId"`
    Status                      PlanStatus

// parent associations as their child

}

