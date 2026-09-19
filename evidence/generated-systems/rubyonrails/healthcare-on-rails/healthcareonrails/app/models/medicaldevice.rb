class MedicalDevice < ApplicationRecord
  enum DeviceType: [:Pacemaker, :InsulinPump, :BloodPressureMonitor, :GlucoseMeter, :PulseOximeter, :Ventilator, :InfusionPump, :WearableTracker]
  enum ConnectivityStatus: [:Connected, :Disconnected, :Standby, :Fault]


  has_many :Patient, class_name: 'Patient'
  has_many :Observations, class_name: 'Observation'
  has_many :SoftwareUpdates, class_name: 'SoftwareUpdate'

end
