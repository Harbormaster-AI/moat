package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Policy Declaration
//==============================================================
type Policy struct {
    gorm.Model
     PolicyNumber                                                            string
    EffectivePeriod                                                            string
    TotalPremium                                                            string
    InsurerId         *uint
    Insurer           *Insurer `gorm:"foreignKey:InsurerId"`
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    ProductId         *uint
    Product           *InsuranceProduct `gorm:"foreignKey:ProductId"`
    AgentId         *uint
    Agent           *Agent `gorm:"foreignKey:AgentId"`
     Coverages           []PolicyCoverage `gorm:"foreignKey:CoveragesFromPolicyId"`
     InsuredObjects           []InsuredObject `gorm:"foreignKey:InsuredObjectsFromPolicyId"`
     Endorsements           []Endorsement `gorm:"foreignKey:EndorsementsFromPolicyId"`
    BillingAccountId         *uint
    BillingAccount           *BillingAccount `gorm:"foreignKey:BillingAccountId"`
     Beneficiaries           []Beneficiary `gorm:"foreignKey:BeneficiariesFromPolicyId"`
     Claims           []Claim `gorm:"foreignKey:ClaimsFromPolicyId"`
     ReinsuranceAgreements           []ReinsuranceAgreement `gorm:"foreignKey:ReinsuranceAgreementsFromPolicyId"`
    Status                      PolicyStatus
    PaymentPlan                      PaymentPlanType

// parent associations as their child

}

