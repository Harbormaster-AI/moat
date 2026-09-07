package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CompetencyRating Declaration
//==============================================================
type CompetencyRating struct {
    gorm.Model
     Comment                                    string
    ReviewId         *uint
    Review           *PerformanceReview `gorm:"foreignKey:ReviewId"`
    CompetencyId         *uint
    Competency           *Competency `gorm:"foreignKey:CompetencyId"`
    Rating                      PerformanceRating

// parent associations as their child

}

