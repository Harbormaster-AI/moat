package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Regulation Declaration
//==============================================================
type Regulation struct {
    gorm.Model
     Name                                    string
    Citation                                    string
    Jurisdiction                                    string
    PublicationUrl                                                            string
     Obligations           []Obligation `gorm:"foreignKey:ObligationsFromRegulationId"`
     CompliancePrograms           []ComplianceProgram `gorm:"foreignKey:ComplianceProgramsFromRegulationId"`

// parent associations as their child

}

