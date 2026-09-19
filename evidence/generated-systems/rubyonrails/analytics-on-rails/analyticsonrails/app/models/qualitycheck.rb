class QualityCheck < ApplicationRecord
  enum Status: [:Passed, :Failed, :Warning, :Skipped]


  has_many :Rule, class_name: 'QualityRule'
  has_many :Dataset, class_name: 'DataSet'

end
