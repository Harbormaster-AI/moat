class Document < ApplicationRecord
  enum DocumentType: [:Resume, :CoverLetter, :ID, :Certification, :Contract, :Policy, :Other]


  has_many :Candidate, class_name: 'Candidate'
  has_many :Employee, class_name: 'Employee'

end
