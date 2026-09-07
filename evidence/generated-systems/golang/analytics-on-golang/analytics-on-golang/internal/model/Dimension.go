package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Dimension Declaration
//==============================================================
type Dimension struct {
    gorm.Model
     Name                                    string
    TypeTime                                    bool
    SemanticModelId         *uint
    SemanticModel           *SemanticModel `gorm:"foreignKey:SemanticModelId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromDimensionId"`
     GlossaryTerms           []BusinessGlossaryTerm `gorm:"foreignKey:GlossaryTermsFromDimensionId"`
    DimensionType                      DimensionType

// parent associations as their child

}

