class JobRequisition < ApplicationRecord
  enum Status: [:Draft, :Open, :OnHold, :Closed, :Cancelled]
  enum Priority: [:Low, :Medium, :High, :Critical]


  has_many :Department, class_name: 'Department'
  has_many :HiringManager, class_name: 'Employee'
  has_many :Recruiter, class_name: 'Employee'
  has_many :JobProfile, class_name: 'JobProfile'
  has_many :Candidates, class_name: 'Candidate'
  has_many :Interviews, class_name: 'Interview'
  has_many :Offers, class_name: 'Offer'

end
