package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// InspectionLot Declaration
//==============================================================
type InspectionLot struct {
    gorm.Model
     LotNumber                                    string
    Quantity                                                            string
    SampleSize                                                            string
    CreatedOn                                                            time.Time
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`
    WorkOrderId         *uint
    WorkOrder           *WorkOrder `gorm:"foreignKey:WorkOrderId"`
    GoodsReceiptId         *uint
    GoodsReceipt           *GoodsReceipt `gorm:"foreignKey:GoodsReceiptId"`
     Results           []InspectionResult `gorm:"foreignKey:ResultsFromInspectionLotId"`
    InspectionType                      InspectionType
    Status                      InspectionStatus

// parent associations as their child

}

