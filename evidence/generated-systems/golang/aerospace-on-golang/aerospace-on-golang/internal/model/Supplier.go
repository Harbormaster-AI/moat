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
     Manufacturers           []AerospaceManufacturer `gorm:"foreignKey:ManufacturersFromSupplierId"`
     Components           []Component_ `gorm:"foreignKey:ComponentsFromSupplierId"`
     EngineTypes           []EngineType `gorm:"foreignKey:EngineTypesFromSupplierId"`
     AvionicsSuites           []AvionicsSuite `gorm:"foreignKey:AvionicsSuitesFromSupplierId"`
     Apus           []APU `gorm:"foreignKey:ApusFromSupplierId"`
     LandingGears           []LandingGear `gorm:"foreignKey:LandingGearsFromSupplierId"`
    SupplierType                      SupplierType
    ApprovalStatus                      SupplierApprovalStatus

// parent associations as their child

}

