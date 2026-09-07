package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Risk Declaration
//==============================================================
type Risk struct {
    gorm.Model
     Name                                    string
    Description                                    string
    InherentRiskScore                                                            string
    ResidualRiskScore                                                            string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Controls           []Control `gorm:"foreignKey:ControlsFromRiskId"`
     Assessments           []RiskAssessment `gorm:"foreignKey:AssessmentsFromRiskId"`
     Issues           []Issue `gorm:"foreignKey:IssuesFromRiskId"`
     Findings           []AuditFinding `gorm:"foreignKey:FindingsFromRiskId"`
    Category                      RiskCategory
    Impact                      RiskImpact
    Likelihood                      RiskLikelihood
    Status                      RiskStatus

// parent associations as their child

}

