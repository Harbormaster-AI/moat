class Admission < ApplicationRecord
  enum AdmissionType: [:Elective, :Emergency, :Urgent, :Newborn, :Trauma]


  has_many :Encounter, class_name: 'Encounter'
  has_many :Facility, class_name: 'Facility'

end
