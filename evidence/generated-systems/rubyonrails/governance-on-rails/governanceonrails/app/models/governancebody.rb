class GovernanceBody < ApplicationRecord
  enum BodyType: [:Board, :Committee, :Council, :WorkingGroup]


  composed_of :uRL,
    class_name: "URL",
    mapping: [
      %w[uRL_value value]
    ]

  has_many :Organization, class_name: 'Organization'
  has_many :RoleAssignments, class_name: 'RoleAssignment'
  has_many :Policies, class_name: 'Policy'

end
