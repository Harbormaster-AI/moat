class BackgroundCheck < ApplicationRecord
  enum Status: [:Ordered, :InProgress, :Clear, :Adverse, :Cancelled]


  has_many :Candidate, class_name: 'Candidate'
  has_many :Requisition, class_name: 'JobRequisition'
  has_many :Report, class_name: 'Document'

end
