package model

import (
    "gorm.io/gorm"
)

//==============================================================
// PaymentMethod Declaration
//==============================================================
type PaymentMethod struct {
    gorm.Model
     Preferred                                    bool
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    BankAccountId         *uint
    BankAccount           *BankAccount `gorm:"foreignKey:BankAccountId"`
    MethodType                      PaymentMethodType

// parent associations as their child

}

