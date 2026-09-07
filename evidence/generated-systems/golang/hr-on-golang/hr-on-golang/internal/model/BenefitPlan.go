package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BenefitPlan Declaration
//==============================================================
type BenefitPlan struct {
    gorm.Model
     Name                                    string
    ProviderName                                    string
    EmployeeContributionRate                                                            string
    EmployerContributionRate                                                            string
    EligibilityRules                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Enrollments           []BenefitEnrollment `gorm:"foreignKey:EnrollmentsFromBenefitPlanId"`
    BenefitType                      BenefitType

// parent associations as their child

}

