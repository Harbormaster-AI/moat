class ReturnItem < ApplicationRecord
  enum Reason: [:Defective, :Damaged, :NotAsDescribed, :WrongItem, :NoLongerNeeded, :SizeFitIssue, :Other]
  enum Condition: [:New, :OpenBox, :Used, :Damaged, :MissingParts]


  has_many :ReturnRequest, class_name: 'ReturnRequest'
  has_many :OrderLine, class_name: 'OrderLine'

end
