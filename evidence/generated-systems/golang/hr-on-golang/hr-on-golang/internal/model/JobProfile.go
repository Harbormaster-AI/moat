package model

import (
    "gorm.io/gorm"
)

//==============================================================
// JobProfile Declaration
//==============================================================
type JobProfile struct {
    gorm.Model
     Title                                    string
    JobCode                                    string
    JobFamilyId         *uint
    JobFamily           *JobFamily `gorm:"foreignKey:JobFamilyId"`
     Competencies           []Competency `gorm:"foreignKey:CompetenciesFromJobProfileId"`
     TrainingRecommendations           []TrainingCourse `gorm:"foreignKey:TrainingRecommendationsFromJobProfileId"`
     Positions           []Position `gorm:"foreignKey:PositionsFromJobProfileId"`
    JobLevel                      JobLevel
    ExemptStatus                      ExemptStatus

// parent associations as their child

}

