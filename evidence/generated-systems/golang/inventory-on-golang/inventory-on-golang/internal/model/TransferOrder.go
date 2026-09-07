package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// TransferOrder Declaration
//==============================================================
type TransferOrder struct {
    gorm.Model
     OrderNumber                                    string
    RequestedShipDate                                                            time.Time
    RequestedReceiveDate                                                            time.Time
    ShippedDate                                                            time.Time
    ReceivedDate                                                            time.Time
    OriginWarehouseId         *uint
    OriginWarehouse           *Warehouse `gorm:"foreignKey:OriginWarehouseId"`
    DestinationWarehouseId         *uint
    DestinationWarehouse           *Warehouse `gorm:"foreignKey:DestinationWarehouseId"`
     Lines           []TransferOrderLine `gorm:"foreignKey:LinesFromTransferOrderId"`
     Transactions           []InventoryTransaction `gorm:"foreignKey:TransactionsFromTransferOrderId"`
    Status                      TransferOrderStatus

// parent associations as their child

}

