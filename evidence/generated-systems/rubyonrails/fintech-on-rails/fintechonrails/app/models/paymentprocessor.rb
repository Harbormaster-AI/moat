class PaymentProcessor < ApplicationRecord


  has_many :Institutions, class_name: 'FinancialInstitution'
  has_many :Contracts, class_name: 'PaymentContract'
  has_many :Settlements, class_name: 'SettlementBatch'

end
