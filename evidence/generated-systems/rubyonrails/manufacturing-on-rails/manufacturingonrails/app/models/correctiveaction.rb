class CorrectiveAction < ApplicationRecord
  enum Status: [:Proposed, :Approved, :Implemented, :Verified, :Closed]


  has_many :Nonconformance, class_name: 'Nonconformance'
  has_many :Owner, class_name: 'Employee'

end
