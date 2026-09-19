class Operation < ApplicationRecord
  enum OperationType: [:Setup, :Run, :Teardown, :Inspection, :Transfer]


  composed_of :timeDuration,
    class_name: "TimeDuration",
    mapping: [
      ${$mapping}, 
      %w[timeDuration_unit unit]
    ]

  composed_of :timeDuration,
    class_name: "TimeDuration",
    mapping: [
      ${$mapping}, 
      %w[timeDuration_unit unit]
    ]

  has_many :Routing, class_name: 'Routing'
  has_many :WorkCenter, class_name: 'WorkCenter'
  has_many :InspectionPlan, class_name: 'InspectionPlan'

end
