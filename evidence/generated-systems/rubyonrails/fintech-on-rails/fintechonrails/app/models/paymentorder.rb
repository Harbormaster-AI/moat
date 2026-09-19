class PaymentOrder < ApplicationRecord
  enum PaymentMethod: [:Card, :BankTransfer, :DirectDebit, :Wallet, :Cash]
  enum Status: [:Draft, :Submitted, :Processing, :Completed, :Cancelled, :Failed]
  enum Priority: [:Normal, :Urgent]


  has_many :SourceAccount, class_name: 'Account'
  has_many :DestinationAccount, class_name: 'Account'
  has_many :Beneficiary, class_name: 'Beneficiary'
  has_many :Transactions, class_name: 'Transaction'
  has_many :FxDeal, class_name: 'FXDeal'
  has_many :Fees, class_name: 'AppliedFee'

end
