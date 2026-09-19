class Interview < ApplicationRecord
  enum Stage: [:PhoneScreen, :Technical, :Onsite, :Panel, :HR, :Executive]
  enum Result: [:Pending, :Proceed, :Reject, :OfferRecommended]


  has_many :Requisition, class_name: 'JobRequisition'
  has_many :Candidate, class_name: 'Candidate'
  has_many :Interviewers, class_name: 'Employee'

end
