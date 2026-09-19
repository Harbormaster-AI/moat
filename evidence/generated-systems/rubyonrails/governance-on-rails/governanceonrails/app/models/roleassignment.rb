class RoleAssignment < ApplicationRecord


  has_many :Person, class_name: 'Person'
  has_many :Role, class_name: 'Role'
  has_many :GovernanceBody, class_name: 'GovernanceBody'
  has_many :Organization, class_name: 'Organization'

end
