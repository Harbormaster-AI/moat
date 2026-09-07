package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Experiment Declaration
//==============================================================
type Experiment struct {
    gorm.Model
     Name                                    string
    Hypothesis                                    string
    StartDate                                                            time.Time
    EndDate                                                            time.Time
    CampaignId         *uint
    Campaign           *Campaign `gorm:"foreignKey:CampaignId"`
     Variants           []ExperimentVariant `gorm:"foreignKey:VariantsFromExperimentId"`
    Status                      ExperimentStatus

// parent associations as their child

}

