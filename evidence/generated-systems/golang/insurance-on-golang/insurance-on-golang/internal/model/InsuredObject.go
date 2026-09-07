package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InsuredObject Declaration
//==============================================================
type InsuredObject struct {
    gorm.Model
     Description                                    string
    SerialOrId                                    string
    PrimaryAddress                                                            string
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
     Coverages           []PolicyCoverage `gorm:"foreignKey:CoveragesFromInsuredObjectId"`
    ObjectType                      InsuredObjectType

// parent associations as their child

}

