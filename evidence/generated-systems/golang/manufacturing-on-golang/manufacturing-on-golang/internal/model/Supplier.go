package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Supplier Declaration
//==============================================================
type Supplier struct {
    gorm.Model
     Name                                    string
    SupplierCode                                    string
    Address                                                            string
     Enterprises           []Enterprise `gorm:"foreignKey:EnterprisesFromSupplierId"`
     Items           []Item `gorm:"foreignKey:ItemsFromSupplierId"`
     PurchaseOrders           []PurchaseOrder `gorm:"foreignKey:PurchaseOrdersFromSupplierId"`
    SupplierTier                      SupplierTier
    PaymentTerms                      PaymentTerms

// parent associations as their child

}

