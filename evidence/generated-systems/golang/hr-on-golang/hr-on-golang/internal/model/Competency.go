package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Competency Declaration
//==============================================================
type Competency struct {
    gorm.Model
     Name                                    string
    Category                                    string
     JobProfiles           []JobProfile `gorm:"foreignKey:JobProfilesFromCompetencyId"`
     CompetencyRatings           []CompetencyRating `gorm:"foreignKey:CompetencyRatingsFromCompetencyId"`

// parent associations as their child

}

