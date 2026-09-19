class InspectionResult < ApplicationRecord
  enum ResultStatus: [:Pass, :Fail, :Rework, :Scrap]


  composed_of :measurement,
    class_name: "Measurement",
    mapping: [
      ${$mapping}, 
      %w[measurement_unit unit]
    ]

  has_many :InspectionLot, class_name: 'InspectionLot'
  has_many :Characteristic, class_name: 'InspectionCharacteristic'

end
