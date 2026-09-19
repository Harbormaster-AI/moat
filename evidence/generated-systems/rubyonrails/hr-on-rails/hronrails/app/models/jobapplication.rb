class JobApplication < ApplicationRecord
  enum Status: [:New, :Screening, :Interview, :Offer, :Hired, :Rejected, :Withdrawn]


  has_many :Candidate, class_name: 'Candidate'
  has_many :Requisition, class_name: 'JobRequisition'
  has_many :Screenings, class_name: 'Screening'

end
