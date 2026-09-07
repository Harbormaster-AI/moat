package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Publisher Declaration
//==============================================================
type Publisher struct {
    gorm.Model
     Name                                    string
    Website                                    string
     InventorySources           []InventorySource `gorm:"foreignKey:InventorySourcesFromPublisherId"`
     Deals           []Deal `gorm:"foreignKey:DealsFromPublisherId"`
     CreativeApprovals           []CreativeApproval `gorm:"foreignKey:CreativeApprovalsFromPublisherId"`
     InsertionOrders           []InsertionOrder `gorm:"foreignKey:InsertionOrdersFromPublisherId"`
     RateCards           []RateCard `gorm:"foreignKey:RateCardsFromPublisherId"`
    PublisherType                      PublisherType

// parent associations as their child

}

