class Person < ApplicationRecord


  composed_of :emailAddress,
    class_name: "EmailAddress",
    mapping: [
      %w[emailAddress_value value]
    ]

  has_many :RoleAssignments, class_name: 'RoleAssignment'
  has_many :OwnedPolicies, class_name: 'Policy'
  has_many :CorrectiveActions, class_name: 'CorrectiveAction'

end
