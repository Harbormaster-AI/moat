class Goal < ApplicationRecord
  enum Status: [:NotStarted, :InProgress, :Completed, :Deferred, :Cancelled]


  composed_of :percentage,
    class_name: "Percentage",
    mapping: [
      %w[percentage_value value]
    ]

  has_many :Employee, class_name: 'Employee'
  has_many :Cycle, class_name: 'PerformanceCycle'
  has_many :ParentGoal, class_name: 'Goal'
  has_many :ChildGoals, class_name: 'Goal'

end
