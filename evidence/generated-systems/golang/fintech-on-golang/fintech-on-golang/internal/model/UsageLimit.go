package model

import (
    "gorm.io/gorm"
)

//==============================================================
// UsageLimit Declaration
//==============================================================
type UsageLimit struct {
    gorm.Model
     Name                                    string
    Amount                                                            string
    Count                                                            string
    PricingPlanId         *uint
    PricingPlan           *PricingPlan `gorm:"foreignKey:PricingPlanId"`
    Scope                      LimitScope
    Period                      LimitPeriod

// parent associations as their child

}

