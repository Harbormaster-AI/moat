package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Account Declaration
//==============================================================
type Account struct {
    gorm.Model
     AccountNumber                                                            string
    Iban                                                            string
    Bic                                                            string
    OpenedDate                                                            time.Time
    Currency                                    string
    Balance                                                            string
    AvailableBalance                                                            string
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    InstitutionId         *uint
    Institution           *FinancialInstitution `gorm:"foreignKey:InstitutionId"`
     Transactions           []Transaction `gorm:"foreignKey:TransactionsFromAccountId"`
     Cards           []PaymentCard `gorm:"foreignKey:CardsFromAccountId"`
     Statements           []AccountStatement `gorm:"foreignKey:StatementsFromAccountId"`
     Mandates           []DirectDebitMandate `gorm:"foreignKey:MandatesFromAccountId"`
    AccountType                      AccountType
    Status                      AccountStatus

// parent associations as their child

}

