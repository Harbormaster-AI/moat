package model

import (
    "gorm.io/gorm"
)

//==============================================================
// StockKeepingUnit Declaration
//==============================================================
type StockKeepingUnit struct {
    gorm.Model
     SkuCode                                                            string
    Name                                    string
    Weight                                                            string
    WeightUnit                                    string
    Volume                                                            string
    VolumeUnit                                    string
    ShelfLifeDays                                                            string
    HazardousMaterial                                    bool
     InventoryItems           []InventoryItem `gorm:"foreignKey:InventoryItemsFromStockKeepingUnitId"`
     UomConversions           []UoMConversion `gorm:"foreignKey:UomConversionsFromStockKeepingUnitId"`
     ReplenishmentPolicies           []ReplenishmentPolicy `gorm:"foreignKey:ReplenishmentPoliciesFromStockKeepingUnitId"`
     Lots           []Lot `gorm:"foreignKey:LotsFromStockKeepingUnitId"`
     SerialNumbers           []SerialNumber `gorm:"foreignKey:SerialNumbersFromStockKeepingUnitId"`
    ItemType                      ItemType
    UnitOfMeasure                      UnitOfMeasure

// parent associations as their child

}

