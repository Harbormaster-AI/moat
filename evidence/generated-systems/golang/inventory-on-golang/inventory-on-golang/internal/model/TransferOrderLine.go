package model

import (
    "gorm.io/gorm"
)

//==============================================================
// TransferOrderLine Declaration
//==============================================================
type TransferOrderLine struct {
    gorm.Model
     LineNumber                                                            string
    Quantity                                                            string
    TransferOrderId         *uint
    TransferOrder           *TransferOrder `gorm:"foreignKey:TransferOrderId"`
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    LotId         *uint
    Lot           *Lot `gorm:"foreignKey:LotId"`
     SerialNumbers           []SerialNumber `gorm:"foreignKey:SerialNumbersFromTransferOrderLineId"`
    FromLocationId         *uint
    FromLocation           *StorageLocation `gorm:"foreignKey:FromLocationId"`
    ToLocationId         *uint
    ToLocation           *StorageLocation `gorm:"foreignKey:ToLocationId"`
    UnitOfMeasure                      UnitOfMeasure
    StockStatus                      StockStatus

// parent associations as their child

}

