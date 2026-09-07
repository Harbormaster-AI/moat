package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ExperimentVariant Declaration
//==============================================================
type ExperimentVariant struct {
    gorm.Model
     Name                                    string
    Allocation                                                            string
    ExperimentId         *uint
    Experiment           *Experiment `gorm:"foreignKey:ExperimentId"`
    CreativeVariationId         *uint
    CreativeVariation           *CreativeVariation `gorm:"foreignKey:CreativeVariationId"`
    LineItemId         *uint
    LineItem           *LineItem `gorm:"foreignKey:LineItemId"`

// parent associations as their child

}

