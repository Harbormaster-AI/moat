package model

import (
    "gorm.io/gorm"
)

//==============================================================
// MedicalSupplier Declaration
//==============================================================
type MedicalSupplier struct {
    gorm.Model
     Name                                    string
    Website                                    string
     Facilities           []Facility `gorm:"foreignKey:FacilitiesFromMedicalSupplierId"`
     InventoryItems           []InventoryItem `gorm:"foreignKey:InventoryItemsFromMedicalSupplierId"`
    SupplierTier                      SupplierTier

// parent associations as their child

}

