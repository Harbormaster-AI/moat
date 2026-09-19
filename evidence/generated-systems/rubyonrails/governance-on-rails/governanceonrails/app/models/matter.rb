class Matter < ApplicationRecord
  enum MatterType: [:Litigation, :Investigation, :RegulatoryInquiry, :Complaint, :Arbitration]
  enum Status: [:Open, :ActiveDiscovery, :Negotiation, :Settled, :Closed]


  has_many :LegalHolds, class_name: 'LegalHold'
  has_many :Organization, class_name: 'Organization'
  has_many :DataBreaches, class_name: 'DataBreach'
  has_many :Contracts, class_name: 'Contract'

end
