class InspectionPlan < ApplicationRecord
  enum SamplingPlan: [:Fixed, :Percentage, :C0]
  enum Status: [:Draft, :Released, :Retired]


  has_many :Item, class_name: 'Item'
  has_many :Characteristics, class_name: 'InspectionCharacteristic'

end
