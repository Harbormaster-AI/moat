class Shift < ApplicationRecord
  enum ShiftType: [:Day, :Swing, :Night, :Weekend]


  has_many :Plant, class_name: 'Plant'
  has_many :Assignments, class_name: 'ShiftAssignment'

end
