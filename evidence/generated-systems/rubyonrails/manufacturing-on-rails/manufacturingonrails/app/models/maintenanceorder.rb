class MaintenanceOrder < ApplicationRecord
  enum Status: [:Created, :Approved, :Scheduled, :InProgress, :Completed, :Cancelled]


  has_many :Asset, class_name: 'Asset'
  has_many :Plan, class_name: 'MaintenancePlan'
  has_many :WorkCenter, class_name: 'WorkCenter'

end
