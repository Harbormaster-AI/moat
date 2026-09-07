package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ComplianceRequirement Declaration
//==============================================================
type ComplianceRequirement struct {
    gorm.Model
     Name                                    string
    Source                                    string
    Citation                                    string
    ComplianceProgramId         *uint
    ComplianceProgram           *ComplianceProgram `gorm:"foreignKey:ComplianceProgramId"`
     Policies           []Policy `gorm:"foreignKey:PoliciesFromComplianceRequirementId"`
     Controls           []Control `gorm:"foreignKey:ControlsFromComplianceRequirementId"`
     Obligations           []Obligation `gorm:"foreignKey:ObligationsFromComplianceRequirementId"`
    Applicability                      Applicability
    Status                      ComplianceStatus

// parent associations as their child

}

