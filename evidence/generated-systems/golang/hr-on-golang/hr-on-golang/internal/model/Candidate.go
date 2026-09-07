package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Candidate Declaration
//==============================================================
type Candidate struct {
    gorm.Model
     Name                                                            string
    Email                                                            string
    Phone                                                            string
     Applications           []JobApplication `gorm:"foreignKey:ApplicationsFromCandidateId"`
     Interviews           []Interview `gorm:"foreignKey:InterviewsFromCandidateId"`
     Offers           []Offer `gorm:"foreignKey:OffersFromCandidateId"`
     Documents           []Document `gorm:"foreignKey:DocumentsFromCandidateId"`
    Source                      CandidateSource

// parent associations as their child

}

