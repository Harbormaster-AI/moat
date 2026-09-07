package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ComplianceProgram Declaration
//==============================================================
type ComplianceProgram struct {
    gorm.Model
     Name                                    string
    Framework                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Requirements           []ComplianceRequirement `gorm:"foreignKey:RequirementsFromComplianceProgramId"`
     Controls           []Control `gorm:"foreignKey:ControlsFromComplianceProgramId"`
     Attestations           []Attestation `gorm:"foreignKey:AttestationsFromComplianceProgramId"`
     Regulations           []Regulation `gorm:"foreignKey:RegulationsFromComplianceProgramId"`
    Status                      ComplianceStatus

// parent associations as their child

}

