class ServiceProvider < ApplicationRecord
  enum ProviderType: [:RepairShop, :Towing, :MedicalProvider, :Attorney, :ForensicEngineer, :RentalCar]
  enum NetworkStatus: [:InNetwork, :OutOfNetwork]


  has_many :Claims, class_name: 'Claim'

end
