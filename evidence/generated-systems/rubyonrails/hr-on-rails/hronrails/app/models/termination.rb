class Termination < ApplicationRecord
  enum Reason: [:Voluntary, :Involuntary, :Retirement, :Redundancy, :EndOfContract]
  enum Type: [:Resignation, :Dismissal, :Layoff, :Retirement, :EndOfAssignment]


  has_many :Employee, class_name: 'Employee'
  has_many :Assignment, class_name: 'EmploymentAssignment'

end
