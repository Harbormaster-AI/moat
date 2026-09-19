class Facility < ApplicationRecord
  enum FacilityType: [:Hospital, :Clinic, :AmbulatorySurgeryCenter, :UrgentCare, :Laboratory, :ImagingCenter, :Pharmacy]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :HealthSystem, class_name: 'HealthSystem'
  has_many :Departments, class_name: 'Department'
  has_many :CareTeams, class_name: 'CareTeam'
  has_many :Laboratories, class_name: 'Laboratory'
  has_many :ImagingCenters, class_name: 'ImagingCenter'
  has_many :Pharmacies, class_name: 'Pharmacy'
  has_many :InventoryItems, class_name: 'InventoryItem'

end
