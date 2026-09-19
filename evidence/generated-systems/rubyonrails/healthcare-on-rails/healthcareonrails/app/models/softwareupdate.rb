class SoftwareUpdate < ApplicationRecord
  enum UpdateType: [:SecurityPatch, :FeatureUpdate, :BugFix, :FirmwareUpgrade]


  has_many :Device, class_name: 'MedicalDevice'

end
