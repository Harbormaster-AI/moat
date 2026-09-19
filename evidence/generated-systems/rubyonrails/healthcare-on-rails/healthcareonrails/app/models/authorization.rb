class Authorization < ApplicationRecord
  enum Status: [:Requested, :PendingReview, :Approved, :Denied, :Expired]


  has_many :Coverage, class_name: 'Coverage'
  has_many :Order, class_name: 'ClinicalOrder'

end
