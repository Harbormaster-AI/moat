package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Warehouse Declaration
//==============================================================
type Warehouse struct {
    gorm.Model
     Name                                    string
    Code                                    string
    Address                                                            string
    TimeZone                                    string
    AllowsOverAllocation                                    bool
     StorageLocations           []StorageLocation `gorm:"foreignKey:StorageLocationsFromWarehouseId"`
     InventoryItems           []InventoryItem `gorm:"foreignKey:InventoryItemsFromWarehouseId"`
     InboundShipments           []InboundShipment `gorm:"foreignKey:InboundShipmentsFromWarehouseId"`
     OutboundAllocations           []OutboundAllocation `gorm:"foreignKey:OutboundAllocationsFromWarehouseId"`
     OriginTransfers           []TransferOrder `gorm:"foreignKey:OriginTransfersFromWarehouseId"`
     DestinationTransfers           []TransferOrder `gorm:"foreignKey:DestinationTransfersFromWarehouseId"`
     CycleCounts           []CycleCount `gorm:"foreignKey:CycleCountsFromWarehouseId"`

// parent associations as their child

}

