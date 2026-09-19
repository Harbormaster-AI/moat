class DataCategory < ApplicationRecord
  enum Classification: [:Public, :Internal, :Confidential, :Restricted, :HighlyRestricted]


  has_many :ProcessingActivities, class_name: 'DataProcessingActivity'
  has_many :Records, class_name: 'Record_'
  has_many :DataBreaches, class_name: 'DataBreach'

end
