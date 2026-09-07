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
     Score                                                            string
    AssessedAt                                                            time.Time
    ModelVersion                                    string
    Notes                                    string
    ApplicationId         *uint
    Application           *LoanApplication `gorm:"foreignKey:ApplicationId"`
    Decision                      DecisionOutcome

// parent associations as their child

}

