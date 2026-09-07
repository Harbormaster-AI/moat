package model

import (
    "gorm.io/gorm"
)

//==============================================================
// UoMConversion Declaration
//==============================================================
type UoMConversion struct {
    gorm.Model
     Factor                                                            string
    Precision                                                            string
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    FromUnit                      UnitOfMeasure
    ToUnit                      UnitOfMeasure

// parent associations as their child

}

