package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PurchaseOrder Declaration
//==============================================================
type PurchaseOrder struct {
    gorm.Model
     PoNumber                                    string
    OrderDate                                                            time.Time
    TotalAmount                                                            string
    SupplierId         *uint
    Supplier           *Supplier `gorm:"foreignKey:SupplierId"`
    PlantId         *uint
    Plant           *Plant `gorm:"foreignKey:PlantId"`
     Lines           []PurchaseOrderLine `gorm:"foreignKey:LinesFromPurchaseOrderId"`
     GoodsReceipts           []GoodsReceipt `gorm:"foreignKey:GoodsReceiptsFromPurchaseOrderId"`
    Status                      PurchaseOrderStatus

// parent associations as their child

}

