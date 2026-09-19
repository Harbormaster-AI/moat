class Document < ApplicationRecord
  enum DocumentType: [:ApplicationForm, :PolicyDocument, :Endorsement, :Invoice, :ClaimForm, :PoliceReport, :Estimate, :Photo, :MedicalRecord, :Correspondence]


  has_many :Policy, class_name: 'Policy'
  has_many :Claim, class_name: 'Claim'
  has_many :Application, class_name: 'Application'
  has_many :Customer, class_name: 'Customer'

end
