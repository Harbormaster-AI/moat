class DataBreach < ApplicationRecord
  enum Severity: [:Low, :Medium, :High, :Critical]
  enum Status: [:Identified, :Contained, :Notified, :Resolved, :Closed]


  has_many :Organization, class_name: 'Organization'
  has_many :ProcessingActivities, class_name: 'DataProcessingActivity'
  has_many :DataCategories, class_name: 'DataCategory'
  has_many :ThirdParties, class_name: 'ThirdParty'
  has_many :Matter, class_name: 'Matter'

end
