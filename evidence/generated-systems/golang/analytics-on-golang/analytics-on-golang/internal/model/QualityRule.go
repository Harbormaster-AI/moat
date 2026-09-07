package model

import (
    "gorm.io/gorm"
)

//==============================================================
// QualityRule Declaration
//==============================================================
type QualityRule struct {
    gorm.Model
     Name                                    string
    Threshold                                                            string
    TargetField                                    string
    DatasetId         *uint
    Dataset           *DataSet `gorm:"foreignKey:DatasetId"`
     Checks           []QualityCheck `gorm:"foreignKey:ChecksFromQualityRuleId"`
    Dimension                      QualityDimension
    Operator                      ComparisonOperator

// parent associations as their child

}

