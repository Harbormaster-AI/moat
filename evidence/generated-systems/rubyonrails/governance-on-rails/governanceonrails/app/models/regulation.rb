class Regulation < ApplicationRecord


  composed_of :uRL,
    class_name: "URL",
    mapping: [
      %w[uRL_value value]
    ]

  has_many :Obligations, class_name: 'Obligation'
  has_many :CompliancePrograms, class_name: 'ComplianceProgram'

end
