package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Policy Declaration
//==============================================================
type Policy struct {
    gorm.Model
     Title                                    string
    VersionLabel                                    string
    ApprovalDate                                                            time.Time
    NextReviewDate                                                            time.Time
    DocumentUrl                                                            string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Owners           []Person `gorm:"foreignKey:OwnersFromPolicyId"`
     RelatedRequirements           []ComplianceRequirement `gorm:"foreignKey:RelatedRequirementsFromPolicyId"`
     Controls           []Control `gorm:"foreignKey:ControlsFromPolicyId"`
     Procedures           []Procedure `gorm:"foreignKey:ProceduresFromPolicyId"`
     Exceptions           []Exception_ `gorm:"foreignKey:ExceptionsFromPolicyId"`
     Attestations           []Attestation `gorm:"foreignKey:AttestationsFromPolicyId"`
    PolicyType                      PolicyType
    Status                      DocumentStatus

// parent associations as their child

}

