class Employee < ApplicationRecord
  enum Role: [:Operator, :Technician, :Supervisor, :Planner, :QualityEngineer, :Buyer]
  enum SkillLevel: [:Novice, :Competent, :Proficient, :Expert]


  has_many :WorkCenter, class_name: 'WorkCenter'
  has_many :ShiftAssignments, class_name: 'ShiftAssignment'
  has_many :CorrectiveActions, class_name: 'CorrectiveAction'

end
