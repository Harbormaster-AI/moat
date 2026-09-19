class DataSubjectRequest < ApplicationRecord
  enum RequestType: [:Access, :Rectification, :Erasure, :Restriction, :Portability, :Objection, :AutomatedDecisioningReview]
  enum Status: [:Received, :InValidation, :InProgress, :OnHold, :Fulfilled, :Rejected]


  has_many :Organization, class_name: 'Organization'
  has_many :ProcessingActivities, class_name: 'DataProcessingActivity'
  has_many :Records, class_name: 'Record_'

end
