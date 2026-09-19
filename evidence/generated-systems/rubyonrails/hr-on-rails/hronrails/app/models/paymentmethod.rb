class PaymentMethod < ApplicationRecord
  enum MethodType: [:DirectDeposit, :Check, :Cash, :InternationalTransfer]


  has_many :Employee, class_name: 'Employee'
  has_many :BankAccount, class_name: 'BankAccount'

end
