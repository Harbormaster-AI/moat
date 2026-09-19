class InspectionCharacteristic < ApplicationRecord
  enum MeasurementType: [:Attribute, :Variable]


  composed_of :measurement,
    class_name: "Measurement",
    mapping: [
      ${$mapping}, 
      %w[measurement_unit unit]
    ]

  composed_of :measurement,
    class_name: "Measurement",
    mapping: [
      ${$mapping}, 
      %w[measurement_unit unit]
    ]

  composed_of :measurement,
    class_name: "Measurement",
    mapping: [
      ${$mapping}, 
      %w[measurement_unit unit]
    ]

  has_many :InspectionPlan, class_name: 'InspectionPlan'

end
