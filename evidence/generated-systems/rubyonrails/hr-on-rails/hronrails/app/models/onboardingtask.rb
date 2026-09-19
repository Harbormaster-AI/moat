class OnboardingTask < ApplicationRecord
  enum Status: [:NotStarted, :InProgress, :Blocked, :Completed]


  has_many :Employee, class_name: 'Employee'
  has_many :AssignedTo, class_name: 'Employee'
  has_many :Dependencies, class_name: 'OnboardingTask'
  has_many :RelatedOffer, class_name: 'Offer'

end
