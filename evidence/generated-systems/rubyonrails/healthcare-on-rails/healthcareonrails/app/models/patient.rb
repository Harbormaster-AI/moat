class Patient < ApplicationRecord
  enum SexAtBirth: [:Male, :Female, :Unknown]
  enum BloodType: [:APositive, :ANegative, :BPositive, :BNegative, :ABPositive, :ABNegative, :OPositive, :ONegative]


  composed_of :mRN,
    class_name: "MRN",
    mapping: [
      %w[mRN_value value]
    ]

  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Appointments, class_name: 'Appointment'
  has_many :Encounters, class_name: 'Encounter'
  has_many :CarePlans, class_name: 'CarePlan'
  has_many :Allergies, class_name: 'Allergy'
  has_many :Conditions, class_name: 'Condition'
  has_many :MedicationOrders, class_name: 'MedicationOrder'
  has_many :LabOrders, class_name: 'LaboratoryOrder'
  has_many :ImagingOrders, class_name: 'ImagingOrder'
  has_many :Coverages, class_name: 'Coverage'
  has_many :Claims, class_name: 'Claim'
  has_many :Devices, class_name: 'MedicalDevice'
  has_many :Observations, class_name: 'Observation'

end
