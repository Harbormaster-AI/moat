package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ProductOffering Declaration
//==============================================================
type ProductOffering struct {
    gorm.Model
     Name                                    string
    ProductCode                                    string
    InstitutionId         *uint
    Institution           *FinancialInstitution `gorm:"foreignKey:InstitutionId"`
     PricingPlans           []PricingPlan `gorm:"foreignKey:PricingPlansFromProductOfferingId"`
    Category                      ProductCategory

// parent associations as their child

}

