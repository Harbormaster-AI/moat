class Account < ApplicationRecord
  enum AccountType: [:Checking, :Savings, :Current, :Brokerage, :Settlement, :Escrow]
  enum Status: [:Pending, :Active, :Frozen, :Closed]


  composed_of :accountNumber,
    class_name: "AccountNumber",
    mapping: [
      %w[accountNumber_value value]
    ]

  composed_of :iBAN,
    class_name: "IBAN",
    mapping: [
      %w[iBAN_value value]
    ]

  composed_of :bIC,
    class_name: "BIC",
    mapping: [
      %w[bIC_value value]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Customer, class_name: 'Customer'
  has_many :Institution, class_name: 'FinancialInstitution'
  has_many :Transactions, class_name: 'Transaction'
  has_many :Cards, class_name: 'PaymentCard'
  has_many :Statements, class_name: 'AccountStatement'
  has_many :Mandates, class_name: 'DirectDebitMandate'

end
