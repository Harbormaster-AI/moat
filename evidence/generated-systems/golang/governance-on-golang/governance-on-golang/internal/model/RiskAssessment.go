package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// RiskAssessment Declaration
//==============================================================
type RiskAssessment struct {
    gorm.Model
     AssessmentDate                                                            time.Time
    Assessor                                    string
    Summary                                    string
    RiskId         *uint
    Risk           *Risk `gorm:"foreignKey:RiskId"`
    AssessmentType                      AssessmentType

// parent associations as their child

}

