package model

import (
    "gorm.io/gorm"
)

//==============================================================
// GoodsReceiptLine Declaration
//==============================================================
type GoodsReceiptLine struct {
    gorm.Model
     LineNumber                                                            string
    ReceivedQuantity                                                            string
    AcceptedQuantity                                                            string
    RejectedQuantity                                                            string
    Lot                                                            string
    GoodsReceiptId         *uint
    GoodsReceipt           *GoodsReceipt `gorm:"foreignKey:GoodsReceiptId"`
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`
    InventoryTransactionId         *uint
    InventoryTransaction           *InventoryTransaction `gorm:"foreignKey:InventoryTransactionId"`

// parent associations as their child

}

