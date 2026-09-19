class Screening < ApplicationRecord
  enum Status: [:Ordered, :InProgress, :Clear, :Adverse, :Cancelled]


  has_many :Application, class_name: 'JobApplication'

end
