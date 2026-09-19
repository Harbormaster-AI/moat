class Role < ApplicationRecord


  has_many :Assignments, class_name: 'RoleAssignment'

end
