package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Component_ Declaration
//==============================================================
type Component_ struct {
    gorm.Model
     PartNumber                                    string
    Name                                    string
    SupplierId         *uint
    Supplier           *Supplier `gorm:"foreignKey:SupplierId"`
    ComponentCategory                      ComponentCategory
    SerializationMethod                      SerializationMethod

// parent associations as their child

}

