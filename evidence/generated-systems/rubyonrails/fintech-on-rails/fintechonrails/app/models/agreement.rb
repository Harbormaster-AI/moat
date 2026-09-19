class Agreement < ApplicationRecord
  enum AgreementType: [:TermsOfService, :PrivacyPolicy, :LoanAgreement, :AccountAgreement]
  enum Status: [:Active, :Suspended, :Terminated]


  has_many :Customer, class_name: 'Customer'
  has_many :ProductOffering, class_name: 'ProductOffering'

end
