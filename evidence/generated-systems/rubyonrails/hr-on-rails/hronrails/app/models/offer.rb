class Offer < ApplicationRecord
  enum Status: [:Draft, :Sent, :Accepted, :Declined, :Withdrawn, :Expired]


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

  has_many :Requisition, class_name: 'JobRequisition'
  has_many :Candidate, class_name: 'Candidate'
  has_many :ApprovedBy, class_name: 'Employee'
  has_many :Contract, class_name: 'EmploymentContract'

end
