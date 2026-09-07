package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Item Declaration
//==============================================================
type Item struct {
    gorm.Model
     ItemNumber                                    string
    Name                                    string
    StandardCost                                                            string
    Weight                                                            string
    AsSerialControlled                                    bool
    BusinessUnitId         *uint
    BusinessUnit           *BusinessUnit `gorm:"foreignKey:BusinessUnitId"`
     Boms           []BOM `gorm:"foreignKey:BomsFromItemId"`
     Routings           []Routing `gorm:"foreignKey:RoutingsFromItemId"`
     Suppliers           []Supplier `gorm:"foreignKey:SuppliersFromItemId"`
     QualitySpecifications           []QualitySpecification `gorm:"foreignKey:QualitySpecificationsFromItemId"`
     InventoryItems           []InventoryItem `gorm:"foreignKey:InventoryItemsFromItemId"`
    ItemType                      ItemType
    ProcurementType                      ProcurementType
    UnitOfMeasure                      UnitOfMeasure
    LifecycleStatus                      ProductLifecycleStatus

// parent associations as their child

}

