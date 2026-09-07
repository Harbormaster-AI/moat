package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ThirdPartyAssessment Declaration
//==============================================================
type ThirdPartyAssessment struct {
    gorm.Model
     AssessmentDate                                                            time.Time
    Assessor                                    string
    ThirdPartyId         *uint
    ThirdParty           *ThirdParty `gorm:"foreignKey:ThirdPartyId"`
     Issues           []Issue `gorm:"foreignKey:IssuesFromThirdPartyAssessmentId"`
    AssessmentType                      AssessmentType
    Result                      AssessmentResult

// parent associations as their child

}

