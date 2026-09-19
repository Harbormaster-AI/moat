class DispositionReview < ApplicationRecord
  enum Outcome: [:Approved, :Deferred, :Rejected, :Executed]


  has_many :Record, class_name: 'Record_'
  has_many :RetentionSchedule, class_name: 'RetentionSchedule'

end
