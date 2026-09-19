class Discharge < ApplicationRecord
  enum Disposition: [:Home, :HomeWithHomeCare, :SkilledNursingFacility, :AcuteCareFacility, :Expired, :AgainstMedicalAdvice]


  has_many :Encounter, class_name: 'Encounter'

end
