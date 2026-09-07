package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// GoodsReceipt Declaration
//==============================================================
type GoodsReceipt struct {
    gorm.Model
     ReceiptNumber                                    string
    ReceiptDate                                                            time.Time
    PurchaseOrderId         *uint
    PurchaseOrder           *PurchaseOrder `gorm:"foreignKey:PurchaseOrderId"`
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
     Lines           []GoodsReceiptLine `gorm:"foreignKey:LinesFromGoodsReceiptId"`
    Status                      ReceiptStatus

// parent associations as their child

}

