package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Beneficiary Declaration
//==============================================================
type Beneficiary struct {
    gorm.Model
     Name                                    string
    Share                                                            string
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    Relationship                      RelationshipType

// parent associations as their child

}

