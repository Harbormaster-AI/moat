package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InventoryItem Declaration
//==============================================================
type InventoryItem struct {
    gorm.Model
     Sku                                    string
    Name                                    string
    QuantityOnHand                                                            string
    QuantityReserved                                                            string
    FacilityId         *uint
    Facility           *Facility `gorm:"foreignKey:FacilityId"`
    SupplierId         *uint
    Supplier           *MedicalSupplier `gorm:"foreignKey:SupplierId"`

// parent associations as their child

}

